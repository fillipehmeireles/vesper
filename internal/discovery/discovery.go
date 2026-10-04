package discovery

import (
	"net"
	"time"
	"vesper/internal/node"
	logUtils "vesper/internal/utils"
)


type Transport interface {
	Broadcast([]byte) error
	Receive() ([]byte, net.Addr, error)
}


type Service struct {
	self node.Node
	// peerTable *PeerTable
	tx Transport
}

func (s *Service) Run() error {
	go s.broadcastLoop()

	for {
		payload, remote, err := s.tx.Receive()
		if err != nil {
			continue
		}

		msg, err := protocol.Decode(payload)
		if err != nil {
			continue
		}

		if msg.NodeID == s.self.ID {
			continue
		}

		s.handleMessage(msg, remote)
	}
}

func (s *Service) broadcastLoop() {
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		msg := protocol.Message{
			Type:     protocol.Discover,
			NodeID:   s.self.Id,
			NodeName: s.self.Name,
		}

		payload, err := protocol.Encode(msg)
		if err != nil {
			continue
		}

		_ = s.tx.Broadcast(payload)
	}
}

func (s *Service) handleMessage(msg protocol.Message, remote net.Addr) {
	logUtils.LogInfo(msg)
	/*
	switch msg.Type {
	case protocol.Discover:
		s.peers.Upsert(node.Node{
			ID:         msg.NodeID,
			Name:       msg.NodeName,
			IpAddr:     remote.(*net.UDPAddr).IP.String(),
			LastSeenTS: time.Now(),
		})
	}
	*/
}
