package main

import (
	"github.com/charmbracelet/lipgloss"
)

// styles

type Styles struct {
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
		GameBox:         ren.NewStyle().Padding(1, 0),
		ColorBox:        ren.NewStyle().Width(2).Height(1).Margin(0, 1),
		InputBox:        ren.NewStyle().Border(lipgloss.RoundedBorder()),
		StateMessageBox: ren.NewStyle().MarginLeft(1).Italic(false),
		TriesBox:        ren.NewStyle().Padding(0, 2, 0, 1),
		TryBox:          ren.NewStyle().MarginTop(1),
		CharGrade:       ren.NewStyle(),
	}
}
