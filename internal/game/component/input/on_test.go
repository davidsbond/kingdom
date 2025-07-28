package input_test

import (
	"testing"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/davidsbond/kingdom/internal/game/component/input"
)

func TestOn(t *testing.T) {
	t.Parallel()

	tt := []struct {
		Name     string
		Input    tea.KeyType
		Message  tea.Msg
		Expected tea.Msg
	}{
		{
			Name:  "input matches",
			Input: tea.KeyEnter,
			Message: tea.KeyMsg{
				Type: tea.KeyEnter,
			},
			Expected: tea.Quit(),
		},
		{
			Name:  "input does not match",
			Input: tea.KeyEnter,
			Message: tea.KeyMsg{
				Type: tea.KeyUp,
			},
			Expected: nil,
		},
	}

	for _, tc := range tt {
		t.Run(tc.Name, func(t *testing.T) {
			on := input.On(tc.Input, func() tea.Cmd {
				return func() tea.Msg {
					return tc.Expected
				}
			})

			_, command := on.Update(tc.Message)
			switch {
			case tc.Expected == nil && command != nil:
				t.Errorf("expected: nil, got: %v", command())
				return
			case tc.Expected != nil && command == nil:
				t.Errorf("expected: %v, got: nil", tc.Expected)
				return
			case tc.Expected == nil && command == nil:
				return
			}

			actual := command()
			if actual != tc.Expected {
				t.Errorf("expected: %v\nactual: %v", tc.Expected, actual)
			}
		})
	}
}
