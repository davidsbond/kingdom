package component

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss/v2"
)

// Init performs initialisation of all the provided models, returning a batched command if any models returned a
// non-nil command.
func Init(models ...tea.Model) tea.Cmd {
	commands := make([]tea.Cmd, 0)
	for _, model := range models {
		if command := model.Init(); command != nil {
			commands = append(commands, command)
		}
	}

	if len(commands) > 0 {
		return tea.Batch(commands...)
	}

	return nil
}

// Update performs updates of all the provided models for the given message, returning a batched command if any
// models returned a non-nil command.
func Update(msg tea.Msg, models ...tea.Model) tea.Cmd {
	commands := make([]tea.Cmd, 0)
	for _, model := range models {
		if _, command := model.Update(msg); command != nil {
			commands = append(commands, command)
		}
	}

	if len(commands) > 0 {
		return tea.Batch(commands...)
	}

	return nil
}

// View returns a string representing the views of all the provided models joined vertically.
func View(models ...tea.Model) string {
	strings := make([]string, 0)
	for _, m := range models {
		if view := m.View(); view != "" {
			strings = append(strings, view)
		}
	}

	return lipgloss.JoinVertical(
		lipgloss.Top,
		strings...,
	)
}

type (
	// NoUpdate is to be embedded into components that lack the use of an Update method.
	NoUpdate struct{}

	// NoView is to be embedded into components that lack the use of a View method.
	NoView struct{}

	// NoInit is to be embedded into components that lack the use of an Init method.
	NoInit struct{}
)

func (nu *NoUpdate) Update(_ tea.Msg) (tea.Model, tea.Cmd) {
	return nil, nil
}

func (nu *NoView) View() string {
	return ""
}

func (nu *NoInit) Init() tea.Cmd {
	return nil
}
