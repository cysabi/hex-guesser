package main

import (
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type Board struct {
	state *state
	Table table.Model
}

func (m Board) New() Board {
	columns := []table.Column{
		{Title: "name", Width: 20},
		{Title: m.state.styles.TableRight.Render("guesses"), Width: 53},
	}
	m.Table = table.New(
		table.WithColumns(columns),
		table.WithHeight(m.state.height-8),
		table.WithFocused(false),
		table.WithStyles(table.Styles{
			Header:   m.state.styles.Subtitle.PaddingBottom(1),
			Cell:     m.state.styles.NormalText,
			Selected: m.state.styles.NormalText,
		}),
	)
	m.refreshRows()
	return m
}

func (m Board) Init() tea.Cmd {
	return textinput.Blink
}

func (m Board) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyLeft:
			m.state.dayPage -= 1
			m.refreshRows()
		case tea.KeyRight:
			m.state.dayPage += 1
			m.refreshRows()
		}
	}

	m.Table, cmd = m.Table.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m Board) View() string {
	return m.Table.View()
}

func (m *Board) refreshRows() {
	// find all of day:**:done = true
	// for each p find name
	// for each p day find all their tries

	// cap at showing the first 25 guesses
	// cap at showing the number 99

	// display a table of it
	rows := []table.Row{
		{"claire", m.state.styles.TableRight.Render("9")},
		// {"claire", m.state.styles.TableRight.Render("2139123981239 12")},
		// {"claire", m.state.styles.TableRight.Render("111111111122222222223333333333 12")},
		// {"claire", m.state.styles.TableRight.Render("00%%$$##((!!**@@))&&^^%%$$SSOOVVMM  ##((DDOOPP##)) 12")},
	}
	m.Table.SetRows(rows)
}
