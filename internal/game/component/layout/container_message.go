package layout

import (
	"image/color"

	tea "github.com/charmbracelet/bubbletea"
)

type (
	// The ContainerBorderForegroundChangeMessage type is a tea.Msg implementation that modifies the border foreground
	// for a container component.
	ContainerBorderForegroundChangeMessage struct {
		// The ID of the container to modify.
		ID string
		// The desired border foreground.
		Foreground color.Color
	}
)

// ChangeContainerBorderForeground returns a tea.Cmd that generates a ContainerBorderForegroundChangeMessage which,
// when handled, modifies the foreground colour of a container's border.
func ChangeContainerBorderForeground(id string, foreground color.Color) tea.Cmd {
	return func() tea.Msg {
		return ContainerBorderForegroundChangeMessage{
			ID:         id,
			Foreground: foreground,
		}
	}
}
