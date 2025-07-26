package scene

import (
	"runtime/debug"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss/v2"

	"github.com/davidsbond/kingdom/internal/game/component/image"
	"github.com/davidsbond/kingdom/internal/game/component/layout"
	"github.com/davidsbond/kingdom/internal/game/component/text"
	"github.com/davidsbond/kingdom/internal/game/component/timing"
)

// Splash returns a tea.Model implementation describing the splash screen scene.
func Splash(window tea.Model) tea.Model {
	version := "Development Version"
	if info, ok := debug.ReadBuildInfo(); ok && info.Main.Version != "" {
		version = info.Main.Version
	}

	const (
		duration  = time.Second * 5
		logoWidth = 59
		logoName  = "logo.txt"
	)

	return create(
		window,
		timing.After(duration, change(Splash)),
		layout.Centered(
			layout.Vertical(
				image.Image(logoName, image.Foreground(lipgloss.Red)),
				text.Text(version,
					text.Width(logoWidth),
					text.Align(lipgloss.Center),
				),
			),
		),
	)
}
