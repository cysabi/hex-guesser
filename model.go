package main

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	state *state

	// screens
	Title Title
	Play  Play
	Board Board
}

func (m Model) New() Model {
	m.Title = Title{state: m.state}.New()
	m.Play = Play{state: m.state}
	m.Board = Board{state: m.state}
	return m
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(m.Title.Init(), textinput.Blink)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.state.height = msg.Height
		m.state.width = msg.Width
		return m, nil

	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}
	}

	switch m.state.screen {
	case TitleScreen:
		title, cmd := m.Title.Update(msg)
		if t, ok := title.(Title); ok {
			m.Title = t
		}
		cmds = append(cmds, cmd)

		if m.Title.Form.State == huh.StateCompleted {

			m.state.screen = m.Title.Form.Get("screen").(Screen)

			m.state.SetName(m.Title.Form.Get("name").(string))

			if m.state.screen == PlayScreen {
				m.Play = m.Play.New()
			}

			if m.state.screen == BoardScreen {
				m.Board = m.Board.New()
			}
		}

	case PlayScreen:
		game, cmd := m.Play.Update(msg)
		if g, ok := game.(Play); ok {
			m.Play = g
		}
		cmds = append(cmds, cmd)

	case BoardScreen:
		board, cmd := m.Board.Update(msg)
		if b, ok := board.(Board); ok {
			m.Board = b
		}
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {

	subtitle := m.state.styles.Subtitle.Render("day " + fmt.Sprint(m.state.day))
	if m.state.gameState != Idle {
		subtitle = m.state.styles.Subtitle.Foreground(lipgloss.Color(m.state.gameState)).Render(m.Play.StateMsg())
	}

	banner := m.state.styles.CharGrade.Margin(2).AlignHorizontal(lipgloss.Center).Render(
		lipgloss.JoinVertical(0.5,
			m.state.styles.Title.Foreground(lipgloss.Color("#"+m.state.secret)).AlignHorizontal(lipgloss.Center).Render("dailyhex!"),
			subtitle,
		),
	)
	switch m.state.screen {
	case TitleScreen:
		return lipgloss.Place(m.state.width, m.state.height, lipgloss.Center, lipgloss.Top,
			lipgloss.JoinVertical(0.5, banner, m.Title.View()))
	case PlayScreen:
		return lipgloss.Place(m.state.width, m.state.height, lipgloss.Center, lipgloss.Top,
			lipgloss.JoinVertical(0.5, banner, m.Play.View()))
	case BoardScreen:
		return lipgloss.Place(m.state.width, m.state.height, lipgloss.Center, lipgloss.Top,
			lipgloss.JoinVertical(0.5, banner, m.Board.View()))
	}
	return "uh oh"
}
