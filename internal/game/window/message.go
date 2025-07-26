package window

import (
	tea "github.com/charmbracelet/bubbletea"
)

type (
	// The SizeMessage type is a tea.Msg implementation that describes the window dimensions for a single client.
	SizeMessage struct {
		// The client's window width.
		Width int
		// The client's window height.
		Height int
	}

	// The GetSizeMessage type is a tea.Msg implementation used to trigger a SizeMessage from a Window instance.
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

// Size returns a tea.Cmd that triggers a SizeMessage from a Window instance.
func Size() tea.Cmd {
	return func() tea.Msg {
		return GetSizeMessage{}
	}
}
