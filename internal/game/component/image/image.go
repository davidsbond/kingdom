package image

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss/v2"

	"github.com/davidsbond/kingdom/internal/game/asset"
)

type (
	image struct {
		content []string
		style   lipgloss.Style
	}
)

// Image returns a tea.Model implementation that is used to render an image. In the context of this game, an image is
// pretty much just a text file that we split into lines and place vertically.
func Image(name string) tea.Model {
	return &image{
		content: asset.Image(name),
		style:   lipgloss.NewStyle(),
	}
}

func (i *image) Init() tea.Cmd {
	return nil
}

func (i *image) Update(_ tea.Msg) (tea.Model, tea.Cmd) {
	return i, nil
}

func (i *image) View() string {
	return i.style.Render(
		lipgloss.JoinVertical(
			lipgloss.Center,
			i.content...,
		),
	)
}
