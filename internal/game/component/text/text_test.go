package text_test

import (
	"testing"

	"github.com/charmbracelet/lipgloss/v2"
	"github.com/charmbracelet/x/exp/golden"

	"github.com/davidsbond/kingdom/internal/game/component/text"
)

func TestText_View(t *testing.T) {
	t.Parallel()

	tt := []struct {
		Name    string
		Content string
		Options []text.Option
	}{
		{
			Name:    "no options",
			Content: "hello world",
		},
		{
			Name:    "width",
			Content: "hello world",
			Options: []text.Option{
				text.Width(35),
			},
		},
		{
			Name:    "width and align",
			Content: "hello world",
			Options: []text.Option{
				text.Width(35),
				text.Align(lipgloss.Right),
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.Name, func(t *testing.T) {
			txt := text.Text(tc.Content, tc.Options...)
			txt.Init()

			golden.RequireEqual(t, []byte(txt.View()))
		})
	}
}
