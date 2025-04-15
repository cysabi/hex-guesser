package main

import (
	"encoding/hex"
	"math/rand"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
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
	wsize    tea.WindowSizeMsg
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
	random := rand.New(rand.NewSource(day))
	secret := make([]byte, 3)
	random.Read(secret)
	return hex.EncodeToString(secret)
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.wsize = msg
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
		return m.Title.View()
	case PlayScreen:
		return m.Game.View()
	case NameScreen:
		return m.Name.View()
	case LeaderboardScreen:
		return "leaderboard"
	}
	return ""
}
