package ui

import  tea "charm.land/bubbletea/v2"


func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {

    case tea.KeyPressMsg:
        switch msg.String() {

        case "ctrl+c", "q":
            return m, tea.Quit

        // The "up" and "k" keys move the cursor up
        case "up", "k":
            if m.cursor > 0 {
                m.cursor--
            }

        // The "down" and "j" keys move the cursor down
        case "down", "j":
            if m.cursor < len(m.choices)-1 {
                m.cursor++
            }

        // The "enter" key and the space bar toggle the selected state
        // for the item that the cursor is pointing at.
        case "enter":
			switch m.cursor {
			case 0:
                // Discover Peers
            case 1:
                // List Peers
            case 2:
                return m, tea.Quit
            }

		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}
