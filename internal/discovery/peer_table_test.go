package discovery

import (
	"testing"
	"time"
	"vesper/internal/node"

	"github.com/google/uuid"

	"github.com/stretchr/testify/assert"
)


func TestPeerTable_UpsertShouldUpdatePeerTable(t *testing.T) {
	newNodes := []node.Node {
		{
			ID: uuid.NewString(),
				Name: "vsp01-test",
				IpAddr: "192.168.0.25",
				Port: 6970,
				LastSeenTS: time.Now(),
			},
			{
			ID: uuid.NewString(),
				Name: "vsp01-test",
				IpAddr: "192.168.0.26",
				Port: 6969,
				LastSeenTS: time.Now(),
			},	{
			ID: uuid.NewString(),
			Name: "vsp01-test",
			IpAddr: "192.168.0.27",
			Port: 6071,
			LastSeenTS: time.Now(),
		},
		}


	peerTable := NewPeerTable(3)
	for _, n := range newNodes {
		peerTable.Upsert(n)
	}

	assert.Equal(t, peerTable.table[newNodes[0].ID], newNodes[0])
	assert.Equal(t, peerTable.table[newNodes[1].ID], newNodes[1])
	assert.Equal(t, peerTable.table[newNodes[2].ID], newNodes[2])
}
