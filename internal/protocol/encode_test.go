package protocol

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestEncode_ValidateMessageParsing(t *testing.T) {
	message := Message{
		OriginNodeID: 1,
		OriginNodeIPAddr: "192.168.0.25",
		Body: "pied piper",
	}

	expectedRawMessage := "Origin-Node-Id:1\r\nOrigin-Node-IPAddr:192.168.0.25\r\nBody:pied piper"

	rawMsg, err := Encode(message)
	assert.NoError(t, err)
	assert.Equal(t, rawMsg, expectedRawMessage)
}
