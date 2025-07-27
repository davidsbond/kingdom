package layout

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss/v2"

	"github.com/davidsbond/kingdom/internal/game/component"
)

type (
	horizontal struct {
		models []tea.Model
	}
)

// Horizontal returns a tea.Model implementation that renders all provided models horizontally relative to one another.
func Horizontal(models ...tea.Model) tea.Model {
	return &horizontal{models: models}
}

func (v *horizontal) Init() tea.Cmd {
	return component.Init(v.models...)
}

func (v *horizontal) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return v, component.Update(msg, v.models...)
}

func (v *horizontal) View() string {
	strings := make([]string, 0)
	for _, e := range v.models {
		if view := e.View(); view != "" {
			strings = append(strings, view)
		}
	}

	return lipgloss.JoinHorizontal(
		lipgloss.Center,
		strings...,
	)
}
