package main

import (
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

// styles

type Styles struct {
	Title           lipgloss.Style
	Subtitle        lipgloss.Style
	GameBox         lipgloss.Style
	ColorBox        lipgloss.Style
	InputBox        lipgloss.Style
	StateMessageBox lipgloss.Style
	TriesBox        lipgloss.Style
	TryBox          lipgloss.Style
	CharGrade       lipgloss.Style
}

func (s Styles) New(ren *lipgloss.Renderer) Styles {
	return Styles{
		Title:           ren.NewStyle().Foreground(lipgloss.Color("15")).Bold(true),
		Subtitle:        ren.NewStyle().Foreground(lipgloss.Color("08")),
		GameBox:         ren.NewStyle().Padding(1, 0, 1, 1),
		ColorBox:        ren.NewStyle().Width(2).Height(1).Margin(0, 1),
		InputBox:        ren.NewStyle().Border(lipgloss.RoundedBorder()),
		StateMessageBox: ren.NewStyle().MarginLeft(1).Italic(false),
		TriesBox:        ren.NewStyle().Padding(0, 2, 0, 1),
		TryBox:          ren.NewStyle().MarginTop(1),
		CharGrade:       ren.NewStyle(),
	}
}

func Theme() *huh.Theme {
	t := huh.ThemeBase()

	t.FieldSeparator = lipgloss.NewStyle().SetString("\n\n\n")
	t.Group.Width(16)

	// button := lipgloss.NewStyle().
	// 	Padding(buttonPaddingVertical, buttonPaddingHorizontal).
	// 	MarginRight(1)

	// // Focused styles.
	// t.Focused.Base = lipgloss.NewStyle().PaddingLeft(1).BorderStyle(lipgloss.ThickBorder()).BorderLeft(true)
	// t.Focused.Card = lipgloss.NewStyle().PaddingLeft(1)
	// t.Focused.ErrorIndicator = lipgloss.NewStyle().SetString(" *")
	// t.Focused.ErrorMessage = lipgloss.NewStyle().SetString(" *")
	t.Focused.SelectSelector = lipgloss.NewStyle()
	// t.Focused.NextIndicator = lipgloss.NewStyle().MarginLeft(1).SetString("→")
	// t.Focused.PrevIndicator = lipgloss.NewStyle().MarginRight(1).SetString("←")
	t.Focused.MultiSelectSelector = lipgloss.NewStyle()
	// t.Focused.SelectedPrefix = lipgloss.NewStyle().SetString("[•] ")
	// t.Focused.UnselectedPrefix = lipgloss.NewStyle().SetString("[ ] ")
	// t.Focused.FocusedButton = button.Foreground(lipgloss.Color("0")).Background(lipgloss.Color("7"))
	// t.Focused.BlurredButton = button.Foreground(lipgloss.Color("7")).Background(lipgloss.Color("0"))
	// t.Focused.TextInput.Placeholder = lipgloss.NewStyle().Foreground(lipgloss.Color("8"))

	// t.Help = help.New().Styles

	// // Blurred styles.
	// t.Blurred = t.Focused
	// t.Blurred.Base = t.Blurred.Base.BorderStyle(lipgloss.HiddenBorder())
	// t.Blurred.MultiSelectSelector = lipgloss.NewStyle().SetString("  ")
	// t.Blurred.NextIndicator = lipgloss.NewStyle()
	// t.Blurred.PrevIndicator = lipgloss.NewStyle()

	return t
}
