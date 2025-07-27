package layout_test

import (
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/x/exp/golden"

	"github.com/davidsbond/kingdom/internal/game/component/layout"
	"github.com/davidsbond/kingdom/internal/game/component/text"
	"github.com/davidsbond/kingdom/internal/game/window"
)

func TestCentered(t *testing.T) {
	t.Parallel()

	tt := []struct {
		Name     string
		Messages []tea.Msg
		Model    tea.Model
	}{
		{
			Name:  "text",
			Model: text.Text("test"),
			Messages: []tea.Msg{
				window.SizeMessage{
					Width:  104,
					Height: 25,
				},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.Name, func(t *testing.T) {
			centered := layout.Centered(tc.Model)
			centered.Init()

			for _, msg := range tc.Messages {
				centered.Update(msg)
			}

			golden.RequireEqual(t, []byte(centered.View()))
		})
	}
}
