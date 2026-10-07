package udp

import (
	"net"
	logUtils "vesper/internal/utils"
)


func (tx *UDPTransport) Broadcast(payload []byte) error {
	addr := &net.UDPAddr{
		IP:   net.IPv4bcast,
		Port: tx.port,
	}

    if _, err := tx.conn.WriteToUDP(payload, addr); err != nil {
		logUtils.LogError(err.Error())
		return err
	}

	return nil
}
