package message

import (
	tea "github.com/charmbracelet/bubbletea"

	"github.com/davidsbond/kingdom/internal/game/component"
)

type (
	initialise struct {
		component.NoUpdate
		component.NoView

		command tea.Cmd
	}
)

// Init returns a tea.Model implementation that can fire off arbitrary tea.Cmd implementations as part of model
// initialisation.
func Init(command tea.Cmd) tea.Model {
	return &initialise{command: command}
}

func (i *initialise) Init() tea.Cmd {
	return i.command
}
