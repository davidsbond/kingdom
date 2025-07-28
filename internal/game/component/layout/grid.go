package layout

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss/v2"
)

type (
	grid struct {
		models [][]tea.Model
	}
)

func Grid(w, h int, models ...tea.Model) tea.Model {
	g := make([][]tea.Model, 0, h)

	for i := 0; i < h; i++ {
		start := i * w
		end := start + w

		if start >= len(models) {
			break
		}
		if end > len(models) {
			end = len(models)
		}

		row := models[start:end]
		g = append(g, row)
	}

	return &grid{models: g}
}

func (g *grid) Init() tea.Cmd {
	commands := make([]tea.Cmd, 0)
	for y := range g.models {
		for x := range g.models[y] {
			if command := g.models[y][x].Init(); command != nil {
				commands = append(commands, command)
			}
		}
	}

	return tea.Batch(commands...)
}

func (g *grid) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	commands := make([]tea.Cmd, 0)
	for y := range g.models {
		for x := range g.models[y] {
			if _, command := g.models[y][x].Update(msg); command != nil {
				commands = append(commands, command)
			}
		}
	}

	return g, tea.Batch(commands...)
}

func (g *grid) View() string {
	rows := make([]string, len(g.models))

	for y := range g.models {
		columns := make([]string, len(g.models[y]))
		for x := range g.models[y] {
			columns[x] = g.models[y][x].View()
		}

		rows[y] = lipgloss.JoinHorizontal(lipgloss.Center, columns...)
	}

	return lipgloss.JoinVertical(lipgloss.Top, rows...)
}
