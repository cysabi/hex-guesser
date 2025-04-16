package main

import (
	"errors"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

type Title struct {
	state *state

	Form *huh.Form
}

func (m Title) New() Title {
	placeholder := "your name"
	if len(m.state.playerid) > 0 {
		placeholder = m.state.playerid
	}
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().Key("name").Placeholder(placeholder).Prompt("? ").Validate(
				func(str string) error {
					if len(str) == 0 {
						return errors.New("what's ur name?")
					}
					return nil
				},
			),
			huh.NewSelect[Screen]().
				Key("screen").
				Options(huh.NewOptions(
					PlayScreen,
					BoardScreen)...),
		),
	).WithWidth(19).WithShowHelp(false).WithTheme(m.state.styles.FormTheme)

	return Title{Form: form}
}

func (m Title) Init() tea.Cmd {
	return tea.Batch(m.Form.Init(), textinput.Blink)
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
