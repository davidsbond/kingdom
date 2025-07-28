package layout

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss/v2"

	"github.com/davidsbond/kingdom/internal/game/window"
)

type (
	centered struct {
		model  tea.Model
		width  int
		height int
	}
)

// Centered returns a tea.Model implementation that wraps another tea.Model implementation, ensuring it is centered
// relative to the window size when rendered.
func Centered(model tea.Model) tea.Model {
	return &centered{
		model: model,
	}
}

func (b *centered) Init() tea.Cmd {
	commands := []tea.Cmd{window.Size()}
	if command := b.model.Init(); command != nil {
		commands = append(commands, command)
	}

	return tea.Batch(commands...)
}

func (b *centered) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if message, ok := msg.(window.SizeMessage); ok {
		b.width = message.Width
		b.height = message.Height
	}

	return b.model.Update(msg)
}

func (b *centered) View() string {
	if b.width == 0 || b.height == 0 {
		return ""
	}
	
	return lipgloss.Place(
		b.width,
		b.height,
		lipgloss.Center,
		lipgloss.Center,
		b.model.View(),
	)
}
