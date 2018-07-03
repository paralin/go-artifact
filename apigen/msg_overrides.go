package main

import (
	dm "github.com/faceit/go-artifact/protocol"
	"github.com/golang/protobuf/proto"
)

// msgSenderOverrides overrides the heuristic-generated sender parties for each message
// Most of the MsgSenderNone messages are not GC<->Client messages.
var msgSenderOverrides = map[MsgID]MsgSender{
	// dm.EDOTAGCMsg_k_EMsgGCLiveScoreboardUpdate:      MsgSenderNone,
	dm.EGCDCGClientMessages_k_EMsgClientToGCSendChatMessage: MsgSenderClient,
}

// msgMethodNameOverrides overrides the generated client method names.
var msgMethodNameOverrides = map[MsgID]string{
	// dm.EDOTAGCMsg_k_EMsgGameAutographReward:               "AutographReward",
	// dm.EGCDCGClientMessages_k_EMsgClientToGCStartMatchmaking: "StartMatchmaking",
}

// msgResponseOverrides maps request message IDs to response message IDs.
// Setting zero as the value indicates it is an action and not a query
var msgResponseOverrides = map[MsgID]MsgID{
	// dm.EGCDCGClientMessages_k_EMsgClientToGCAIMatchStarted: 0,
	// dm.EDOTAGCMsg_k_EMsgClientToGCMyTeamInfoRequest:    dm.EDOTAGCMsg_k_EMsgGCToClientTeamInfo,
}

// msgProtoTypeOverrides overrides the GC message to proto mapping.
var msgProtoTypeOverrides = map[MsgID]proto.Message{
	// dm.EDOTAGCMsg_k_EMsgGCToClientTeamInfo: &dota_gcmessages_client_team.CMsgDOTATeamInfo{},
}

// msgArgAsParameterOverrides indicates that call arguments should be exposed as
// parameters and not an object.
var msgArgAsParameterOverrides = map[MsgID]bool{
	// dm.EDOTAGCMsg_k_EMsgGCPracticeLobbySetDetails: true,
}

// msgEventNameOverrides are overrides for event names
var msgEventNameOverrides = map[MsgID]string{
	// dm.EDOTAGCMsg_k_EMsgGCTeamInvite_GCRequestToInvitee:  "TeamInviteReceived",
}
