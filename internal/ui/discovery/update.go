package discovery

import (
	"strconv"
	"time"

	"charm.land/bubbles/v2/table"
	tea "charm.land/bubbletea/v2"
)

func (m Model) WaitForPeers() tea.Cmd {
    return func() tea.Msg {
        peers, ok := <-m.peerTableChan
        if !ok {
            return nil
        }

        return PeersUpdatedMsg{Peers: peers}
    }
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
    switch msg := msg.(type) {
    case PeersUpdatedMsg:
        var rows []table.Row

        for _, peer := range msg.Peers {
            rows = append(rows, table.Row{
                peer.Name,
                peer.IpAddr,
                strconv.FormatFloat(
                    time.Since(peer.LastSeenTS).Seconds(),
                    'f', 2, 64,
                ),
            })
        }

        m.peerTable.SetRows(rows)

        return m, m.WaitForPeers()
    }

    var cmd tea.Cmd
    m.peerTable, cmd = m.peerTable.Update(msg)

    return m, cmd
}
