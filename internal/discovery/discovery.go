package discovery

import (
	"fmt"
	"net"
	"time"
	"vesper/internal/node"
	"vesper/internal/protocol"
	logUtils "vesper/internal/utils"
)


const DiscoveryPort = 6969

type Transport interface {
	Broadcast([]byte) error
	Receive() ([]byte, net.Addr, error)
}


type Service struct {
	self *node.Node
	peerTable *PeerTable
	tx Transport
}

func NewService(n *node.Node, peerTable *PeerTable, tx Transport) *Service{
	return &Service{
		self: n,
		peerTable: peerTable,
		tx: tx,
	}
}

func (s *Service) Run() error {
	go s.broadcastLoop()

	logUtils.LogInfo(fmt.Sprintf("node: %s", s.self.Name))
	logUtils.LogInfo("listening for peers...")

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
	}
}

func (s *Service) handleMessage(msg protocol.Message, remote net.Addr) {
	switch msg.Type {
	case protocol.MessageTypeDiscover:
		remoteIp := remote.(*net.UDPAddr).IP.String()
		logUtils.LogInfo(
			fmt.Sprintf(
				"peer discovered: %s @ %s",
				msg.OriginNodeName, remoteIp))
		s.peerTable.Upsert(node.Node{
			ID:         msg.OriginNodeID,
			Name:       msg.OriginNodeName,
			IpAddr:     remoteIp,
			LastSeenTS: time.Now(),
		})
	}
}
