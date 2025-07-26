package layout

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type (
	centered struct {
		model  tea.Model
		width  int
		height int
	}
)

func Centered(model tea.Model) tea.Model {
	return &centered{
		model: model,
	}
}

func (b *centered) Init() tea.Cmd {
	return tea.Batch(tea.WindowSize(), b.model.Init())
}

func (b *centered) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if message, ok := msg.(tea.WindowSizeMsg); ok {
		b.width = message.Width
		b.height = message.Height
	}

	return b.model.Update(msg)
}

func (b *centered) View() string {
	return lipgloss.Place(
		b.width,
		b.height/2,
		lipgloss.Center,
		lipgloss.Top,
		b.model.View(),
	)
}
