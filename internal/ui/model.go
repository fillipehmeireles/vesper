package ui

import  tea "charm.land/bubbletea/v2"

type Screen int

const (
    ScreenMenu Screen = iota
    ScreenPeers
)

type ModelNodeData struct {
	NodeName string
	NodeAddr string
}

type Model struct {
	screen Screen
    choices  []string
    cursor   int
	nodeData ModelNodeData
}

func InitialModel(modelNodeData ModelNodeData) Model {
	return Model { choices: []string{
		"Discover Peers"
		"Test: Say 'Hi' to a peer",
	},
		nodeData: modelNodeData,
	}
}

func (m Model) Init() tea.Cmd {
    return nil
}

type VesperUIModel interface {
	SetupNodeData()
}
