package game

import (
	tea "github.com/charmbracelet/bubbletea"
)

type (
	// The PlayerJoinedMessage is a tea.Msg implementation that is produced when a new player joins the game.
	PlayerJoinedMessage struct {
		// The player's number (1 or 2).
		Number int
		// The player's name.
		Name string
	}
)

// PlayerJoined returns a tea.Cmd implementation that produces a PlayerJoinedMessage for the given player.
func PlayerJoined(name string, number int) tea.Cmd {
	return func() tea.Msg {
		return PlayerJoinedMessage{
			Number: number,
			Name:   name,
		}
	}
}
