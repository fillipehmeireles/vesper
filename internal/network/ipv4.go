package ipv4

import (
	"net"
	"errors"
)

const (
	ErrGetNodeIP string = "Error on getting the node IP"
)

func GetNodeIP() (string, error) {
	conn, err := net.Dial("udp4", "8.8.8.8:80")
	if err != nil {
		return "", err
	}
	defer conn.Close()

	addr, ok := conn.LocalAddr().(*net.UDPAddr)
	if !ok {
		return "", errors.New("unexpected local address type")
	}

	return addr.IP.String(), nil
}
