package tcp

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"strconv"
	"vesper/internal/transport"
)

type TCPTransport struct {
	conn *net.Conn
	port int
	listener *net.Listener
}


func NewTCPTransport(port int) (transport.Transport, error) {
	return &TCPTransport {
		conn: nil,
			listener:nil,
			port: port,
		}, nil
}


func (tx *TCPTransport) verifyConnection() error {
	if tx.conn == nil {
		return errors.New("connection is nil")
	}
	return nil
}

func (tx *TCPTransport) Connect(addr string, port int) error {
	if conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", addr,port)); err != nil {
		return err
	} else {
		tx.conn = &conn
	}
	return nil
}

func (tx *TCPTransport) CreateServer() error {
	if l, err := net.Listen("tcp", strconv.Itoa(tx.port)); err != nil {
		return err
	} else {
		tx.listener = &l
	}

	return nil

}

func (tx *TCPTransport) Send(msg []byte, address string) error {
	if err := tx.verifyConnection(); err != nil {
		return err
	}

	_, err := (*tx.conn).Write(msg)
	return err
}

func (tx *TCPTransport) Receive() (string, error) {
	if err := tx.verifyConnection(); err != nil {
		return "", err
	}

	message, err :=  bufio.NewReader(*(tx.conn)).ReadString('\n')
	if err != nil {
		return "", err
	}
	return message, nil
	//return nil, transport.Peer{}, nil
}

func (tx *TCPTransport) CloseConnection () error {
	panic("not implemented yet")
	//return nil
}
