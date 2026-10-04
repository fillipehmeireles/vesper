package protocol

import (
	"fmt"
	"strings"
)



func Decode(rawMsg string) (Message, error) {
	rawSplittedMsg := strings.Split(rawMsg, "\r\n")

	messageMap := make(map[string]string, 3)
	for _, msgData := range rawSplittedMsg {
		kv := strings.Split(msgData,":")
		messageMap[kv[0]] = kv[1]
	}

	var message Message

	if value, exists := messageMap[KeyOriginNodeID]; !exists {
		return Message{},fmt.Errorf("message should contain: %s"  , value)
	} else {
		message.OriginNodeID = value
	}

	if value, exists := messageMap[KeyOriginNodeIPAddr]; !exists {
		return Message{}, fmt.Errorf("message should contain: %s"  , value)
	} else {
		message.OriginNodeIPAddr = value
	}

	if value, exists := messageMap[KeyBody]; !exists {
		return Message{}, fmt.Errorf("message should contain: %s"  , value)
	} else {
		message.Body = value
	}


	return message, nil
}
