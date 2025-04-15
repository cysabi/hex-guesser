package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Game struct {
	Secret   string
	Day      int64
	PlayerId string
	input    textinput.Model
	State    GameState
	Styles   Styles
}

func (game Game) New() Game {
	ti := textinput.New()
	ti.CharLimit = 6
	ti.Width = 6
	ti.Prompt = ""
	ti.Focus()

	game.input = ti
	game.State = Idle
	return game
}

type GameState string

const (
	Idle    GameState = "0"
	Invalid GameState = "9"
	Win     GameState = "10"
)

func (game Game) Init() tea.Cmd {
	return nil
}

func (game Game) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {

		case tea.KeyEnter:
			move := game.input.Value()

			if len(move) != 6 {
				game.State = Invalid
			} else if game.Secret == move {
				game.State = Win
			} else {
				memory.AppendTry(game.Day, game.PlayerId, Try{move: move}.New(game.Secret))
				game.State = Idle
			}

			if game.State == Win {
				game.input.Blur()
			}
			if game.State == Idle {
				game.input.SetValue("")
			}

		default:
			if game.State == Invalid {
				game.State = Idle
			}
		}
	}

	var cmd tea.Cmd
	game.input, cmd = game.input.Update(msg)

	var hex []rune
	for _, r := range game.input.Value() {
		if (r >= '0' && r <= '9') || (r >= 'a' && r <= 'f') || (r >= 'A' && r <= 'F') {
			hex = append(hex, r)
		}
	}

	game.input.SetValue(strings.ToLower(string(hex)))
	return game, cmd
}

func (g Game) View() string {
	tries := slices.Clone(memory.GetDay(g.Day)[g.PlayerId])
	slices.Reverse(tries)

	display := make([]string, len(memory.GetDay(g.Day)[g.PlayerId]))
	for i, t := range tries {
		display[i] = t.View(g.Styles)
	}

	return g.Styles.GameBox.Render(lipgloss.JoinVertical(0,
		lipgloss.JoinHorizontal(lipgloss.Center,
			g.Styles.InputBox.BorderForeground(lipgloss.Color(g.State)).Render(
				lipgloss.JoinHorizontal(0,
					g.Styles.ColorBox.Background(lipgloss.Color("#"+g.Secret)).Render(),
					g.input.View(),
				),
			),
			g.Styles.StateMessageBox.Foreground(lipgloss.Color(g.State)).Render(g.StateMsg()),
		),
		g.Styles.TriesBox.Render(
			lipgloss.JoinVertical(0,
				display...,
			),
		),
	))
}

func (game Game) StateMsg() string {
	if game.State == Invalid {
		return "invalid hex!"
	} else if game.State == Win {
		return fmt.Sprintf("you got it! (%d turns)", len(memory.GetDay(game.Day)[game.PlayerId])+1)
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

func (t Try) New(inSecret string) Try {
	grade := make([]Grade, len(inSecret))
	secret := []rune(inSecret)

	for i, s := range secret {
		if s == []rune(t.move)[i] {
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
		for i, m := range t.move {
			if m == s && grade[i] == Gray {
				grade[i] = Yellow
				break
			}
		}
	}

	t.grade = grade
	return t
}

func (t Try) View(styles Styles) string {

	var display strings.Builder
	for index, letter := range t.move {
		str := styles.CharGrade.Foreground(lipgloss.Color(t.grade[index])).Render(string(letter))
		display.WriteString(str)
	}

	return styles.TryBox.Render(
		lipgloss.JoinHorizontal(0,
			styles.ColorBox.Background(lipgloss.Color("#"+t.move)).Render(),
			display.String(),
		))
}
