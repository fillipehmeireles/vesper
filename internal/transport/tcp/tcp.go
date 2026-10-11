package tcp

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"strconv"
)

type TCPTransport struct {
	connClient *net.Conn
	connServer *net.Conn
	port int
	listener net.Listener
}


func NewTCPTransport(port int) (*TCPTransport, error) {
	return &TCPTransport {
		connClient: nil,
			connServer: nil,
			listener:nil,
			port: port,
		}, nil
}


func (tx *TCPTransport) verifyClientConnection() error {
	if tx.connClient == nil {
		return errors.New("client connection is nil")
	}
	return nil
}

func (tx *TCPTransport) verifyServerConnection() error {
	if tx.connClient == nil {
		return errors.New("server connection is nil")
	}
	return nil
}

func (tx *TCPTransport) verifyListener() error {
	if tx.listener == nil {
		return errors.New("listener is nil")
	}
	return nil
}

func (tx *TCPTransport) Connect(addr string, port int) error {
	if conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", addr,port)); err != nil {
		return err
	} else {
		tx.connClient = &conn
	}
	return nil
}

func (tx *TCPTransport) StartServer() error {
	if err := tx.verifyListener(); err != nil {
		return err
	}

	if ac, err := tx.listener.Accept(); err != nil {
		return err
	} else {
		tx.connServer = &ac
	}

	return nil
}
func (tx *TCPTransport) CreateServer() error {
	if l, err := net.Listen("tcp", strconv.Itoa(tx.port)); err != nil {
		return err
	} else {
		tx.listener = l
	}

	return nil

}

func (tx *TCPTransport) Send(msg []byte, address string) error {
	if err := tx.verifyServerConnection(); err != nil {
		return err
	}

	_, err := (*tx.connServer).Write(msg)
	return err
}

func (tx *TCPTransport) Receive() (string, error) {
	if err := tx.verifyClientConnection(); err != nil {
		return "", err
	}

	message, err :=  bufio.NewReader(*(tx.connClient)).ReadString('\n')
	if err != nil {
		return "", err
	}
	return message, nil
}

func (tx *TCPTransport) CloseConnection () error {
	panic("not implemented yet")
}
