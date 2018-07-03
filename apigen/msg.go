package main

import (
	"github.com/faceit/go-artifact/protocol"
	"sort"
)

// IsValidMsg checks if the message is valid.
func IsValidMsg(msg MsgID) bool {
	_, ok := protocol.EGCDCGClientMessages_name[int32(msg)]
	return ok
}

func getSortedMsgIDs() []MsgID {
	var msgIds []MsgID
	for msgIDNum := range protocol.EGCDCGClientMessages_name {
		msgID := MsgID(msgIDNum)
		msgIds = append(msgIds, msgID)
	}

	sort.Slice(msgIds, func(i int, j int) bool {
		return msgIds[i] < msgIds[j]
	})
	return msgIds
}
