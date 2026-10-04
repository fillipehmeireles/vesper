package main

import (
	"time"
	"vesper/internal/cli"
	"vesper/internal/discovery"
	ipv4 "vesper/internal/network"
	"vesper/internal/node"
	"vesper/internal/transport/udp"
	logUtils "vesper/internal/utils"

	"github.com/google/uuid"
)

func main() {
	cliArgs, err := cli.GetArgs()
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
		ID: uuid.NewString(),
			Name: cliArgs.NodeName,
			IpAddr: ip,
			Port: cliArgs.NodePort,
			LastSeenTS: time.Now(),
		}

	transport, err:= udp.New(cliArgs.NodePort)
    defer transport.CloseConnection()
	if err != nil {
		logUtils.LogError(err.Error())
		return
	}

	service := discovery.NewService(&node, transport)

	service.Run()
}
