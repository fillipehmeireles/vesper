package udp

import (
	"errors"
	"fmt"
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

	conn, err := net.ListenUDP("udp4", addr)
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
func getBroadcastIPs() ([]net.IP, error) {
    ifaces, err := net.Interfaces()
    if err != nil {
        return nil, err
    }

    var result []net.IP

    for _, iface := range ifaces {
        if iface.Flags&net.FlagUp == 0 ||
            iface.Flags&net.FlagBroadcast == 0 ||
            iface.Flags&net.FlagLoopback != 0 {
            continue
        }

        addrs, err := iface.Addrs()
        if err != nil {
            continue
        }

        for _, addr := range addrs {
            ipnet, ok := addr.(*net.IPNet)
            if !ok {
                continue
            }

            ip := ipnet.IP.To4()
            if ip == nil || len(ipnet.Mask) != net.IPv4len {
                continue
            }

            broadcast := make(net.IP, net.IPv4len)

            for i := range broadcast {
                broadcast[i] = ip[i] | ^ipnet.Mask[i]
            }

            result = append(result, broadcast)
        }
    }

    return result, nil
}

func (tx *UDPTransport) Broadcast(payload []byte) error {
    ips, err := getBroadcastIPs()
    if err != nil {
        return err
    }

    if len(ips) == 0 {
        return fmt.Errorf("no broadcast interfaces available")
    }

    var errs []error
    success := false

    for _, ip := range ips {
        addr := &net.UDPAddr{
            IP:   ip,
            Port: tx.port,
        }

        _, err := tx.conn.WriteToUDP(payload, addr)
        if err != nil {
            errs = append(errs, fmt.Errorf(
                "broadcast to %s: %w", ip, err,
            ))
            continue
        }

        success = true
    }

    if !success {
        return errors.Join(errs...)
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
