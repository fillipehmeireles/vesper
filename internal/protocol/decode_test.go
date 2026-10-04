package protocol

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecode_ValidateMessageParsing(t *testing.T) {
	message := "Origin-Node-Id: 1\r\nOrigin-Node-IPAddr: 192.168.0.25\r\nBody: pied piper"
	msg, err := Decode(message)
	assert.NoError(t,err)
	assert.Equal(t,msg.OriginNodeID, 1)
	assert.Equal(t,msg.OriginNodeIPAddr, "192.168.0.25")
	assert.Equal(t,msg.Body, "pied piper")
}
