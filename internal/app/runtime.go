package app

import (
	"vesper/internal/discovery"
	"vesper/internal/node"
	"vesper/internal/transport"
)


type Runtime struct {
	node      *node.Node
    discovery *discovery.Service
    peers     *discovery.PeerTable
    transport transport.Transport
}


func NewRuntime(node *node.Node, service *discovery.Service, peers *discovery.PeerTable, transport transport.Transport) *Runtime{
	return &Runtime{
		node: node,
		discovery: service,
		peers: peers,
		transport: transport,
	}
}

func (r *Runtime) Run() error {
	return r.discovery.Run()
}

func (r *Runtime) Peers() discovery.NodeTable {
	snapshot := make(discovery.NodeTable)

	for id, peer := range r.discovery.PeerTable.GetPeerTable() {
		snapshot[id] = peer
	}

	return snapshot
}
