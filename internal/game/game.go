package game

import (
	"strings"
	"sync"

	tea "github.com/charmbracelet/bubbletea"
)

type (
	// The State type is used to represent the current game state shared across all players.
	State struct {
		mux sync.Mutex

		playerOne *Player
		playerTwo *Player

		playerOneJoined sync.Once
		playerTwoJoined sync.Once
	}
)

// NewState returns a new State instance.
func NewState() *State {
	return &State{}
}

// Player returns the player with the matching name, creating player one or two on-demand.
func (s *State) Player(name string) *Player {
	s.mux.Lock()
	defer s.mux.Unlock()

	if s.playerOne == nil {
		s.playerOne = &Player{
			name:   name,
			number: 1,
		}

		return s.playerOne
	}

	if s.playerOne != nil && s.playerOne.name == name {
		return s.playerOne
	}

	if s.playerTwo == nil {
		s.playerTwo = &Player{
			name:   name,
			number: 2,
		}

		return s.playerTwo
	}

	if s.playerTwo != nil && s.playerTwo.name == name {
		return s.playerTwo
	}

	// TODO(davidsbond): we need to limit to two players connecting.
	panic("todo")
}

func (s *State) Update(msg tea.Msg) tea.Cmd {
	switch message := msg.(type) {
	case PlayerJoinedMessage:
		return s.handlePlayerJoined(message)
	case KingdomSelectionChangedMessage:
		return s.handleKingdomSelectionChanged(message)
	}

	return nil
}

func (s *State) sendToPlayer(number int, msg tea.Msg) {
	if number == 1 && s.playerOne != nil {
		s.playerOne.program.Send(msg)
	}

	if number == 2 && s.playerTwo != nil {
		s.playerTwo.program.Send(msg)
	}
}

// PlayerN returns the player associated with the given number.
func (s *State) PlayerN(n int) *Player {
	s.mux.Lock()
	defer s.mux.Unlock()

	if n == 1 {
		return s.playerOne
	}

	return s.playerTwo
}

func (s *State) handlePlayerJoined(message PlayerJoinedMessage) tea.Cmd {
	if message.Number == 2 {
		s.playerTwoJoined.Do(func() {
			s.sendToPlayer(1, message)
		})
	}

	if message.Number == 1 {
		s.playerOneJoined.Do(func() {
			s.sendToPlayer(2, message)
		})
	}

	return nil
}

func (s *State) handleKingdomSelectionChanged(message KingdomSelectionChangedMessage) tea.Cmd {
	if message.handled {
		return nil
	}

	// We set this boolean to true so that when a player receives it, it is not then repropagated by their program
	// loop.
	message.handled = true

	if message.Player == 1 {
		s.sendToPlayer(2, message)
	}

	if message.Player == 2 {
		s.sendToPlayer(1, message)
	}

	return nil
}

type (
	// The Player type represents a single player.
	Player struct {
		name    string
		number  int
		program *tea.Program
	}
)

// Name returns the upper-cased version of the player's name. This is usually their SSH username.
func (p *Player) Name() string {
	return strings.ToUpper(p.name)
}

// Number returns the player number (1 or 2).
func (p *Player) Number() int {
	return p.number
}

// SetProgram sets the given tea.Program as belonging to the player and will be used for message propagation.
func (p *Player) SetProgram(program *tea.Program) {
	p.program = program
}
