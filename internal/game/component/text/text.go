package text

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss/v2"
)

type (
	text struct {
		style   lipgloss.Style
		content string
	}
)

// Text returns a tea.Model implementation that displays some static text.
func Text(content string) tea.Model {
	return &text{
		content: content,
		style:   lipgloss.NewStyle(),
	}
}

func (t *text) Init() tea.Cmd {
	return nil
}

func (t *text) Update(_ tea.Msg) (tea.Model, tea.Cmd) {
	return t, nil
}

func (t *text) View() string {
	return t.style.Render(t.content)
}
