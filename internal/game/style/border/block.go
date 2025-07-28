package border

import (
	"github.com/charmbracelet/lipgloss/v2"
)

// Block returns a lipgloss.Border using UTF8 block elements.
func Block() lipgloss.Border {
	return lipgloss.Border{
		Top:          "█",
		Bottom:       "█",
		Left:         "█",
		Right:        "█",
		TopLeft:      "█",
		TopRight:     "█",
		BottomLeft:   "█",
		BottomRight:  "█",
		MiddleLeft:   "█",
		MiddleRight:  "█",
		Middle:       "",
		MiddleTop:    "",
		MiddleBottom: "",
	}
}
