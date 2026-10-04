package protocol

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecode_ValidateMessageParsing(t *testing.T) {
	message := "OriginNodeID: 1\r\nOriginNodeIPAddr: 192.168.0.25\r\nOriginNodeName:vs01\r\nType:0\r\nBody: pied piper"
	msg, err := Decode(message)
	assert.NoError(t,err)
	assert.Equal(t,msg.OriginNodeID, "1")
	assert.Equal(t,msg.OriginNodeIPAddr, "192.168.0.25")
	assert.Equal(t,msg.Type, MessageTypeDiscover)
	assert.Equal(t,msg.Body, "pied piper")
}
