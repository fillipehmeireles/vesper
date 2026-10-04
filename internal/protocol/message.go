package protocol


import "errors"
/*
 * Message:
 Origin-Node-Id: 0\r\n
 Origin-Node-IPAddr: 192.168.0.25\r\n
 Body: pied piper\r\n
 \r\n
*/

const (
	KeyOriginNodeID string = "Origin-Node-Id"
	KeyOriginNodeIPAddr string = "Origin-Node-IPAddr"
	KeyBody string = "Body"
)

type Message struct {
	OriginNodeID int
	OriginNodeIPAddr string
	Body string
}

func (m Message) Validate() error {
	if m.OriginNodeID == 0 {
		return errors.New("origin node id can't be empty")
	}
	if m.OriginNodeIPAddr == "" {
		return errors.New("origin node ip address can't be empty")
	}
	return nil
}
