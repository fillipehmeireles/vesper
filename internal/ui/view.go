package ui

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
)

func (m Model) View() tea.View {
	var s strings.Builder

    fmt.Fprintf(&s, "Vesper\n%s @ %s\n\n", m.nodeData.NodeName, m.nodeData.NodeAddr)
	for i, choice := range m.choices {
        cursor := " "
        if m.cursor == i {
            cursor = ">"
        }

        fmt.Fprintf(&s, "%s %s\n", cursor, choice)
    }

    // The footer
	fmt.Fprintf(&s, "\nPress q to quit.\n")

    // Send the UI for rendering
    return tea.NewView(s.String())
}
