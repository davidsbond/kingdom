package entity_test

import (
	"io"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
	"github.com/charmbracelet/x/exp/golden"

	"github.com/davidsbond/kingdom/internal/game"
	"github.com/davidsbond/kingdom/internal/game/entity"
)

func TestLobbyPlayerTitle(t *testing.T) {
	t.Parallel()

	tt := []struct {
		Name     string
		Messages []tea.Msg
	}{
		{
			Name: "handles player join",
			Messages: []tea.Msg{
				game.PlayerJoinedMessage{
					Number: 1,
					Name:   "Test",
				},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.Name, func(t *testing.T) {
			logger := log.New(io.Discard)
			state := game.NewState()

			ent := entity.LobbyPlayerTitle(logger, state, 1)
			pump(t, ent, tc.Messages...)

			golden.RequireEqual(t, []byte(ent.View()))
		})
	}
}
