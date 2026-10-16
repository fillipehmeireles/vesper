package discovery

import (
	"vesper/internal/discovery"

	"charm.land/bubbles/v2/table"
)

type DiscoveryUIOpts struct {
	PeerTableChan <- chan discovery.NodeTable
}

type Model struct {
    peerTable table.Model
	peerTableChan <- chan discovery.NodeTable
}

func NewModel(discoveryUIOpts DiscoveryUIOpts) Model {
	columns := []table.Column{
        {Title: "Name", Width: 20},
        {Title: "IP Address", Width: 20},
        {Title: "Last Seen", Width: 15},
    }

    t := table.New(
        table.WithColumns(columns),
        table.WithRows([]table.Row{}),
        table.WithFocused(true),
        table.WithHeight(10),
    )


    return Model{
        peerTable: t,
		peerTableChan: discoveryUIOpts.PeerTableChan,
    }
}
