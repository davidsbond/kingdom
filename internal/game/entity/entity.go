package entity

import (
	tea "github.com/charmbracelet/bubbletea"

	"github.com/davidsbond/kingdom/internal/game/component"
)

type (
	entity struct {
		id string

		models []tea.Model
	}
)

func create(id string, models ...tea.Model) tea.Model {
	return &entity{
		id:     id,
		models: models,
	}
}

func (e *entity) Init() tea.Cmd {
	return component.Init(e.models...)
}

func (e *entity) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return e, component.Update(msg, e.models...)
}

func (e *entity) View() string {
	return component.View(e.models...)
}
