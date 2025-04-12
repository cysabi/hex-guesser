package main

import (
	"slices"

	"github.com/charmbracelet/lipgloss"
)

type board struct {
	secret string
	tries  []string
}

type State int

const (
	Idle State = iota
	Invalid
	Win
)

func (c *board) submit(move string) State {
	if len(move) != 6 {
		return Invalid
	}
	if c.secret == move {
		return Win
	}
	grade := c.gradeMove(move)
	c.tries = append(c.tries, gradeDisplay(grade, move))
	return Idle
}

type Grade string

const (
	Green  = "2"
	Yellow = "3"
	Gray   = "8"
)

func (c board) gradeMove(move string) []string {
	grade := make([]string, len(move))
	secret := []rune(c.secret)

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

func gradeDisplay(grade []string, move string) string {
	display := ""

	for i, g := range grade {
		display = display + lipgloss.NewStyle().Foreground(lipgloss.Color(g)).Render(string(move[i]))
	}

	return lipgloss.NewStyle().MarginTop(1).Render(
		lipgloss.JoinHorizontal(0,
			colorBoxStyle.Background(lipgloss.Color("#"+move)).Render(),
			display,
		))
}

var colorBoxStyle = lipgloss.NewStyle().Width(2).Height(1).Margin(0, 1)

func (c board) View() string {
	tries := slices.Clone(c.tries)
	slices.Reverse(tries)
	return lipgloss.NewStyle().Padding(0, 2, 0, 1).Render(
		lipgloss.JoinVertical(0,
			tries...,
		),
	)
}
