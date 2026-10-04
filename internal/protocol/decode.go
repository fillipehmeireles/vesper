package protocol

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// TODO: decode []byte instead of string
func Decode(rawMsg string) (Message, error) {
	rawSplittedMsg := strings.Split(rawMsg, "\r\n")

	messageMap := make(map[string]string, 3)
	for _, msgData := range rawSplittedMsg {
		kv := strings.Split(msgData,":")
		messageMap[kv[0]] = strings.TrimSpace(kv[1])
	}

	var message Message
	for _, key := range getMessageKeysStrArr() {
		if value, exists := messageMap[key]; !exists {
			return Message{}, fmt.Errorf("message should contain: %s"  , value)
		} else {
			v := reflect.ValueOf(&message).Elem()
			field := v.FieldByName(key)
			if field.IsValid() && field.CanSet() {
				field.SetString(value)
			}
		}
	}

	for _, key := range getMessageKeysIntArr() {
		if value, exists := messageMap[key]; !exists {
			return Message{}, fmt.Errorf("message should contain: %s"  , value)
		} else {
			v := reflect.ValueOf(&message).Elem()
			field := v.FieldByName(key)
			if field.IsValid() && field.CanSet() {
				n, err := strconv.Atoi(value)
				if err != nil {
					return Message{}, fmt.Errorf("error on converting number %s from string to int", value)
				}
				field.SetInt(int64(n))
			}
		}
	}
	if err := message.Validate(); err != nil {
		return Message{}, err
	}

	return message, nil
}
