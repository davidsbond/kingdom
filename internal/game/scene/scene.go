package scene

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss/v2"
)

type (
	scene struct {
		models []tea.Model
	}
)

func create(models ...tea.Model) tea.Model {
	return &scene{models: models}
}

func (s *scene) Init() tea.Cmd {
	commands := make([]tea.Cmd, 0)
	for _, model := range s.models {
		if cmd := model.Init(); cmd != nil {
			commands = append(commands, cmd)
		}
	}

	return tea.Batch(commands...)
}

func (s *scene) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	commands := make([]tea.Cmd, 0)
	for _, model := range s.models {
		if _, cmd := model.Update(msg); cmd != nil {
			commands = append(commands, cmd)
		}
	}

	return s, tea.Batch(commands...)
}

func (s *scene) View() string {
	strings := make([]string, 0)
	for _, e := range s.models {
		if view := e.View(); view != "" {
			strings = append(strings, view)
		}
	}

	return lipgloss.JoinVertical(
		lipgloss.Top,
		strings...,
	)
}
