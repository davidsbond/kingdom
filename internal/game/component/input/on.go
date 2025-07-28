package input

import (
	tea "github.com/charmbracelet/bubbletea"

	"github.com/davidsbond/kingdom/internal/game/component"
)

type (
	on struct {
		component.NoInit
		component.NoView

		key    tea.KeyType
		action func() tea.Cmd
	}
)

// On returns a tea.Model implementation that can handle a single key input and perform a function that optionally
// returns a tea.Cmd.
func On(key tea.KeyType, action func() tea.Cmd) tea.Model {
	return &on{
		key:    key,
		action: action,
	}
}

func (o *on) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if message, ok := msg.(tea.KeyMsg); ok && message.Type == o.key {
		return o, o.action()
	}

	return o, nil
}
