package main

import (
	"fmt"
	"time"
	"vesper/internal/cli"
	"vesper/internal/discovery"
	ipv4 "vesper/internal/network"
	"vesper/internal/node"
	"vesper/internal/transport/udp"
	"vesper/internal/ui"
	discoveryUI "vesper/internal/ui/discovery"
	logUtils "vesper/internal/utils"

	tea "charm.land/bubbletea/v2"
	"github.com/google/uuid"
)


func setupServer(peerTableChan chan <- discovery.NodeTable) {
	cliArgs, err := cli.GetArgs()
	if err != nil 	{
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

	// TODO: get discoveryPort from config/args
	transport, err:= udp.NewUDPTransport(discovery.DiscoveryPort)
    defer transport.CloseConnection()
	if err != nil {
		logUtils.LogError(err.Error())
		return
	}

	service := discovery.NewService(&node,discovery.PeerTableOpts{
		Capacity: 3,
		PeerTableChan: peerTableChan,
	}, transport)

	service.Run()
}
func setupUI(peerTableChan <- chan discovery.NodeTable) {
	uiNodeData := ui.ModelNodeData{NodeName: "node-01", NodeAddr: "192.168.0.25"}
	discoveryUIOpts := discoveryUI.DiscoveryUIOpts{PeerTableChan: peerTableChan}
    p := tea.NewProgram(ui.InitialModel(uiNodeData, discoveryUIOpts))
    if _, err := p.Run(); err != nil {
        logUtils.LogError(fmt.Sprintf("Alas, there's been an error: %v", err))
		return
    }

}
func main() {
	peerTableChan := make (chan discovery.NodeTable, 10)
	go setupServer(peerTableChan)
	setupUI(peerTableChan)
}
