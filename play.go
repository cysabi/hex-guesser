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
				m.state.SetDone(true)
				m.state.gameState = Win
			} else {
				m.state.AppendMove(move)
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
	return m.state.styles.GameBox.Render(lipgloss.JoinVertical(0.5,
		m.state.styles.InputBox.BorderForeground(lipgloss.Color(m.state.gameState)).Render(
			lipgloss.JoinHorizontal(0,
				m.state.styles.ColorBox.Background(lipgloss.Color("#"+m.state.secret)).Render(),
				m.Input.View(),
			),
		),
		m.state.styles.GameBox.Render(lipgloss.JoinVertical(0,
			m.displayMoves()...,
		)),
	))
}

func (m Play) StateMsg() string {
	if m.state.gameState == Invalid {
		return "invalid hex"
	} else if m.state.gameState == Win {
		return fmt.Sprintf("you got it! (%d turns)", len(m.state.GetMoves())+1)
	}
	return ""
}

type CharGrade string

const (
	Green  CharGrade = "2"
	Yellow CharGrade = "3"
	Gray   CharGrade = "8"
)

func (m Play) displayMoves() []string {
	moves := m.state.GetMoves()
	slices.Reverse(moves)
	out := make([]string, len(moves))

	for i, move := range moves {
		grade := m.gradeMove(move)
		out[i] = m.displayMove(move, grade)
	}

	return out

}

func (m Play) gradeMove(move string) []CharGrade {
	grade := make([]CharGrade, len(m.state.secret))
	secret := []rune(m.state.secret)

	for i, s := range secret {
		if s == []rune(move)[i] {
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
		for i, m := range move {
			if m == s && grade[i] == Gray {
				grade[i] = Yellow
				break
			}
		}
	}
	return grade
}

func (m Play) displayMove(move string, grade []CharGrade) string {
	var text strings.Builder
	for index, letter := range move {
		str := m.state.styles.CharGrade.Foreground(lipgloss.Color(grade[index])).Render(string(letter))
		text.WriteString(str)
	}

	return m.state.styles.MoveBox.Render(
		lipgloss.JoinHorizontal(0,
			m.state.styles.ColorBox.Background(lipgloss.Color("#"+move)).Render(),
			text.String(),
		))
}
