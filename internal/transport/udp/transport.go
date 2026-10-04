package udp

import "net"

type Transport struct {
	conn *net.UDPConn
	port int
}

func New(port int) (*Transport, error) {
	addr := &net.UDPAddr{
		IP:   net.IPv4zero,
		Port: port,
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		return nil, err
	}

	return &Transport{
		conn: conn,
		port: port,
	}, nil
}

func (tx *Transport) CloseConnection() error {
	return tx.conn.Close()
}
