package tcp

import (
	"net"
	"vesper/internal/transport"
)

type TCPTransport struct {
	conn *net.Conn
}


func NewTCPTransport(addr string) (transport.Transport, error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}

	return &TCPTransport { conn: &conn }, nil
}

func (tx *TCPTransport) Send(msg []byte, address string) error {
	panic("not implemented yet")
	//return nil
}

func (tx *TCPTransport) Receive() ([]byte, transport.Peer, error) {
	panic("not implemented yet")
	//return nil, transport.Peer{}, nil
}

func (tx *TCPTransport) CloseConnection () error {
	panic("not implemented yet")
	//return nil
}
