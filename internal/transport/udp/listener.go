package udp

import (
	"vesper/internal/transport"
	logUtils "vesper/internal/utils"
)

func (tx *UDPTransport) Receive() ([]byte, transport.Peer, error) {
	// TODO: get buffer size from config
	buf := make([]byte, 1024)

	n, remote, err := tx.conn.ReadFromUDP(buf)
	if err != nil {
		logUtils.LogError(err.Error())
		return nil, transport.Peer{}, err
	}

	return buf[:n], transport.Peer{
		Address: remote.IP.String(),
		Port: remote.Port,
	}, nil
}
