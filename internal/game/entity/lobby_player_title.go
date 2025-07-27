package entity

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss/v2"
	"github.com/charmbracelet/log"

	"github.com/davidsbond/kingdom/internal/game"
	"github.com/davidsbond/kingdom/internal/game/component/layout"
	"github.com/davidsbond/kingdom/internal/game/component/message"
	"github.com/davidsbond/kingdom/internal/game/component/text"
	"github.com/davidsbond/kingdom/internal/game/style/border"
)

// LobbyPlayerTitle returns a tea.Model implementation that displays the name of the player. If the player has yet
// to join it will display "WAITING..." and update via messaging once the player joins.
func LobbyPlayerTitle(logger *log.Logger, state *game.State, player int) tea.Model {
	id := fmt.Sprintf("lobby-player-title-%d", player)

	align := lipgloss.Left
	colour := lipgloss.Red
	if player == 2 {
		align = lipgloss.Right
		colour = lipgloss.Blue
	}

	const (
		width   = 36
		margin  = 1
		padding = 1
	)

	name := "WAITING..."
	if p := state.PlayerN(player); p != nil {
		name = p.Name()
	}

	logger = logger.With("entity", "lobby_player_title")

	return create(id,
		message.Handler(func(msg tea.Msg) tea.Cmd {
			if m, ok := msg.(game.PlayerJoinedMessage); ok && m.Number == player {
				logger.
					With("player_name", m.Name, "player_number", m.Number).
					Debug("detected player join")
				
				return text.Change(id, m.Name)
			}

			return nil
		}),

		layout.Vertical(
			layout.Container(
				text.Dynamic(id,
					text.Foreground(lipgloss.Black),
					text.Background(colour),
					text.Width(width),
					text.Align(align),
					text.Content(name),
					text.Padding(0, padding, 0),
				),
				layout.ContainerAlign(lipgloss.Center),
				layout.ContainerMargin(margin),
				layout.ContainerBorder(border.Wing(align), colour),
			),
		),
	)
}
