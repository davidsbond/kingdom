package text_test

import (
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/x/exp/golden"

	"github.com/davidsbond/kingdom/internal/game/component/text"
)

func TestDynamic(t *testing.T) {
	t.Parallel()

	tt := []struct {
		Name     string
		ID       string
		Options  []text.Option
		Messages []tea.Msg
	}{
		{
			Name: "changes with message",
			ID:   "test",
			Options: []text.Option{
				text.Content("first"),
			},
			Messages: []tea.Msg{
				text.ChangeMessage{
					ID:      "test",
					Content: "second",
				},
			},
		},
		{
			Name: "ignores other messages",
			ID:   "test",
			Options: []text.Option{
				text.Content("first"),
			},
			Messages: []tea.Msg{
				text.ChangeMessage{
					ID:      "test2",
					Content: "second",
				},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.Name, func(t *testing.T) {
			txt := text.Dynamic(tc.ID, tc.Options...)
			txt.Init()

			for _, msg := range tc.Messages {
				txt.Update(msg)
			}

			golden.RequireEqual(t, []byte(txt.View()))
		})
	}
}
