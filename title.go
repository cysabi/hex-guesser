package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

type Title struct {
	form *huh.Form // huh.Form is just a tea.Model
}

func (t Title) New() Title {
	return Title{
		form: huh.NewForm(
			huh.NewGroup(
				huh.NewSelect[string]().
					Key("class").
					Options(huh.NewOptions("Warrior", "Mage", "Rogue")...).
					Title("Choose your class"),

				huh.NewSelect[int]().
					Key("level").
					Options(huh.NewOptions(1, 20, 9999)...).
					Title("Choose your level"),
			),
		),
	}
}

func (m Title) Init() tea.Cmd {
	return m.form.Init()
}

func (m Title) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	form, cmd := m.form.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		m.form = f
	}

	return m, cmd
}

func (m Title) View() string {
	return m.form.View()
}
