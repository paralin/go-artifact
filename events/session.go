package events

import (
	"github.com/faceit/go-artifact/protocol"
	"github.com/faceit/go-steam/protocol/gamecoordinator"
)

// GCConnectionStatusChanged is emitted when the client connection state is updated.
type GCConnectionStatusChanged struct {
	// OldState contains the old connection status.
	OldState protocol.GCConnectionStatus
	// NewState contains the new connection status.
	NewState protocol.GCConnectionStatus
	// Update contains the message from the server that triggered this change, may be nil.
	Update *protocol.CMsgConnectionStatus
}

// ClientWelcomed is emitted when the client receives the GC welcome
type ClientWelcomed struct {
	// Welcome is the welcome message from the GC.
	Welcome *protocol.CMsgClientWelcome
}

// UnhandledGCPacket is called when the client ignores an unhandled packet.
type UnhandledGCPacket struct {
	// Packet is the unhandled packet.
	Packet *gamecoordinator.GCPacket
}
