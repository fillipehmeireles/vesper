package discovery


func (m Model) Render() string {
	return m.peerTable.View()
}
