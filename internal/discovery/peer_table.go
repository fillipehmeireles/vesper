package discovery

import (
	"vesper/internal/node"
	"time"
)



type PeerTable struct {
	Capacity uint
	table map[string]node.Node
	historyTable map[string]node.Node
}

func NewPeerTable(capacity uint) *PeerTable {
	return &PeerTable{
		Capacity: capacity,
		table: make(map[string]node.Node, capacity),
		historyTable: make(map[string]node.Node, capacity),
	}
}


func (pt *PeerTable) GetPeerTable() map[string]node.Node{
	return pt.table
}

func (pt *PeerTable) GetHistoryPeerTable() map[string]node.Node{
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
