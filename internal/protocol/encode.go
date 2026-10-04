package protocol

import "fmt"

// TODO: encode to bytes instead of string
func Encode(message Message) (string, error) {
	if err := message.Validate(); err != nil {
		return "", err
	}
	originNodeIdRaw := fmt.Sprintf("%s:%s", KeyOriginNodeID, message.OriginNodeID)
	originNodeIPAddrRaw := fmt.Sprintf("%s:%s", KeyOriginNodeIPAddr, message.OriginNodeIPAddr)
	originNodeNameRaw := fmt.Sprintf("%s:%s", KeyOriginNodeName, message.OriginNodeName)
	originNodePortRaw := fmt.Sprintf("%s:%d", KeyOriginNodePort, message.OriginNodePort)
	messageTypeRaw := fmt.Sprintf("%s:%d", KeyMessageType, message.Type)
	bodyRaw := fmt.Sprintf("%s:%s", KeyBody, message.Body)

	rawMsg := fmt.Sprintf("%s\r\n%s\r\n%s\r\n%s\r\n%s\r\n%s",
		originNodeIdRaw,
		originNodeIPAddrRaw,
		originNodePortRaw,
		originNodeNameRaw,
		messageTypeRaw,
		bodyRaw)

	return rawMsg, nil
}
