package scene

import (
	tea "github.com/charmbracelet/bubbletea"

	"github.com/davidsbond/kingdom/internal/game/component/layout"
	"github.com/davidsbond/kingdom/internal/game/component/text"
	"github.com/davidsbond/kingdom/internal/game/entity"
)

// Lobby returns a tea.Model implementation describing the player lobby scene.
func Lobby(ctx Context) tea.Model {
	ctx.Logger.With("scene", "lobby").Debug("initialising scene")

	return create(ctx,
		layout.Centered(
			layout.Horizontal(
				entity.LobbyPlayerTitle(ctx.Logger, ctx.State, 1),
				text.Text("VS"),
				entity.LobbyPlayerTitle(ctx.Logger, ctx.State, 2),
			),
		),
	)
}
