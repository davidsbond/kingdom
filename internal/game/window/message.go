package window

import (
	tea "github.com/charmbracelet/bubbletea"
)

type (
	SizeMessage struct {
		Width  int
		Height int
	}

	GetSizeMessage struct{}
)

func size(w, h int) tea.Cmd {
	return func() tea.Msg {
		return SizeMessage{
			Width:  w,
			Height: h,
		}
	}
}

func Size() tea.Cmd {
	return func() tea.Msg {
		return GetSizeMessage{}
	}
}
