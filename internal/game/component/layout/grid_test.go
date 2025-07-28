package layout_test

import (
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/x/exp/golden"

	"github.com/davidsbond/kingdom/internal/game/component/layout"
	"github.com/davidsbond/kingdom/internal/game/component/text"
)

func TestGrid(t *testing.T) {
	t.Parallel()

	tt := []struct {
		Name   string
		Width  int
		Height int
		Models []tea.Model
	}{
		{
			Name:   "3x3",
			Width:  3,
			Height: 3,
			Models: []tea.Model{
				text.Text("1"),
				text.Text("2"),
				text.Text("3"),
				text.Text("4"),
				text.Text("5"),
				text.Text("6"),
				text.Text("7"),
				text.Text("8"),
				text.Text("9"),
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.Name, func(t *testing.T) {
			grid := layout.Grid(tc.Width, tc.Height, tc.Models...)
			grid.Init()

			golden.RequireEqual(t, []byte(grid.View()))
		})
	}
}
