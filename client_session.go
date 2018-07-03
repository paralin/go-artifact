package artifact

import (
	"context"

	devents "github.com/faceit/go-artifact/events"
	"github.com/faceit/go-artifact/protocol"
	"github.com/faceit/go-artifact/state"
	"github.com/faceit/go-steam/protocol/gamecoordinator"
)

// SetPlaying informs Steam we are playing / not playing Artifact.
func (d *Artifact) SetPlaying(playing bool) {
	if playing {
		d.client.GC.SetGamesPlayed(AppID)
	} else {
		d.client.GC.SetGamesPlayed()
		_ = d.accessState(func(ns *state.State) (changed bool, err error) {
			ns.ClearState()
			return true, nil
		})
	}
}

// SayHello says hello to the Dota2 server, in an attempt to get a session.
func (d *Artifact) SayHello(haveCacheVersions ...*protocol.CMsgSOCacheHaveVersion) {
	d.le.Debug("sending hello to GC")
	partnerAccType := protocol.PartnerAccountType_PARTNER_NONE
	engine := protocol.ESourceEngine_k_ESE_Source2
	var clientSessionNeed uint32 = 104
	d.write(uint32(protocol.EGCBaseClientMsg_k_EMsgGCClientHello), &protocol.CMsgClientHello{
		ClientLauncher:      &partnerAccType,
		Engine:              &engine,
		ClientSessionNeed:   &clientSessionNeed,
		SocacheHaveVersions: haveCacheVersions,
	})
}

// validateConnectionContext checks if the client is ready or not.
func (d *Artifact) validateConnectionContext() (context.Context, error) {
	d.connectionCtxMtx.Lock()
	defer d.connectionCtxMtx.Unlock()

	cctx := d.connectionCtx
	if cctx == nil {
		return nil, ErrNotReady
	}

	select {
	case <-cctx.Done():
		return nil, ErrNotReady
	default:
		return cctx, nil
	}
}

// handleClientWelcome handles an incoming client welcome event.
func (d *Artifact) handleClientWelcome(packet *gamecoordinator.GCPacket) error {
	welcome := &protocol.CMsgClientWelcome{}
	if err := d.unmarshalBody(packet, welcome); err != nil {
		return err
	}

	d.le.Debug("received GC welcome")
	for _, cache := range welcome.GetUptodateSubscribedCaches() {
		d.RequestCacheSubscriptionRefresh(cache.GetOwnerSoid())
	}

	for _, cache := range welcome.GetOutofdateSubscribedCaches() {
		if err := d.cache.HandleSubscribed(cache); err != nil {
			d.le.WithError(err).Warn("unable to handle welcome cache")
		}
	}

	d.setConnectionStatus(protocol.GCConnectionStatus_GCConnectionStatus_HAVE_SESSION, nil)
	d.emit(&devents.ClientWelcomed{Welcome: welcome})
	return nil
}

// handleConnectionStatus handles the connection status update event.
func (d *Artifact) handleConnectionStatus(packet *gamecoordinator.GCPacket) error {
	stat := &protocol.CMsgConnectionStatus{}
	if err := d.unmarshalBody(packet, stat); err != nil {
		return err
	}

	if stat.Status == nil {
		return nil
	}

	d.setConnectionStatus(*stat.Status, stat)
	return nil
}

// setConnectionStatus sets the connection status, and emits an event.
// NOTE: do not call from inside accessState.
func (d *Artifact) setConnectionStatus(
	connStatus protocol.GCConnectionStatus,
	update *protocol.CMsgConnectionStatus,
) {
	_ = d.accessState(func(ns *state.State) (changed bool, err error) {
		if ns.ConnectionStatus == connStatus {
			return false, nil
		}

		oldState := ns.ConnectionStatus
		d.le.WithField("old", oldState.String()).
			WithField("new", connStatus.String()).
			Debug("connection status changed")
		d.emit(&devents.GCConnectionStatusChanged{
			OldState: oldState,
			NewState: connStatus,
			Update:   update,
		})

		ns.ClearState() // every time the state changes, we lose the lobbies / etc
		ns.ConnectionStatus = connStatus
		d.connectionCtxMtx.Lock()
		if d.connectionCtxCancel != nil {
			d.connectionCtxCancel()
			d.connectionCtxCancel = nil
			d.connectionCtx = nil
		}
		if connStatus == protocol.GCConnectionStatus_GCConnectionStatus_HAVE_SESSION {
			d.connectionCtx, d.connectionCtxCancel = context.WithCancel(context.Background())
		}
		d.connectionCtxMtx.Unlock()
		return true, nil
	})
}
