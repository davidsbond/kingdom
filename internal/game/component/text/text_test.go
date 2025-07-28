package text_test

import (
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss/v2"
	"github.com/charmbracelet/x/exp/golden"

	"github.com/davidsbond/kingdom/internal/game/component/text"
)

func TestText(t *testing.T) {
	t.Parallel()

	tt := []struct {
		Name     string
		Content  string
		Options  []text.Option
		Messages []tea.Msg
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
		{
			Name:    "handles updates",
			Content: "test",
			Options: []text.Option{
				text.ID("test"),
			},
			Messages: []tea.Msg{
				text.ChangeMessage{
					ID:      "test",
					Content: "test2",
				},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.Name, func(t *testing.T) {
			txt := text.Text(tc.Content, tc.Options...)
			txt.Init()

			for _, msg := range tc.Messages {
				txt.Update(msg)
			}

			golden.RequireEqual(t, []byte(txt.View()))
		})
	}
}
