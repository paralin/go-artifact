package state

import (
	"github.com/faceit/go-artifact/protocol"
)

// Looks like: Private Lobby -> Lobby is created, match starts.

// State is a snapshot of the client state at a point in time.
type State struct {
	// ConnectionStatus is the status of the connection to the GC.
	ConnectionStatus protocol.GCConnectionStatus
	// Lobby is the current lobby object.
	Lobby *protocol.CSODCGLobby
	// PrivateLobby is the current private lobby object.
	PrivateLobby *protocol.CSODCGPrivateLobby
	// PrivateLobbyInvite is the active incoming lobby invite.
	PrivateLobbyInvite *protocol.CSODCGPrivateLobby_Invite
	// LastConnectionStatusUpdate is the last connection state update we received.
	LastConnectionStatusUpdate *protocol.CMsgConnectionStatus
}

// ClearState clears everything.
func (s *State) ClearState() {
	*s = State{ConnectionStatus: protocol.GCConnectionStatus_GCConnectionStatus_NO_SESSION}
}

// IsReady checks if the client is ready to receive requests.
func (s *State) IsReady() bool {
	return s.ConnectionStatus == protocol.GCConnectionStatus_GCConnectionStatus_HAVE_SESSION
}
