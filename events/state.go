package events

import (
	"github.com/faceit/go-artifact/state"
)

// ClientStateChanged is emitted whenever anything about the client state changes.
type ClientStateChanged struct {
	// OldState is the old state.
	OldState state.State
	// NewState is the new state.
	NewState state.State
}
