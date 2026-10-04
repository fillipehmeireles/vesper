package udp

import "net"

func (t *Transport) Receive() ([]byte, net.Addr, error) {
	// TODO: get buffer size from config
	buf := make([]byte, 1024)

	n, remote, err := t.conn.ReadFromUDP(buf)
	if err != nil {
		return nil, nil, err
	}

	return buf[:n], remote, nil
}
