package layout

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss/v2"

	"github.com/davidsbond/kingdom/internal/game/component"
)

type (
	vertical struct {
		models []tea.Model
	}
)

// Vertical returns a tea.Model implementation that renders all provided models vertically relative to one another.
func Vertical(models ...tea.Model) tea.Model {
	return &vertical{models: models}
}

func (v *vertical) Init() tea.Cmd {
	return component.Init(v.models...)
}

func (v *vertical) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return v, component.Update(msg, v.models...)
}

func (v *vertical) View() string {
	strings := make([]string, 0)
	for _, e := range v.models {
		if view := e.View(); view != "" {
			strings = append(strings, view)
		}
	}

	return lipgloss.JoinVertical(
		lipgloss.Top,
		strings...,
	)
}
