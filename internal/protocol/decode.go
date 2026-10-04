package protocol

import (
	"fmt"
	"strconv"
	"strings"
)



func Decode(rawMsg string) (Message, error) {
	rawSplittedMsg := strings.Split(rawMsg, "\r\n")

	messageMap := make(map[string]string, 3)
	for _, msgData := range rawSplittedMsg {
		kv := strings.Split(msgData,":")
		messageMap[kv[0]] = strings.TrimSpace(kv[1])
	}

	var message Message

	if value, exists := messageMap[KeyOriginNodeID]; !exists {
		return Message{},fmt.Errorf("message should contain: %s"  , value)
	} else {
		n, err := strconv.Atoi(value)
		if err != nil {
			return Message{}, fmt.Errorf("error on converting number %s from string to int", value)
		}
		message.OriginNodeID = n
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
