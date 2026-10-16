package ui

import (

	tea "charm.land/bubbletea/v2"
)


func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {

    case tea.KeyPressMsg:
        switch msg.String() {

        case "ctrl+c", "q":
            return m, tea.Quit

        case "up", "k":
            if m.cursor > 0 {
                m.cursor--
            }

        case "down", "j":
            if m.cursor < len(m.choices)-1 {
                m.cursor++
            }

		case "esc":
			m.screen = ScreenMenu
        case "enter":
			switch m.cursor {
			case 0:
				m.screen = ScreenPeers
                // Discover Peers
            case 1:
                // List Peers
            case 2:
                return m, tea.Quit
            }

		}
	}
	if m.screen == ScreenPeers {
        model, c := m.discovery.Update(msg)
		m.discovery = model
        return m, c
    }

	return m, nil
}
