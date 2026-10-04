package main

import (
	"fmt"
	"net"
	"time"
	"vesper/internal/cli"
	ipv4 "vesper/internal/network"
	"vesper/internal/node"
	logUtils "vesper/internal/utils"
)

func listen(conn *net.UDPConn) {
	buf := make([]byte, 1024)

	for {
		n, remote, err := conn.ReadFromUDP(buf)
		if err != nil {
			logUtils.LogError(err.Error())
			continue
		}
		logUtils.LogInfo(fmt.Sprintf(
			"received %q from %s\n",
			string(buf[:n]),
			remote,
		))
	}
}

func broadcast(conn *net.UDPConn, nodeName string) error {
	addr := &net.UDPAddr{
		IP:   net.IPv4bcast,
		Port: 6969,
	}

	msg := []byte("VESPER|DISCOVER|" + nodeName)

	_, err := conn.WriteToUDP(msg, addr)
	return err
}

func main() {
	name, err := cli.GeNameFlag()
	if err != nil {
		logUtils.LogError(err.Error())
		return
	}

	ip, err := ipv4.GetNodeIP()
	if err != nil {
		logUtils.LogError(err.Error())
		return
	}

	node := node.Node {
		Id: 0,
			Name: name,
			IpAddr: ip,
			LastSeenTS: time.Now(),
		}

	addr, err := net.ResolveUDPAddr("udp", ":6969")
    if err != nil {
        logUtils.LogError(fmt.Sprintf("Couldn’t resolve address: %s", err.Error()))
		return
    }

    conn, err := net.ListenUDP("udp", addr)
    if err != nil {
        logUtils.LogError(fmt.Sprintf("Listen failed: %s", err.Error()))
		return
    }
	logUtils.LogInfo(fmt.Sprintf("%s (%s) Listening", node.Name, node.IpAddr))
    defer conn.Close()


	go listen(conn)

    for {
		if err := broadcast(conn, node.Name); err != nil {
			logUtils.LogError(err.Error())
		}

		time.Sleep(2 * time.Second)
    }

}
