package main

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type Board struct {
	state *state
}

func (m Board) New() Board {
	return m
}

func (m Board) Init() tea.Cmd {
	return textinput.Blink
}

func (m Board) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// form, cmd := m.Form.Update(msg)
	// if f, ok := form.(*huh.Form); ok {
	// 	m.Form = f
	// }

	// return m, cmd
	return m, nil
}

func (m Board) View() string {
	return "leaderboard\nclaire   5 turns"
}
