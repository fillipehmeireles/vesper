package udp

import (
	"net"
	"vesper/internal/transport"
	logUtils "vesper/internal/utils"
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
