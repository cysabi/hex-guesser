package main

import (
	"encoding/hex"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func createInput() textinput.Model {
	ti := textinput.New()
	ti.CharLimit = 6
	ti.Width = 6
	ti.Prompt = ""
	ti.Focus()
	return ti
}

func filterHex(s string) string {
	var result []rune
	for _, r := range s {
		if (r >= '0' && r <= '9') || (r >= 'a' && r <= 'f') || (r >= 'A' && r <= 'F') {
			result = append(result, r)
		}
	}
	return string(result)
}

func dailySecret() string {
	random := rand.New(rand.NewSource((time.Now().Unix() / (60 * 60 * 24))))

	secret := make([]byte, 3)

	random.Read(secret)
	return hex.EncodeToString(secret)
}

func stateToStyle(m model) (string, string) {
	if m.state == Invalid {
		return "9", "invalid hex!"
	} else if m.state == Win {
		return "10", fmt.Sprintf("you got it! (%d turns)", len(m.board.tries)+1)
	} else {
		return "0", ""
	}
}

// model

type model struct {
	state     State
	textInput textinput.Model
	board     board
	wsize     tea.WindowSizeMsg
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if m.state == Win {
			return m, tea.Quit
		}
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit

		case tea.KeyEnter:

			m.state = m.board.submit(m.textInput.Value())

			if m.state == Win {
				m.textInput.Blur()
			}
			if m.state == Idle {
				m.textInput.SetValue("")
			}
		default:
			if m.state == Invalid {
				m.state = Idle
			}
		}

	case tea.WindowSizeMsg:
		m.wsize = msg
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)
	m.textInput.SetValue(strings.ToLower(filterHex(m.textInput.Value())))

	return m, cmd
}

func (m model) View() string {
	stateAnsi, stateMsg := stateToStyle(m)
	return lipgloss.NewStyle().Padding(1, 0).Render(lipgloss.JoinVertical(0,
		lipgloss.JoinHorizontal(lipgloss.Center,
			lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color(stateAnsi)).Render(
				lipgloss.JoinHorizontal(0,
					colorBoxStyle.Background(lipgloss.Color("#"+m.board.secret)).Render(),
					m.textInput.View(),
				),
			),
			lipgloss.NewStyle().Foreground(lipgloss.Color(stateAnsi)).MarginLeft(1).Italic(false).Render(stateMsg),
		),
		m.board.View(),
	))
}

// main
func main() {
	m := model{
		state:     Idle,
		textInput: createInput(),
	}
	m.board.secret = dailySecret()

	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Println("Uh oh:", err)
		os.Exit(1)
	}
}
