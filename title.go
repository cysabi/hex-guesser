package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

type Title struct {
	Form *huh.Form
}

func (m Title) New() Title {
	return Title{
		Form: huh.NewForm(
			huh.NewGroup(
				huh.NewSelect[Screen]().
					Key("screen").
					Options(huh.NewOptions(
						PlayScreen,
						NameScreen,
						LeaderboardScreen)...),
			),
		),
	}
}

func (m Title) Init() tea.Cmd {
	return m.Form.Init()
}

func (m Title) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	form, cmd := m.Form.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		m.Form = f
	}

	return m, cmd
}

func (m Title) View() string {
	return m.Form.View()
}
