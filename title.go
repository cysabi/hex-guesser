package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

type Title struct {
	Value string
	Form  *huh.Form
}

func (m Title) New() Title {
	input := huh.NewInput().Key("name").Title("enter name").Placeholder(m.Value)
	return Title{
		Form: huh.NewForm(
			huh.NewGroup(
				input,
				huh.NewSelect[Screen]().
					Key("screen").
					Options(huh.NewOptions(
						PlayScreen,
						NameScreen,
						LeaderboardScreen)...),
			),
		).WithWidth(20).WithTheme(huh.ThemeBase16()).WithShowHelp(false),
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
