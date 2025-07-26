package text_test

import (
	"testing"

	"github.com/charmbracelet/x/exp/golden"

	"github.com/davidsbond/kingdom/internal/game/component/text"
)

func TestText_View(t *testing.T) {
	t.Parallel()

	tt := []struct {
		Name    string
		Content string
	}{
		{
			Name:    "text",
			Content: "hello world",
		},
	}

	for _, tc := range tt {
		t.Run(tc.Name, func(t *testing.T) {
			txt := text.Text(tc.Content)
			txt.Init()

			golden.RequireEqual(t, []byte(txt.View()))
		})
	}
}
