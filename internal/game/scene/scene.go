package scene

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss/v2"
)

type (
	scene struct {
		window tea.Model
		models []tea.Model
	}
)

func create(window tea.Model, models ...tea.Model) tea.Model {
	return &scene{window: window, models: models}
}

func (s *scene) Init() tea.Cmd {
	commands := []tea.Cmd{s.window.Init()}
	for _, model := range s.models {
		if cmd := model.Init(); cmd != nil {
			commands = append(commands, cmd)
		}
	}

	return tea.Batch(commands...)
}

func (s *scene) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if message, ok := msg.(ChangeMessage); ok {
		c := message.To(s.window)
		return c, c.Init()
	}

	commands := make([]tea.Cmd, 0)
	for _, model := range s.models {
		if _, cmd := model.Update(msg); cmd != nil {
			commands = append(commands, cmd)
		}
	}

	if _, command := s.window.Update(msg); command != nil {
		commands = append(commands, command)
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
