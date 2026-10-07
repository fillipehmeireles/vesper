package udp

import (
	"net"
	"vesper/internal/transport"
)

type UDPTransport struct {
	conn *net.UDPConn
	port int
}

func NewUDPTransport(port int) (transport.Transport, error) {
	addr := &net.UDPAddr{
		IP:   net.IPv4zero,
		Port: port,
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		return nil, err
	}

	return &UDPTransport{
		conn: conn,
		port: port,
	}, nil
}


func (tx* UDPTransport) Send(payload []byte, address string) error {
	addr, err := net.ResolveUDPAddr("udp", address)
    if err != nil {
        return err
    }

    _, err = tx.conn.WriteToUDP(payload, addr)
    return err
}

func (tx *UDPTransport) CloseConnection() error {
	return tx.conn.Close()
}
