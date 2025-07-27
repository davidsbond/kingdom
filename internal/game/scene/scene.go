package scene

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"

	"github.com/davidsbond/kingdom/internal/game"
	"github.com/davidsbond/kingdom/internal/game/component"
	"github.com/davidsbond/kingdom/internal/game/window"
)

type (
	scene struct {
		models []tea.Model
		ctx    Context
	}

	// The Context type is a bag of things that all scenes depend on that act as global state.
	Context struct {
		// Information on the current window.
		Window *window.Window
		// Information on the current player.
		Player *game.Player
		// The game state shared across all players.
		State *game.State
		// The logger for the current player.
		Logger *log.Logger
	}
)

func create(ctx Context, models ...tea.Model) tea.Model {
	return &scene{
		models: models,
		ctx:    ctx,
	}
}

func (s *scene) Init() tea.Cmd {
	commands := []tea.Cmd{s.ctx.Window.Init()}
	for _, model := range s.models {
		if cmd := model.Init(); cmd != nil {
			commands = append(commands, cmd)
		}
	}

	return tea.Batch(commands...)
}

func (s *scene) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if message, ok := msg.(ChangeMessage); ok {
		c := message.To(s.ctx)

		return c, c.Init()
	}

	commands := make([]tea.Cmd, 0)
	for _, model := range s.models {
		if _, cmd := model.Update(msg); cmd != nil {
			commands = append(commands, cmd)
		}
	}

	if command := s.ctx.State.Update(msg); command != nil {
		commands = append(commands, command)
	}

	if command := s.ctx.Window.Update(msg); command != nil {
		commands = append(commands, command)
	}

	return s, tea.Batch(commands...)
}

func (s *scene) View() string {
	return component.View(s.models...)
}
