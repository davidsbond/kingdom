package text

import (
	tea "github.com/charmbracelet/bubbletea"
)

type (
	// The ChangeMessage type is a tea.Msg implementation that modifies the content of a dynamic text component.
	ChangeMessage struct {
		// The identifier of the component to modify.
		ID string
		// The new text content.
		Content string
	}
)

// Change returns a tea.Cmd implementation that produces a ChangeMessage. This is used to change the content of a
// dynamic text component.
func Change(id, content string) tea.Cmd {
	return func() tea.Msg {
		return ChangeMessage{
			ID:      id,
			Content: content,
		}
	}
}
