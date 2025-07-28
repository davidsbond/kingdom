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

type (
	// The KingdomSelectionChangedMessage is a tea.Msg implementation that is produced when a player modifies their
	// desired kingdom. This is mainly used to display changes of selection within the lobby scene across clients.
	KingdomSelectionChangedMessage struct {
		// The Player that changed their selection.
		Player int
		// The Kingdom the player selected.
		Kingdom int

		// Boolean to prevent infinite loops. Should be set to true once the message has been propagated to another
		// client.
		handled bool
	}
)

// KingdomSelectionChanged returns a tea.Cmd implementation that produces a KingdomSelectionChangedMessage for the given player.
func KingdomSelectionChanged(player, kingdom int) tea.Cmd {
	return func() tea.Msg {
		return KingdomSelectionChangedMessage{
			Player:  player,
			Kingdom: kingdom,
		}
	}
}
