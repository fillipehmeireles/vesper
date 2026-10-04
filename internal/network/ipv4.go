package ipv4

import (
	"os"
	"net"
	"errors"
)

const (
	ErrGetNodeIP string = "Error on getting the node IP"
)

func GetNodeIP() (string, error) {
	host, err := os.Hostname()
	if err != nil {
		return "", err
	}
	addrs, err := net.LookupIP(host)
	if err != nil {
		return "", err
	}

	for _, addr := range addrs {
		if ipv4 := addr.To4(); ipv4 != nil {
			// TODO: FIX
			if ipv4.String()[:3] != "127" {
				return ipv4.String(), nil
			}
		}
	}
	return "", errors.New(ErrGetNodeIP)
}


