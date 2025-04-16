package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

type Screen string

const (
	TitleScreen       Screen = "back to title"
	PlayScreen        Screen = "play today!"
	NameScreen        Screen = "change name"
	LeaderboardScreen Screen = "see leaderboard"
)

type model struct {
	PlayerId string
	State    Screen
	Day      int64
	Title    Title
	Game     Game
	Name     Name
	Styles   Styles
	Height   int
	Width    int
}

func (m model) New() model {
	m.State = TitleScreen
	m.Day = day()
	m.Title = Title{}.New()
	m.Game = Game{
		Secret:   secret(m.Day),
		Day:      m.Day,
		PlayerId: m.PlayerId,
		Styles:   m.Styles,
	}
	m.Name = Name{}
	return m
}

func day() int64 {
	loc, _ := time.LoadLocation("America/New_York")
	now := time.Now().In(loc)

	adjusted := now.Add(-11 * time.Hour)

	dayNumber := adjusted.Unix() / (60 * 60 * 24)
	return dayNumber
}

func secret(day int64) string {
	input := []byte("secret" + fmt.Sprint(day))
	hash := sha256.Sum256(input)
	return hex.EncodeToString(hash[:3])
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Height = msg.Height
		m.Width = msg.Width
		return m, nil

	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}
	}

	switch m.State {
	case TitleScreen:
		title, cmd := m.Title.Update(msg)
		if t, ok := title.(Title); ok {
			m.Title = t
		}
		cmds = append(cmds, cmd)

		if m.Title.Form.State == huh.StateCompleted {
			m.State = m.Title.Form.Get("screen").(Screen)
			if m.State == PlayScreen {
				m.Game = m.Game.New()
			}
			if m.State == NameScreen {
				m.Name = m.Name.New()
			}
		}

	case NameScreen:
		name, cmd := m.Name.Update(msg)
		if t, ok := name.(Name); ok {
			m.Name = t
		}
		cmds = append(cmds, cmd)

		if m.Name.Form.State == huh.StateCompleted {
			m.Name.Value = m.Name.Form.Get("name").(string)
			m.State = TitleScreen
			m.Title = Title{}.New()
		}

	case PlayScreen:
		game, cmd := m.Game.Update(msg)
		if g, ok := game.(Game); ok {
			m.Game = g
		}
		cmds = append(cmds, cmd)

	case LeaderboardScreen:
	}

	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	switch m.State {
	case TitleScreen:
		return lipgloss.Place(m.Width, m.Height, lipgloss.Center, lipgloss.Top,
			lipgloss.JoinVertical(0.5,
				m.Styles.CharGrade.MarginTop(2).Render(),
				m.Styles.Title.Render("dailyhex!"),
				m.Styles.Subtitle.Render("day "+fmt.Sprint(m.Day)),
				m.Title.View(),
			),
		)
	case PlayScreen:
		return lipgloss.Place(m.Width, m.Height, lipgloss.Center, lipgloss.Top,
			lipgloss.JoinVertical(0.5,
				m.Styles.CharGrade.MarginTop(2).Render(),
				m.Styles.Title.Render("dailyhex!"),
				m.Styles.Subtitle.Render("day "+fmt.Sprint(m.Day)),
				m.Game.View(),
			),
		)
	case NameScreen:
		return m.Name.View()
	case LeaderboardScreen:
		return "leaderboard"
	}
	return ""
}
