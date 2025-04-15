package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

type Name struct {
	Value string
	Form  *huh.Form // huh.Form is just a tea.Model
}

func (m Name) New() Name {
	input := huh.NewInput().Key("name").Description("enter name").Placeholder(m.Value)
	input.Focus()
	return Name{
		Form: huh.NewForm(
			huh.NewGroup(
				input,
			),
		),
	}
}

func (m Name) Init() tea.Cmd {
	return m.Form.Init()
}

func (m Name) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	form, cmd := m.Form.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		m.Form = f
	}

	return m, cmd
}

func (m Name) View() string {
	return m.Form.View()
}
