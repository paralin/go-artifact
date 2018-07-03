package artifact

import (
	"context"
	"errors"
	"sync"

	"github.com/sirupsen/logrus"
	devents "github.com/faceit/go-artifact/events"
	pr "github.com/faceit/go-artifact/protocol"
	"github.com/faceit/go-artifact/socache"
	"github.com/faceit/go-artifact/state"
	"github.com/faceit/go-steam"
	"github.com/faceit/go-steam/protocol/gamecoordinator"
	"github.com/golang/protobuf/proto"
)

// AppID is the ID for Artifact
const AppID = 583950

// ErrNotReady is returned when the dota client is not ready.
var ErrNotReady = errors.New("the dota client is not ready to accept requests yet, or has just become unready")

// handlerMap is the map of message types to handler functions.
type handlerMap map[uint32]func(packet *gamecoordinator.GCPacket) error

// Artifact handles the game handler.
type Artifact struct {
	le     *logrus.Entry
	client *steam.Client
	cache  *socache.SOCache

	connectionCtxMtx    sync.Mutex
	connectionCtx       context.Context
	connectionCtxCancel context.CancelFunc

	stateMtx sync.Mutex
	state    state.State

	handlers handlerMap

	pendReqMtx sync.Mutex
	pendReqID  uint32
	pendReq    map[uint32]map[uint32]responseHandler
}

// New builds a new Artifact handler.
func New(client *steam.Client, le *logrus.Entry) *Artifact {
	c := &Artifact{
		le:      le,
		cache:   socache.NewSOCache(le),
		client:  client,
		pendReq: make(map[uint32]map[uint32]responseHandler),

		state: state.State{
			ConnectionStatus: pr.GCConnectionStatus_GCConnectionStatus_NO_SESSION,
		},
	}
	c.buildHandlerMap()
	client.GC.RegisterPacketHandler(c)
	return c
}

// GetCache returns the SO Cache.
func (d *Artifact) GetCache() *socache.SOCache {
	return d.cache
}

// Close kills any ongoing calls.
func (d *Artifact) Close() {
	d.connectionCtxMtx.Lock()
	if d.connectionCtxCancel != nil {
		d.connectionCtxCancel()
	}
	d.connectionCtxMtx.Unlock()
}

// buildHandlerMap builds the map of bound handler functions.
func (d *Artifact) buildHandlerMap() {
	d.handlers = handlerMap{
		// Welcome and conn status
		uint32(pr.EGCBaseClientMsg_k_EMsgGCClientWelcome):          d.handleClientWelcome,
		uint32(pr.EGCBaseClientMsg_k_EMsgGCClientConnectionStatus): d.handleConnectionStatus,

		// Caching
		uint32(pr.ESOMsg_k_ESOMsg_CacheSubscribed):   d.handleCacheSubscribed,
		uint32(pr.ESOMsg_k_ESOMsg_UpdateMultiple):    d.handleCacheUpdateMultiple,
		uint32(pr.ESOMsg_k_ESOMsg_CacheUnsubscribed): d.handleCacheUnsubscribed,
		uint32(pr.ESOMsg_k_ESOMsg_Destroy):           d.handleCacheDestroy,

		// System events
		uint32(pr.EGCBaseClientMsg_k_EMsgGCPingRequest): d.handlePingRequest,
	}

	d.registerGeneratedHandlers()
}

// write sends a message to the game coordinator.
func (d *Artifact) write(messageType uint32, msg proto.Message) {
	d.client.GC.Write(gamecoordinator.NewGCMsgProtobuf(AppID, messageType, msg))
}

// emit emits an event.
func (d *Artifact) emit(event interface{}) {
	d.client.Emit(event)
}

// accessState safely accesses the Artifact state. return true if the state was changed / otherwise updated during the call.
func (d *Artifact) accessState(cb func(nextState *state.State) (bool, error)) error {
	d.stateMtx.Lock()
	defer d.stateMtx.Unlock()

	lastState := d.state
	changed, err := cb(&d.state)
	if err != nil {
		return err
	}
	if changed {
		d.emit(devents.ClientStateChanged{
			OldState: lastState,
			NewState: d.state,
		})
	}
	return nil
}

// unmarshalBody attempts to unmarshal a packet body.
func (d *Artifact) unmarshalBody(packet *gamecoordinator.GCPacket, msg proto.Message) (parseErr error) {
	defer func() {
		if parseErr != nil {
			d.le.WithError(parseErr).WithField("msgtype", packet.MsgType).Warn("unable to parse message")
		}
	}()

	return proto.Unmarshal(packet.Body, msg)
}

// HandleGCPacket handles an incoming game coordinator packet.
func (d *Artifact) HandleGCPacket(packet *gamecoordinator.GCPacket) {
	if packet.AppId != AppID {
		return
	}

	le := d.le.WithField("msgtype", packet.MsgType)
	handler, ok := d.handlers[packet.MsgType]
	if ok && handler != nil {
		if err := handler(packet); err != nil {
			le.WithError(err).Warn("error handling gc msg")
			ok = false
		}
	}

	respHandled := d.handleResponsePacket(le, packet)
	if !ok && !respHandled {
		le.Debug("unhandled gc packet")
		d.emit(&devents.UnhandledGCPacket{
			Packet: packet,
		})
	}
}

// handlePingRequest handles an incoming ping request from the gc.
func (d *Artifact) handlePingRequest(packet *gamecoordinator.GCPacket) error {
	d.write(uint32(pr.EGCBaseClientMsg_k_EMsgGCPingResponse), &pr.CMsgGCClientPing{})
	return nil
}

// getEventEmitter returns a handler that emits an event, used by the generated code.
func (d *Artifact) getEventEmitter(ctor func() devents.Event) func(packet *gamecoordinator.GCPacket) error {
	return func(packet *gamecoordinator.GCPacket) error {
		obj := ctor()
		if err := d.unmarshalBody(packet, obj.GetEventBody()); err != nil {
			return err
		}

		d.emit(obj)
		return nil
	}
}
