package discovery

import (
	"fmt"
	"time"
	"vesper/internal/node"
	"vesper/internal/protocol"
	"vesper/internal/transport"
	logUtils "vesper/internal/utils"
)


const DiscoveryPort = 6969

type Service struct {
	self *node.Node
	peerTable *PeerTable
	tx transport.Transport
}

func NewService(n *node.Node, peerTable *PeerTable, tx transport.Transport) *Service{
	return &Service{
		self: n,
		peerTable: peerTable,
		tx: tx,
	}
}

func (s *Service) Run() error {
	go s.broadcastLoop()
	fmt.Printf("node: %s\n", s.self.Name)
	fmt.Println("listening for peers...")
	fmt.Println()
	fmt.Print("\033[s")

	for {
		fmt.Print("\033[u")
		fmt.Print("\033[J")

		fmt.Println("peers:")

		for _, peer := range s.peerTable.GetPeerTable() {
			fmt.Printf(
				"  %-10s %-15s seen %.0fs ago\n",
				peer.Name,
				peer.IpAddr,
				time.Since(peer.LastSeenTS).Seconds(),
			)
		}
		payload, remotePeer, err := s.tx.Receive()
		if err != nil {
			logUtils.LogError(err.Error())
			continue
		}

		msg, err := protocol.Decode(string(payload))
		if err != nil {
			logUtils.LogError(err.Error())
			continue
		}


		if msg.OriginNodeID == s.self.ID{
			continue
		}

		s.handleMessage(msg, remotePeer.Address)
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
			logUtils.LogError(err.Error())
			continue
		}

		if b, ok := s.tx.(transport.Broadcaster); ok {
			if err := b.Broadcast([]byte(payload)); err != nil {
				logUtils.LogError(err.Error())
			}
		}
	}
}

func (s *Service) handleMessage(msg protocol.Message, peerIPAddr string) {
	switch msg.Type {
	case protocol.MessageTypeDiscover:
		s.peerTable.Upsert(node.Node{
			ID:         msg.OriginNodeID,
			Name:       msg.OriginNodeName,
			IpAddr:     peerIPAddr,
			LastSeenTS: time.Now(),
		})
	}
}
