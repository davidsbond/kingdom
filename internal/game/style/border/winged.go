package border

import (
	"github.com/charmbracelet/lipgloss/v2"
)

// Wing returns a border with a wing on either its left or right side.
func Wing(p lipgloss.Position) lipgloss.Border {
	var border lipgloss.Border

	switch p {
	case lipgloss.Left:
		border.Left = "▜"
	case lipgloss.Right:
		border.Right = "▛"
	}

	return border
}
