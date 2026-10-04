package protocol

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSomething(t *testing.T) {
	message := "Origin-Node-Id: 0\r\nOrigin-Node-IPAddr: 192.168.0.25\r\nBody: pied piper"
	msg, err := Decode(message)
	assert.NoError(t,err)
	fmt.Println(msg)
}
