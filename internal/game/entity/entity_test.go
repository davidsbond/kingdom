package entity_test

import (
	"testing"

	tea "github.com/charmbracelet/bubbletea"
)

func pump(t *testing.T, model tea.Model, messages ...tea.Msg) {
	t.Helper()

	for _, message := range messages {
		switch m := message.(type) {
		case tea.BatchMsg:
			for _, msg := range m {
				pump(t, model, msg)
			}
		}

		if _, cmd := model.Update(message); cmd != nil {
			pump(t, model, cmd())
		}
	}
}
