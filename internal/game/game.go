package game

import (
	"sync"
)

type (
	// The State type is used to represent the current game state shared across all players.
	State struct {
		mux sync.Mutex

		playerOne *Player
		playerTwo *Player
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
			name: name,
		}

		return s.playerOne
	}

	if s.playerOne != nil && s.playerOne.name == name {
		return s.playerOne
	}

	if s.playerTwo == nil {
		s.playerTwo = &Player{
			name: name,
		}

		return s.playerTwo
	}

	if s.playerTwo != nil && s.playerTwo.name == name {
		return s.playerTwo
	}

	// TODO(davidsbond): we need to limit to two players connecting.
	panic("todo")
}

type (
	Player struct {
		name string
	}
)

func (p *Player) Name() string {
	return p.name
}
