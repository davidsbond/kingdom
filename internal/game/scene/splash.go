package scene

import (
	tea "github.com/charmbracelet/bubbletea"

	"github.com/davidsbond/kingdom/internal/game/component/layout"
	"github.com/davidsbond/kingdom/internal/game/component/text"
)

func Splash() tea.Model {
	return create(
		layout.Centered(
			text.Text("hello world"),
		),
	)
}
