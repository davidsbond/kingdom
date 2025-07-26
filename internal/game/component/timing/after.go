package timing

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type (
	after struct {
		duration time.Duration
		message  tea.Msg
	}
)

// After returns a tea.Model implementation that will produce the given command after a specified duration.
func After(duration time.Duration, cmd tea.Cmd) tea.Model {
	return &after{
		duration: duration,
		message:  cmd(),
	}
}

func (a *after) Init() tea.Cmd {
	return tea.Tick(a.duration, func(t time.Time) tea.Msg {
		return a.message
	})
}

func (a *after) Update(_ tea.Msg) (tea.Model, tea.Cmd) {
	return a, nil
}

func (a *after) View() string {
	return ""
}
