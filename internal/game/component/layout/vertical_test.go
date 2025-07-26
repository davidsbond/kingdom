package layout_test

import (
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/x/exp/golden"

	"github.com/davidsbond/kingdom/internal/game/component/layout"
	"github.com/davidsbond/kingdom/internal/game/component/text"
)

func TestVertical_View(t *testing.T) {
	t.Parallel()

	tt := []struct {
		Name   string
		Models []tea.Model
	}{
		{
			Name: "text",
			Models: []tea.Model{
				text.Text("a"),
				text.Text("b"),
				text.Text("c"),
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.Name, func(t *testing.T) {
			vertical := layout.Vertical(tc.Models...)
			vertical.Init()

			golden.RequireEqual(t, []byte(vertical.View()))
		})
	}
}
