package main

import (
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

// styles

type Styles struct {
	Title         lipgloss.Style
	Subtitle      lipgloss.Style
	FormBox       lipgloss.Style
	NormalText    lipgloss.Style
	TableRight    lipgloss.Style
	BoardSubtitle lipgloss.Style
	BoardArrows   lipgloss.Style
	GameBox       lipgloss.Style
	ColorBox      lipgloss.Style
	InputBox      lipgloss.Style
	MoveBox       lipgloss.Style
	CharGrade     lipgloss.Style
	Disabled      lipgloss.Style
	FormError     lipgloss.Style
	FormTheme     *huh.Theme
}

func (s Styles) New(r *lipgloss.Renderer, secret string) Styles {
	return Styles{
		Title:         r.NewStyle().Width(23).AlignHorizontal(lipgloss.Center).Bold(true),
		Subtitle:      r.NewStyle().Width(23).AlignHorizontal(lipgloss.Center).Foreground(lipgloss.Color("8")),
		NormalText:    r.NewStyle().Foreground(lipgloss.Color("7")),
		TableRight:    r.NewStyle().Width(53).AlignHorizontal(lipgloss.Right),
		BoardSubtitle: r.NewStyle().Foreground(lipgloss.Color("8")),
		BoardArrows:   r.NewStyle().Foreground(lipgloss.Color("7")).Bold(true),
		GameBox:       r.NewStyle().Width(23),
		ColorBox:      r.NewStyle().Width(2).Height(1).Margin(0, 1),
		InputBox:      r.NewStyle().Border(lipgloss.RoundedBorder()),
		MoveBox:       r.NewStyle().PaddingTop(1),
		CharGrade:     r.NewStyle(),
		Disabled:      r.NewStyle().Strikethrough(true).Foreground(lipgloss.Color("8")),
		FormBox:       r.NewStyle().Width(23).PaddingLeft(2),
		FormError:     r.NewStyle().Width(21).PaddingLeft(1).Foreground(lipgloss.Color("1")),
		FormTheme:     makeFormTheme(r, secret),
	}
}

// 0 8 7 secret
// bold

func makeFormTheme(r *lipgloss.Renderer, secret string) *huh.Theme {
	var t huh.Theme

	t.FieldSeparator = r.NewStyle().SetString("\n\n\n")

	// group
	t.Blurred.Base = r.NewStyle().BorderForeground(lipgloss.Color("0")).BorderStyle(lipgloss.HiddenBorder()).BorderLeft(true)

	// prompts
	t.Blurred.SelectSelector = r.NewStyle().Foreground(lipgloss.Color("8")).Bold(true).SetString("> ")
	t.Blurred.TextInput.Prompt = r.NewStyle().Foreground(lipgloss.Color("8")).Bold(true)

	// text
	t.Blurred.UnselectedOption = r.NewStyle().Foreground(lipgloss.Color("8"))
	t.Blurred.SelectedOption = r.NewStyle().Foreground(lipgloss.Color("8"))
	t.Blurred.TextInput.Text = r.NewStyle().Foreground(lipgloss.Color("8"))
	t.Blurred.TextInput.Placeholder = t.Blurred.TextInput.Text.Foreground(lipgloss.Color("0"))

	// ~ FOCUSED ~
	t.Focused = t.Blurred

	// prompts
	t.Focused.SelectSelector = t.Focused.SelectSelector.Foreground(lipgloss.Color("7"))
	t.Focused.TextInput.Prompt = t.Focused.TextInput.Prompt.Foreground(lipgloss.Color("7"))

	// text
	t.Focused.UnselectedOption = r.NewStyle().Foreground(lipgloss.Color("7"))
	t.Focused.SelectedOption = r.NewStyle().Foreground(lipgloss.Color("#" + secret))
	t.Focused.TextInput.Text = r.NewStyle().Foreground(lipgloss.Color("#" + secret))

	return &t
}
