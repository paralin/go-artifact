package artifact

import (
	"github.com/faceit/go-artifact/protocol"
)

// SendChannelMessage attempts to send a message in a channel, text-only.
// Use SendChatMessage for more fine-grained control.
func (d *Artifact) SendChannelMessage(channelID uint64, message string) {
	d.write(uint32(protocol.EGCDCGClientMessages_k_EMsgClientToGCSendChatMessage), &protocol.CMsgClientToGCSendChatMessage{
		ChatRoomId: &channelID,
		ChatMsg:    &message,
	})
}
