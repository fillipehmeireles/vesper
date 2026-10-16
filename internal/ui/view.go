package ui

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
)


func (m Model) renderMenu() string {
	var s strings.Builder

    fmt.Fprintf(&s, "Vesper\n%s @ %s\n\n", m.nodeData.NodeName, m.nodeData.NodeAddr)
	for i, choice := range m.choices {
        cursor := " "
        if m.cursor == i {
            cursor = ">"
        }

        fmt.Fprintf(&s, "%s %s\n", cursor, choice)
    }

	fmt.Fprintf(&s, "\nPress q to quit.\n")

	return s.String()
}



func (m Model) View() tea.View {
	var content string

	switch m.screen {
	case ScreenMenu:
		content = m.renderMenu()
	case ScreenPeers:
		content = m.discovery.Render()
	}

	return tea.NewView(content)
}
