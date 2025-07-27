package layout_test

import (
	"image/color"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss/v2"
	"github.com/charmbracelet/x/exp/golden"

	"github.com/davidsbond/kingdom/internal/game/component/layout"
	"github.com/davidsbond/kingdom/internal/game/component/text"
)

func TestContainer(t *testing.T) {
	t.Parallel()

	tt := []struct {
		Name    string
		Model   tea.Model
		Options []layout.ContainerOption
	}{
		{
			Name:  "default",
			Model: text.Text("test"),
		},
		{
			Name:  "all options",
			Model: text.Text("test"),
			Options: []layout.ContainerOption{
				layout.ContainerMargin(1),
				layout.ContainerMargin(1),
				layout.ContainerBorder(lipgloss.NormalBorder(), color.White),
				layout.ContainerAlign(lipgloss.Center),
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.Name, func(t *testing.T) {
			centered := layout.Container(tc.Model, tc.Options...)
			centered.Init()

			golden.RequireEqual(t, []byte(centered.View()))
		})
	}
}
