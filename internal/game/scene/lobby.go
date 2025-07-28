package scene

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss/v2"

	"github.com/davidsbond/kingdom/internal/game/component/layout"
	"github.com/davidsbond/kingdom/internal/game/component/text"
	"github.com/davidsbond/kingdom/internal/game/entity"
)

// Lobby returns a tea.Model implementation describing the player lobby scene.
func Lobby(ctx Context) tea.Model {
	logger := ctx.Logger.With("scene", "lobby")
	logger.Debug("initialising scene")

	return create(ctx,
		layout.Centered(
			layout.Vertical(
				text.Text("KINGDOM SELECTION", text.Width(96), text.Align(lipgloss.Center)),
				layout.Horizontal(
					layout.Vertical(
						entity.LobbyPlayerTitle(logger, ctx.State, 1),
						entity.LobbyKingdomSelector(logger, ctx.Player.Number() == 1, 1),
					),
					text.Text("VS"),
					layout.Vertical(
						entity.LobbyPlayerTitle(logger, ctx.State, 2),
						entity.LobbyKingdomSelector(logger, ctx.Player.Number() == 2, 2),
					),
				),
			),
		),
	)
}
