package udp

import (
	"net"
	logUtils "vesper/internal/utils"
)

func (t *Transport) Receive() ([]byte, net.Addr, error) {
	// TODO: get buffer size from config
	buf := make([]byte, 1024)

	n, remote, err := t.conn.ReadFromUDP(buf)
	if err != nil {
		logUtils.LogError(err.Error())
		return nil, nil, err
	}

	return buf[:n], remote, nil
}
