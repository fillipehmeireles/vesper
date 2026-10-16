package discovery

import (
	"vesper/internal/node"
	"time"
)


type NodeTable map[string]node.Node

type PeerTable struct {
	Capacity uint
	table NodeTable
	historyTable  NodeTable
}

func NewPeerTable(capacity uint) *PeerTable {
	return &PeerTable{
		Capacity: capacity,
		table: make(NodeTable, capacity),
		historyTable: make(NodeTable, capacity),
	}
}


func (pt *PeerTable) GetPeerTable() NodeTable {
	return pt.table
}

func (pt *PeerTable) GetHistoryPeerTable() NodeTable {
	return pt.historyTable
}

func (pt *PeerTable) Upsert(node node.Node) {
	pt.table[node.ID] = node
	pt.historyTable[node.ID] = node
}

func (pt *PeerTable) RemoveInactivePeers() {
	for k, peer := range pt.table {
		if time.Since(peer.LastSeenTS).Seconds() > 10 {
			pt.Delete(k)
		}
	}
}

func (pt *PeerTable) Delete(nodeId string) {
	delete(pt.table, nodeId)
}
