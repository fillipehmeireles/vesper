package discovery

import (
	"fmt"
	"net"
	"time"
	"vesper/internal/node"
	"vesper/internal/protocol"
	logUtils "vesper/internal/utils"
)


type Transport interface {
	Broadcast([]byte) error
	Receive() ([]byte, net.Addr, error)
}


type Service struct {
	self *node.Node
	// peerTable *PeerTable
	tx Transport
}

func NewService(n *node.Node, tx Transport) *Service{
	return &Service{
		self: n,
		tx: tx,
	}
}

func (s *Service) Run() error {
	go s.broadcastLoop()

	logUtils.LogInfo(fmt.Sprintf("%s (%s) Listening", s.self.Name, s.self.IpAddr))

	for {
		payload, remote, err := s.tx.Receive()
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


		fmt.Printf("receiving: %+v\n", msg)
		s.handleMessage(msg, remote)
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

		if err := s.tx.Broadcast([]byte(payload)); err != nil {
			logUtils.LogError(err.Error())
		}

		logUtils.LogInfo("broadcasting!")
	}
}

func (s *Service) handleMessage(msg protocol.Message, remote net.Addr) {
	fmt.Printf("OriginNodeID: %s nodeID: %s", msg.OriginNodeID, s.self.ID)
	fmt.Println(msg)
	fmt.Println(remote)
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
