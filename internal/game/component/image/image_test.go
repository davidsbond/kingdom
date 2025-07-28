package image_test

import (
	"io"
	"testing"

	"github.com/charmbracelet/log"
	"github.com/charmbracelet/x/exp/golden"

	"github.com/davidsbond/kingdom/internal/game/component/image"
)

func TestImage(t *testing.T) {
	t.Parallel()

	tt := []struct {
		Name string
	}{
		{
			Name: "logo.txt",
		},
	}

	for _, tc := range tt {
		t.Run(tc.Name, func(t *testing.T) {
			logger := log.New(io.Discard)
			txt := image.Image(logger, tc.Name)
			txt.Init()

			golden.RequireEqual(t, []byte(txt.View()))
		})
	}
}
