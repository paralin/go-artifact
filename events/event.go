package events

import (
	"github.com/faceit/go-artifact/protocol"
	"github.com/golang/protobuf/proto"
)

// Event is a DOTA event.
type Event interface {
	// GetEventMsgID returns the DOTA event message ID.
	GetEventMsgID() protocol.EGCDCGClientMessages
	// GetEventBody event body.
	GetEventBody() proto.Message
	// GetEventName returns the event name.
	GetEventName() string
}
