package protocol


import "errors"
/*
 * Message:
 Origin-Node-Id: 0\r\n
 Origin-Node-IPAddr: 192.168.0.25\r\n
 Origin-Node-Name: vsp-01\r\n
 Type: 0\r\n
 Body: pied piper
*/

const (
	KeyOriginNodeID string = "OriginNodeID"
	KeyOriginNodeIPAddr string = "OriginNodeIPAddr"
	KeyOriginNodeName string = "OriginNodeName"
	KeyOriginNodePort string = "OriginNodePort"
	KeyMessageType string = "Type"
	KeyBody string = "Body"
)

func getMessageKeysStrArr() []string {
	return []string{
		KeyOriginNodeID,
		KeyOriginNodeIPAddr,
		KeyOriginNodeName,
		KeyBody,
	}
}

func getMessageKeysIntArr() []string {
	return []string{
		KeyMessageType,
		KeyOriginNodePort,
	}
}

type MessageType int
const  (
	MessageTypeDiscover MessageType = 0
)

type Message struct {
	OriginNodeID string
	OriginNodeIPAddr string
	OriginNodeName string
	OriginNodePort int
	Type MessageType
	Body string
}

func (m Message) Validate() error {
	if m.OriginNodeID == "" {
		return errors.New("origin node id can't be empty")
	}
	if m.OriginNodeName == ""{
		return errors.New("origin node name can't be empty")
	}
	if m.OriginNodeIPAddr == "" {
		return errors.New("origin node ip address can't be empty")
	}
	return nil
}
