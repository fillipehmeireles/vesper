package udp

import (
	"net"
	logUtils "vesper/internal/utils"
)


func (t *Transport) Broadcast(payload []byte) error {
	addr := &net.UDPAddr{
		IP:   net.IPv4bcast,
		Port: t.port,
	}

	if err := t.SendMsg(payload, addr); err != nil {
		logUtils.LogError(err.Error())
		return err
	}
	return nil
}
