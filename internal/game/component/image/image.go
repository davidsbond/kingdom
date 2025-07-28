package image

import (
	"image/color"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss/v2"
	"github.com/charmbracelet/log"

	"github.com/davidsbond/kingdom/internal/game/asset"
	"github.com/davidsbond/kingdom/internal/game/component"
)

type (
	image struct {
		component.NoUpdate
		component.NoInit

		content []string
		style   lipgloss.Style
	}

	// The Option type is a function used to modify an image.
	Option func(i *image)
)

// Image returns a tea.Model implementation that is used to render an image. In the context of this game, an image is
// pretty much just a text file that we split into lines and place vertically.
func Image(logger *log.Logger, name string, options ...Option) tea.Model {
	i := &image{
		content: asset.Image(logger, name),
		style:   lipgloss.NewStyle(),
	}

	for _, option := range options {
		option(i)
	}

	return i
}

func (i *image) View() string {
	return i.style.Render(
		lipgloss.JoinVertical(
			lipgloss.Top,
			i.content...,
		),
	)
}

// Foreground is an Option that modifies the foreground colour of the image.
func Foreground(c color.Color) Option {
	return func(i *image) {
		i.style = i.style.Foreground(c)
	}
}
