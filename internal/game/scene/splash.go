package scene

import (
	"runtime/debug"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss/v2"

	"github.com/davidsbond/kingdom/internal/game"
	"github.com/davidsbond/kingdom/internal/game/component/image"
	"github.com/davidsbond/kingdom/internal/game/component/layout"
	"github.com/davidsbond/kingdom/internal/game/component/text"
	"github.com/davidsbond/kingdom/internal/game/component/timing"
)

// Splash returns a tea.Model implementation describing the splash screen scene.
func Splash(ctx Context) tea.Model {
	logger := ctx.Logger.With("scene", "splash")

	version := "Development Version"
	if info, ok := debug.ReadBuildInfo(); ok && info.Main.Version != "" {
		version = info.Main.Version
	}

	const (
		duration  = time.Second * 5
		logoWidth = 59
		logoName  = "logo.txt"
	)

	logger.Debug("initialising scene")

	return create(ctx,
		timing.After(duration, tea.Sequence(
			change(Lobby),
			game.PlayerJoined(ctx.Player.Name(), ctx.Player.Number()),
		)),
		layout.Centered(
			layout.Vertical(
				image.Image(logger, logoName, image.Foreground(lipgloss.Red)),
				text.Text(version,
					text.Width(logoWidth),
					text.Align(lipgloss.Center),
				),
			),
		),
	)
}
