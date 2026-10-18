package discovery

import (
	"time"
	"vesper/internal/node"
	"vesper/internal/protocol"
	"vesper/internal/transport"
)


const DiscoveryPort = 6969

type Service struct {
	self *node.Node
	PeerTable *PeerTable
	tx transport.Transport
	peerTableChan chan <- NodeTable
}

type PeerTableOpts struct {
	Capacity uint
	PeerTableChan chan <- NodeTable
}

func NewService(n *node.Node, peerTableOpts PeerTableOpts, tx transport.Transport) *Service{
	PeerTable := NewPeerTable(peerTableOpts.Capacity);
	return &Service{
		self: n,
		PeerTable: PeerTable,
		tx: tx,
		peerTableChan: peerTableOpts.PeerTableChan,
	}
}

func (s *Service) Run() error {
	go s.broadcastLoop()
	for {
		payload, remotePeer, err := s.tx.Receive()
		if err != nil {
			continue
		}

		msg, err := protocol.Decode(string(payload))
		if err != nil {
			continue
		}


		if msg.OriginNodeID == s.self.ID{
			continue
		}

		s.handleMessage(msg, remotePeer.Address)

		s.PeerTable.RemoveInactivePeers()
	}
}


func (s *Service) broadcastLoop() {
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		msg := protocol.Message{
			OriginNodeID:   s.self.ID,
			OriginNodeIPAddr: s.self.IpAddr,
			OriginNodeName: s.self.Name,
			OriginNodePort: s.self.Port,
			Type:     protocol.MessageTypeDiscover,
			Body: "pied piper",
		}

		payload, err := protocol.Encode(msg)
		if err != nil {
			continue
		}

		if b, ok := s.tx.(transport.Broadcaster); ok {
			if err := b.Broadcast([]byte(payload)); err != nil {
			}
		}
	}
}

func (s *Service) handleMessage(msg protocol.Message, peerIPAddr string) {
	switch msg.Type {
	case protocol.MessageTypeDiscover:
		s.PeerTable.Upsert(node.Node{
			ID:         msg.OriginNodeID,
			Name:       msg.OriginNodeName,
			IpAddr:     peerIPAddr,
			LastSeenTS: time.Now(),
		})
	}
}
