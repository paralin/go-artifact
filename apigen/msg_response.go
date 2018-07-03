package main

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/faceit/go-artifact/protocol"
	"github.com/pkg/errors"
)

// GetResponseMessageID returns the response message ID for the request.
// Error is returned if the request ID indicates there should be a response, but there is none.
func GetResponseMessageID(reqID MsgID) (MsgID, error) {
	if override, ok := msgResponseOverrides[reqID]; ok {
		return override, nil
	}

	msgID := reqID
	// k_EMsgClientToGCStartMatchmaking
	msgStr := msgID.String()
	msgStr = strings.TrimPrefix(msgStr, "k_EMsg")

	responseImplied := strings.HasSuffix(msgStr, "Request")
	msgStr = strings.TrimSuffix(msgStr, "Request")

	clientToGC := strings.HasPrefix(msgStr, "ClientToGC")
	msgStr = strings.TrimPrefix(msgStr, "ClientToGC")

	queryRespStr := func(respStr string) (MsgID, bool) {
		val, ok := protocol.EGCDCGClientMessages_value["k_EMsg"+respStr]
		if ok {
			return MsgID(val), true
		}

		if clientToGC {
			val, ok = protocol.EGCDCGClientMessages_value["k_EMsgGCToClient"+respStr]
			if ok {
				return MsgID(val), true
			}

			val, ok = protocol.EGCDCGClientMessages_value["k_EMsgClientToGC"+respStr]
			if ok {
				return MsgID(val), true
			}
		}

		return MsgID(0), false
	}

	if respID, ok := queryRespStr(msgStr + "Response"); ok {
		return respID, nil
	}

	if responseImplied {
		queryStrs := []string{
			msgStr,
			msgStr + "RequestResponse",
		}
		for _, mstr := range queryStrs {
			if respID, ok := queryRespStr(mstr); ok {
				return respID, nil
			}
		}

		return MsgID(0), errors.Errorf("response was implied by request %v but no response type found", msgID.String())
	}

	return MsgID(0), nil
}

// LookupMessageProtoType lookup proto from message ID.
func LookupMessageProtoType(protoMap map[string]*ProtoType, msgID MsgID) (*ProtoType, error) {
	var protoName string
	if overrideMsg, ok := msgProtoTypeOverrides[msgID]; ok {
		protoName = reflect.TypeOf(overrideMsg).Elem().Name()

		pt, ok := protoMap[protoName]
		if !ok {
			return nil, errors.Errorf("proto not found: %s", protoName)
		}

		return pt, nil
	}

	msgStr := msgID.String()
	msgStr = strings.TrimPrefix(msgStr, "k_EMsg")

	msgStr = strings.TrimPrefix(msgStr, "GC")
	withoutToFrom := strings.Replace(msgStr, "GCToClient", "", -1)
	withoutToFrom = strings.Replace(withoutToFrom, "ClientToGC", "", -1)
	responseToResult := strings.Replace(msgStr, "Response", "Result", -1)
	toAttempt := []string{
		msgStr,
		"GC" + msgStr,
		responseToResult,
		"GC" + responseToResult,
		withoutToFrom,
		"DOTA" + msgStr,
	}

	for _, att := range toAttempt {
		fmt.Println(att)
		if pt, ok := protoMap["CMsg"+att]; ok {
			fmt.Printf("Request: %v matched to type: %v\n", msgID.String(), pt.TypeName)
			return pt, nil
		}
	}

	return nil, errors.Errorf("unable to find proto for: %s", msgID.String())
}
