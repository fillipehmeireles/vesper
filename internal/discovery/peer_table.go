package discovery

import (
	"vesper/internal/node"
)



type PeerTable struct {
	Capacity uint
	table map[string]node.Node
}

func NewPeerTable(capacity uint) *PeerTable {
	return &PeerTable{
		Capacity: capacity,
		table: make(map[string]node.Node, capacity),
	}
}


func (pt *PeerTable) GetPeerTable() map[string]node.Node{
	return pt.table
}
func (pt *PeerTable) Upsert(node node.Node) {
	pt.table[node.ID] = node
}

func (pt *PeerTable) Delete(nodeId string) {
	delete(pt.table, nodeId)
}
