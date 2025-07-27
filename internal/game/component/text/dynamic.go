package text

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss/v2"
)

type (
	dynamic struct {
		id string

		text
	}
)

// Dynamic returns a tea.Model implementation that displays a piece of text. This text can be modified via messaging
// using the Change function.
func Dynamic(id string, options ...Option) tea.Model {
	txt := &dynamic{
		id: id,
		text: text{
			style: lipgloss.NewStyle(),
		},
	}

	for _, option := range options {
		option(&txt.text)
	}

	return txt
}

func (d *dynamic) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if message, ok := msg.(ChangeMessage); ok && message.ID == d.id {
		d.content = message.Content
	}

	return d, nil
}

// Content is an Option that modifies the content of a text component.
func Content(content string) Option {
	return func(txt *text) {
		txt.content = content
	}
}
