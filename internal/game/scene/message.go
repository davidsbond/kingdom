package scene

import (
	tea "github.com/charmbracelet/bubbletea"
)

type (
	// The ChangeMessage is a tea.Msg implementation that informs the game of a scene change.
	ChangeMessage struct {
		// The function that creates the new scene.
		To ChangeFunc
	}

	// A ChangeFunc is a function that can generate a new scene.
	ChangeFunc func(window tea.Model) tea.Model
)

func change(to ChangeFunc) tea.Cmd {
	return func() tea.Msg {
		return ChangeMessage{To: to}
	}
}
