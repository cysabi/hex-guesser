package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Play struct {
	state *state

	Input textinput.Model
}

func (m Play) New() Play {
	ti := textinput.New()
	ti.CharLimit = 6
	ti.Width = 6
	ti.Prompt = ""
	ti.Focus()

	m.Input = ti
	m.state.gameState = Idle
	return m
}

func (m Play) Init() tea.Cmd {
	return tea.Batch(textinput.Blink)
}

func (m Play) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {

		case tea.KeyEnter:
			move := m.Input.Value()

			if len(move) != 6 {
				m.state.gameState = Invalid
			} else if m.state.secret == move {
				m.state.gameState = Win
			} else {
				memory.AppendTry(m.state.day, m.state.playerid, Try{move: move}.New(m.state.secret))
				m.state.gameState = Idle
			}

			if m.state.gameState == Win {
				m.Input.Blur()
			}
			if m.state.gameState == Idle {
				m.Input.SetValue("")
			}

		default:
			if m.state.gameState == Invalid {
				m.state.gameState = Idle
			}
		}
	}

	var cmd tea.Cmd
	m.Input, cmd = m.Input.Update(msg)

	var hex []rune
	for _, r := range m.Input.Value() {
		if (r >= '0' && r <= '9') || (r >= 'a' && r <= 'f') || (r >= 'A' && r <= 'F') {
			hex = append(hex, r)
		}
	}

	m.Input.SetValue(strings.ToLower(string(hex)))
	return m, cmd
}

func (m Play) View() string {
	tries := slices.Clone(memory.GetDay(m.state.day)[m.state.playerid])
	slices.Reverse(tries)

	display := make([]string, len(memory.GetDay(m.state.day)[m.state.playerid]))
	for i, t := range tries {
		display[i] = t.View(m.state.styles)
	}

	return m.state.styles.GameBox.Render(lipgloss.JoinVertical(0,
		lipgloss.JoinHorizontal(lipgloss.Center,
			m.state.styles.InputBox.BorderForeground(lipgloss.Color(m.state.gameState)).Render(
				lipgloss.JoinHorizontal(0,
					m.state.styles.ColorBox.Background(lipgloss.Color("#"+m.state.secret)).Render(),
					m.Input.View(),
				),
			),
			m.state.styles.StateMessageBox.Foreground(lipgloss.Color(m.state.gameState)).Render(m.StateMsg()),
		),
		m.state.styles.TriesBox.Render(
			lipgloss.JoinVertical(0,
				display...,
			),
		),
	))
}

func (m Play) StateMsg() string {
	if m.state.gameState == Invalid {
		return "invalid hex!"
	} else if m.state.gameState == Win {
		return fmt.Sprintf("you got it! (%d turns)", len(memory.GetDay(m.state.day)[m.state.playerid])+1)
	}
	return ""
}

type Try struct {
	move  string
	grade []Grade
}

type Grade string

const (
	Green  Grade = "2"
	Yellow Grade = "3"
	Gray   Grade = "8"
)

func (m Try) New(inSecret string) Try {
	grade := make([]Grade, len(inSecret))
	secret := []rune(inSecret)

	for i, s := range secret {
		if s == []rune(m.move)[i] {
			grade[i] = Green
			secret[i] = ' '
		} else {
			grade[i] = Gray
		}
	}

	for _, s := range secret {
		if s == ' ' {
			continue
		}
		for i, m := range m.move {
			if m == s && grade[i] == Gray {
				grade[i] = Yellow
				break
			}
		}
	}

	m.grade = grade
	return m
}

func (m Try) View(styles Styles) string {

	var display strings.Builder
	for index, letter := range m.move {
		str := styles.CharGrade.Foreground(lipgloss.Color(m.grade[index])).Render(string(letter))
		display.WriteString(str)
	}

	return styles.TryBox.Render(
		lipgloss.JoinHorizontal(0,
			styles.ColorBox.Background(lipgloss.Color("#"+m.move)).Render(),
			display.String(),
		))
}
