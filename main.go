package main

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
	"github.com/charmbracelet/wish/activeterm"
	"github.com/charmbracelet/wish/bubbletea"
	"github.com/charmbracelet/wish/logging"
	"github.com/tidwall/buntdb"
)

const (
	host = "0.0.0.0"
	port = "22"
)

// var names = make(Names)
// type Names map[string](string)

type Memory struct {
	names map[string]string
	board map[int64]map[string][]Try
}

func (m Memory) GetDay(day int64) map[string][]Try {
	players, ok := m.board[day]
	if !ok {
		m.board[day] = map[string][]Try{}
		players = m.board[day]
	}
	return players
}

func (m *Memory) AppendTry(day int64, playerid string, try Try) {
	players := m.GetDay(day)
	players[playerid] = append(players[playerid], try)
	m.board[day] = players
}

var memory = &Memory{names: map[string]string{}, board: map[int64]map[string][]Try{}}

type state struct {
	day       int64
	secret    string
	playerid  string
	height    int
	width     int
	gameState GameState
	screen    Screen
	styles    Styles
}

type GameState string

const (
	Idle    GameState = "0"
	Invalid GameState = "9"
	Win     GameState = "10"
)

type Screen string

const (
	TitleScreen Screen = "back to title"
	PlayScreen  Screen = "play today!"
	BoardScreen Screen = "see leaderboard"
)

func main() {
	db, err := buntdb.Open("data.json")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	s, err := wish.NewServer(
		wish.WithAddress(net.JoinHostPort(host, port)),
		wish.WithHostKeyPath(".ssh/id_ed25519"),
		wish.WithMiddleware(
			bubbletea.Middleware(func(s ssh.Session) (tea.Model, []tea.ProgramOption) {
				pty, _, _ := s.Pty()

				day := day()
				secret := secret(day)
				playerId := strings.Split(s.RemoteAddr().String(), ":")[0]

				renderer := bubbletea.MakeRenderer(s)

				state := state{
					day:       day,
					secret:    secret,
					playerid:  playerId,
					height:    pty.Window.Height,
					width:     pty.Window.Width,
					gameState: Idle,
					screen:    TitleScreen,
					styles:    Styles{}.New(renderer, secret),
				}

				m := Model{state: &state}.New()

				return m, []tea.ProgramOption{tea.WithAltScreen()}
			}),
			activeterm.Middleware(),
			logging.Middleware(),
		),
	)
	if err != nil {
		log.Error("Could not start server", "error", err)
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	log.Info("Starting SSH server", "host", host, "port", port)
	go func() {
		if err = s.ListenAndServe(); err != nil && !errors.Is(err, ssh.ErrServerClosed) {
			log.Error("Could not start server", "error", err)
			done <- nil
		}
	}()

	<-done
	log.Info("Stopping SSH server")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer func() { cancel() }()
	if err := s.Shutdown(ctx); err != nil && !errors.Is(err, ssh.ErrServerClosed) {
		log.Error("Could not stop server", "error", err)
	}
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
