package events

import (
	"github.com/faceit/go-artifact/protocol"
	"github.com/golang/protobuf/proto"
)

// AvailableGauntlets event.
// MessageID: k_EMsgGCToClientAvailableGauntlets
type AvailableGauntlets struct {
	protocol.CMsgGCToClientAvailableGauntlets
}

// GetEventMsgID returns the dota message ID of the event.
func (e *AvailableGauntlets) GetEventMsgID() protocol.EGCDCGClientMessages {
	return protocol.EGCDCGClientMessages_k_EMsgGCToClientAvailableGauntlets
}

// GetEventBody returns the event body.
func (e *AvailableGauntlets) GetEventBody() proto.Message {
	return &e.CMsgGCToClientAvailableGauntlets
}

// GetEventName returns the event name.
func (e *AvailableGauntlets) GetEventName() string {
	return "AvailableGauntlets"
}

// ChatChannelJoined event.
// MessageID: k_EMsgGCToClientChatChannelJoined
type ChatChannelJoined struct {
	protocol.CMsgGCToClientChatChannelJoined
}

// GetEventMsgID returns the dota message ID of the event.
func (e *ChatChannelJoined) GetEventMsgID() protocol.EGCDCGClientMessages {
	return protocol.EGCDCGClientMessages_k_EMsgGCToClientChatChannelJoined
}

// GetEventBody returns the event body.
func (e *ChatChannelJoined) GetEventBody() proto.Message {
	return &e.CMsgGCToClientChatChannelJoined
}

// GetEventName returns the event name.
func (e *ChatChannelJoined) GetEventName() string {
	return "ChatChannelJoined"
}

// ChatMessage event.
// MessageID: k_EMsgGCToClientChatMessage
type ChatMessage struct {
	protocol.CMsgGCToClientChatMessage
}

// GetEventMsgID returns the dota message ID of the event.
func (e *ChatMessage) GetEventMsgID() protocol.EGCDCGClientMessages {
	return protocol.EGCDCGClientMessages_k_EMsgGCToClientChatMessage
}

// GetEventBody returns the event body.
func (e *ChatMessage) GetEventBody() proto.Message {
	return &e.CMsgGCToClientChatMessage
}

// GetEventName returns the event name.
func (e *ChatMessage) GetEventName() string {
	return "ChatMessage"
}

// DevLobbyStatus event.
// MessageID: k_EMsgGCToClientDevLobbyStatus
type DevLobbyStatus struct {
	protocol.CMsgGCToClientDevLobbyStatus
}

// GetEventMsgID returns the dota message ID of the event.
func (e *DevLobbyStatus) GetEventMsgID() protocol.EGCDCGClientMessages {
	return protocol.EGCDCGClientMessages_k_EMsgGCToClientDevLobbyStatus
}

// GetEventBody returns the event body.
func (e *DevLobbyStatus) GetEventBody() proto.Message {
	return &e.CMsgGCToClientDevLobbyStatus
}

// GetEventName returns the event name.
func (e *DevLobbyStatus) GetEventName() string {
	return "DevLobbyStatus"
}

// GlobalPhantomLeagues event.
// MessageID: k_EMsgGCToClientGlobalPhantomLeagues
type GlobalPhantomLeagues struct {
	protocol.CMsgGCToClientGlobalPhantomLeagues
}

// GetEventMsgID returns the dota message ID of the event.
func (e *GlobalPhantomLeagues) GetEventMsgID() protocol.EGCDCGClientMessages {
	return protocol.EGCDCGClientMessages_k_EMsgGCToClientGlobalPhantomLeagues
}

// GetEventBody returns the event body.
func (e *GlobalPhantomLeagues) GetEventBody() proto.Message {
	return &e.CMsgGCToClientGlobalPhantomLeagues
}

// GetEventName returns the event name.
func (e *GlobalPhantomLeagues) GetEventName() string {
	return "GlobalPhantomLeagues"
}

// MarketPrices event.
// MessageID: k_EMsgGCToClientMarketPrices
type MarketPrices struct {
	protocol.CMsgGCToClientMarketPrices
}

// GetEventMsgID returns the dota message ID of the event.
func (e *MarketPrices) GetEventMsgID() protocol.EGCDCGClientMessages {
	return protocol.EGCDCGClientMessages_k_EMsgGCToClientMarketPrices
}

// GetEventBody returns the event body.
func (e *MarketPrices) GetEventBody() proto.Message {
	return &e.CMsgGCToClientMarketPrices
}

// GetEventName returns the event name.
func (e *MarketPrices) GetEventName() string {
	return "MarketPrices"
}

// MatchmakingStatus event.
// MessageID: k_EMsgGCToClientMatchmakingStatus
type MatchmakingStatus struct {
	protocol.CMsgGCToClientMatchmakingStatus
}

// GetEventMsgID returns the dota message ID of the event.
func (e *MatchmakingStatus) GetEventMsgID() protocol.EGCDCGClientMessages {
	return protocol.EGCDCGClientMessages_k_EMsgGCToClientMatchmakingStatus
}

