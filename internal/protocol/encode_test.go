package protocol

import (
	"testing"

	"fmt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)


func TestEncode_ValidateMessageParsing(t *testing.T) {
	nodeID:= uuid.NewString()
	message := Message{
		OriginNodeID: nodeID,
		OriginNodeIPAddr: "192.168.0.25",
		OriginNodeName: "vs01",
		Type: MessageTypeDiscover,
		Body: "pied piper",
	}

	expectedRawMessage := fmt.Sprintf("OriginNodeID:%s\r\nOriginNodeIPAddr:192.168.0.25\r\nType:0\r\nBody:pied piper", nodeID)

	rawMsg, err := Encode(message)
	assert.NoError(t, err)
	assert.Equal(t, rawMsg, expectedRawMessage)
}
