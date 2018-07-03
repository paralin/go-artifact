package main

import (
	"fmt"

	gcm "github.com/faceit/go-artifact/protocol"
)

type MsgID = gcm.EGCDCGClientMessages
type generatedEventHandler struct {
	msgID     MsgID
	eventName string
	eventType *ProtoType
}

// buildGeneratedEventHandler builds a generated event handler.
func buildGeneratedEventHandler(
	msgID MsgID,
	protoMap map[string]*ProtoType,
	eventImports map[string]struct{},
) (*generatedEventHandler, error) {
	eventProtoType, err := LookupMessageProtoType(protoMap, msgID)
	if err != nil {
		return nil, err
	}
	eventImports[eventProtoType.Pak.Path()] = struct{}{}

	return &generatedEventHandler{
		msgID:     msgID,
		eventName: GetMessageEventName(msgID),
		eventType: eventProtoType,
	}, nil
}

func (g *generatedEventHandler) generateComment() string {
	return fmt.Sprintf("// %s event.\n// MessageID: %s\n", g.eventName, g.msgID.String())
}
