package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

var (
	Program *tea.Program
)

type MainModel struct {
	Quitting    bool
	initialized bool
}

func New() MainModel {
	return MainModel{
		initialized: true,
	}
}

func (m MainModel) Update(untypedMessage tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch untypedMessage.(type) {
	case tea.KeyMsg:
		m.Quitting = true
	}

	if m.Quitting {
		cmds = append(cmds, tea.Quit)
	}

	return m, tea.Batch(cmds...)
}

func (m MainModel) View() string {
	return "The program is running, press any key to quit"
}

func (m MainModel) Init() tea.Cmd {
	if !m.initialized {
		panic("Main model was not initialized")
	}

	return nil
}
