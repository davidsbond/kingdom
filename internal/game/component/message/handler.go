package message

import (
	tea "github.com/charmbracelet/bubbletea"

	"github.com/davidsbond/kingdom/internal/game/component"
)

type (
	handler struct {
		component.NoInit
		component.NoView

		fn func(msg tea.Msg) tea.Cmd
	}
)

// Handler returns a tea.Model implementation that can handle arbitrary tea.Msg implementations.
func Handler(fn func(msg tea.Msg) tea.Cmd) tea.Model {
	return &handler{
		fn: fn,
	}
}

func (h *handler) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return h, h.fn(msg)
}
