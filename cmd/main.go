package main

import (
	"fmt"
	"log"
	"net"
	"time"
	"vesper/internal/cli"
	ipv4 "vesper/internal/network"
	"vesper/internal/node"
	logUtils "vesper/internal/utils"
)


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

    buffer := make([]byte, 1024)
    for {
        n, clientAddr, err := conn.ReadFromUDP(buffer)
        if err != nil {
            log.Println("Read error:", err)
            continue
        }
        logUtils.LogInfo(fmt.Sprintf("[INFO] %s: Received from %s: %s\n", node.Name, clientAddr, buffer[:n]))
    }

}
