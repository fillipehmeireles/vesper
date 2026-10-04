package protocol

import "fmt"

func Encode(message Message) (string, error) {
	if err := message.Validate(); err != nil {
		return "", err
	}
	originNodeIdRaw := fmt.Sprintf("%s:%d", KeyOriginNodeID, message.OriginNodeID)
	originNodeIPAddrRaw := fmt.Sprintf("%s:%s", KeyOriginNodeIPAddr, message.OriginNodeIPAddr)
	bodyRaw := fmt.Sprintf("%s:%s", KeyBody, message.Body)

	rawMsg := fmt.Sprintf("%s\r\n%s\r\n%s", originNodeIdRaw, originNodeIPAddrRaw, bodyRaw)

	return rawMsg, nil
}
