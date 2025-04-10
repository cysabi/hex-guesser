package main

import (
	"fmt"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// frame stuff
const fps = 60

type frameMsg time.Time // repr a single animation frame

func animate() tea.Cmd {
	return tea.Tick(time.Second/fps, func(t time.Time) tea.Msg {
		return frameMsg(t)
	})
}

// model
type model struct {
	board  board
	width  int
	height int
}

func (m model) Init() tea.Cmd {
	return tea.Sequence(tea.ClearScreen, animate())
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		return m, tea.Quit

	case tea.WindowSizeMsg:
		m.board.init(60, 30)
		m.width = msg.Width
		m.height = msg.Height
		return m, nil

	case tea.MouseMsg:
		if !m.board.ready() {
			return m, nil
		}

		return m, nil

	case frameMsg:
		if !m.board.ready() {
			return m, nil
		}

		return m, animate()

	default:
		return m, nil
	}
}

func (m model) View() string {
	if !m.board.ready() {
		return ""
	}
	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center,
		lipgloss.NewStyle().Width(m.board.width()).Height(m.board.height()).Border(lipgloss.RoundedBorder()).Padding(0).Render(
			m.board.String()))
}

// main
func main() {
	m := model{}
	p := tea.NewProgram(&m, tea.WithAltScreen(), tea.WithMouseCellMotion())
	if _, err := p.Run(); err != nil {
		fmt.Println("Uh oh:", err)
		os.Exit(1)
	}
}
