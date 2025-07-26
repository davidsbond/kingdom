package scene

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/davidsbond/kingdom/internal/game/component/image"
	"github.com/davidsbond/kingdom/internal/game/component/layout"
	"github.com/davidsbond/kingdom/internal/game/component/timing"
)

// Splash returns a tea.Model implementation describing the splash screen scene.
func Splash(window tea.Model) tea.Model {
	return create(
		window,
		timing.After(time.Second*5, change(Splash)),
		layout.Centered(
			image.Image("logo.txt"),
		),
	)
}
