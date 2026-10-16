package discovery

import tea "charm.land/bubbletea/v2"

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
    var cmd tea.Cmd

    m.peerTable, cmd = m.peerTable.Update(msg)

    return m, cmd
}