// GetEventBody returns the event body.
func (e *MatchmakingStatus) GetEventBody() proto.Message {
	return &e.CMsgGCToClientMatchmakingStatus
}

// GetEventName returns the event name.
func (e *MatchmakingStatus) GetEventName() string {
	return "MatchmakingStatus"
}

// MatchmakingStopped event.
// MessageID: k_EMsgGCToClientMatchmakingStopped
type MatchmakingStopped struct {
	protocol.CMsgGCToClientMatchmakingStopped
}

// GetEventMsgID returns the dota message ID of the event.
func (e *MatchmakingStopped) GetEventMsgID() protocol.EGCDCGClientMessages {
	return protocol.EGCDCGClientMessages_k_EMsgGCToClientMatchmakingStopped
}

// GetEventBody returns the event body.
func (e *MatchmakingStopped) GetEventBody() proto.Message {
	return &e.CMsgGCToClientMatchmakingStopped
}

// GetEventName returns the event name.
func (e *MatchmakingStopped) GetEventName() string {
	return "MatchmakingStopped"
}

// MessageOfTheDay event.
// MessageID: k_EMsgGCToClientMessageOfTheDay
type MessageOfTheDay struct {
	protocol.CMsgGCToClientMessageOfTheDay
}

// GetEventMsgID returns the dota message ID of the event.
func (e *MessageOfTheDay) GetEventMsgID() protocol.EGCDCGClientMessages {
	return protocol.EGCDCGClientMessages_k_EMsgGCToClientMessageOfTheDay
}

// GetEventBody returns the event body.
func (e *MessageOfTheDay) GetEventBody() proto.Message {
	return &e.CMsgGCToClientMessageOfTheDay
}

// GetEventName returns the event name.
func (e *MessageOfTheDay) GetEventName() string {
	return "MessageOfTheDay"
}

// PrivateLobbyEvent event.
// MessageID: k_EMsgGCToClientPrivateLobbyEvent
type PrivateLobbyEvent struct {
	protocol.CMsgGCToClientPrivateLobbyEvent
}

// GetEventMsgID returns the dota message ID of the event.
func (e *PrivateLobbyEvent) GetEventMsgID() protocol.EGCDCGClientMessages {
	return protocol.EGCDCGClientMessages_k_EMsgGCToClientPrivateLobbyEvent
}

// GetEventBody returns the event body.
func (e *PrivateLobbyEvent) GetEventBody() proto.Message {
	return &e.CMsgGCToClientPrivateLobbyEvent
}

// GetEventName returns the event name.
func (e *PrivateLobbyEvent) GetEventName() string {
	return "PrivateLobbyEvent"
}

// SDRTicket event.
// MessageID: k_EMsgGCToClientSDRTicket
type SDRTicket struct {
	protocol.CMsgGCToClientSDRTicket
}

// GetEventMsgID returns the dota message ID of the event.
func (e *SDRTicket) GetEventMsgID() protocol.EGCDCGClientMessages {
	return protocol.EGCDCGClientMessages_k_EMsgGCToClientSDRTicket
}

// GetEventBody returns the event body.
func (e *SDRTicket) GetEventBody() proto.Message {
	return &e.CMsgGCToClientSDRTicket
}

// GetEventName returns the event name.
func (e *SDRTicket) GetEventName() string {
	return "SDRTicket"
}

// TournamentState event.
// MessageID: k_EMsgGCToClientTournamentState
type TournamentState struct {
	protocol.CMsgGCToClientTournamentState
}

// GetEventMsgID returns the dota message ID of the event.
func (e *TournamentState) GetEventMsgID() protocol.EGCDCGClientMessages {
	return protocol.EGCDCGClientMessages_k_EMsgGCToClientTournamentState
}

// GetEventBody returns the event body.
func (e *TournamentState) GetEventBody() proto.Message {
	return &e.CMsgGCToClientTournamentState
}

// GetEventName returns the event name.
func (e *TournamentState) GetEventName() string {
	return "TournamentState"
}

// UserJoinedChatChannel event.
// MessageID: k_EMsgGCToClientUserJoinedChatChannel
type UserJoinedChatChannel struct {
	protocol.CMsgGCToClientUserJoinedChatChannel
}

// GetEventMsgID returns the dota message ID of the event.
func (e *UserJoinedChatChannel) GetEventMsgID() protocol.EGCDCGClientMessages {
	return protocol.EGCDCGClientMessages_k_EMsgGCToClientUserJoinedChatChannel
}

// GetEventBody returns the event body.
func (e *UserJoinedChatChannel) GetEventBody() proto.Message {
	return &e.CMsgGCToClientUserJoinedChatChannel
}

// GetEventName returns the event name.
func (e *UserJoinedChatChannel) GetEventName() string {
	return "UserJoinedChatChannel"
}
