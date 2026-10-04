package udp

import "net"

func (t *Transport) Broadcast(payload []byte) error {
	addr := &net.UDPAddr{
		IP:   net.IPv4bcast,
		Port: t.port,
	}

	_, err := t.conn.WriteToUDP(payload, addr)
	return err
}
