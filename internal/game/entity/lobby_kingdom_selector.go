package entity

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss/v2"
	"github.com/charmbracelet/log"

	"github.com/davidsbond/kingdom/internal/game"
	"github.com/davidsbond/kingdom/internal/game/component/image"
	"github.com/davidsbond/kingdom/internal/game/component/input"
	"github.com/davidsbond/kingdom/internal/game/component/layout"
	"github.com/davidsbond/kingdom/internal/game/component/message"
	"github.com/davidsbond/kingdom/internal/game/style/border"
)

func LobbyKingdomSelector(logger *log.Logger, active bool, player int) tea.Model {
	id := fmt.Sprintf("lobby-kindom-selector-%d", player)
	logger = logger.With("entity", id)

	const (
		gridWidth  = 3
		gridHeight = 3
	)

	img := image.Image(logger, "filler.txt")

	cells := make([]tea.Model, 9)
	cellIDs := make([]string, 9)
	selected := 0

	colour := lipgloss.Red
	if player == 2 {
		colour = lipgloss.Blue
	}

	brd := border.Block()

	for i := range len(cells) {
		cellID := fmt.Sprintf("%s-cell-%d", id, i)

		cellIDs[i] = cellID
		cells[i] = layout.Container(img,
			layout.ContainerID(cellID),
			layout.ContainerBorder(brd, lipgloss.White),
			layout.ContainerMargin(0, 1, 1, 1),
			layout.ContainerLogger(logger),
		)

	}

	logger.Debug("initialising entity")
	return create(id,
		message.Init(layout.ChangeContainerBorderForeground(cellIDs[0], colour)),

		message.Handler(func(msg tea.Msg) tea.Cmd {
			if m, ok := msg.(game.KingdomSelectionChangedMessage); ok && m.Player == player && !active {
				previous := selected
				selected = m.Kingdom
				return tea.Batch(
					layout.ChangeContainerBorderForeground(cellIDs[previous], lipgloss.White),
					layout.ChangeContainerBorderForeground(cellIDs[selected], colour),
				)
			}

			return nil
		}),

		input.On(tea.KeyUp, func() tea.Cmd {
			if !active {
				return nil
			}

			previous := selected
			selected -= gridHeight
			if selected < 0 {
				selected += gridHeight
				return nil
			}

			return tea.Batch(
				layout.ChangeContainerBorderForeground(cellIDs[previous], lipgloss.White),
				layout.ChangeContainerBorderForeground(cellIDs[selected], colour),
				game.KingdomSelectionChanged(player, selected),
			)
		}),

		input.On(tea.KeyDown, func() tea.Cmd {
			if !active {
				return nil
			}

			previous := selected
			selected += gridHeight
			if selected > len(cells)-1 {
				selected -= gridHeight
				return nil
			}

			return tea.Batch(
				layout.ChangeContainerBorderForeground(cellIDs[previous], lipgloss.White),
				layout.ChangeContainerBorderForeground(cellIDs[selected], colour),
				game.KingdomSelectionChanged(player, selected),
			)
		}),

		input.On(tea.KeyRight, func() tea.Cmd {
			if !active {
				return nil
			}

			previous := selected
			selected++
			if selected > len(cells)-1 {
				selected = 0
			}

			return tea.Batch(
				layout.ChangeContainerBorderForeground(cellIDs[previous], lipgloss.White),
				layout.ChangeContainerBorderForeground(cellIDs[selected], colour),
				game.KingdomSelectionChanged(player, selected),
			)
		}),

		input.On(tea.KeyLeft, func() tea.Cmd {
			if !active {
				return nil
			}

			previous := selected
			selected--
			if selected < 0 {
				selected = len(cells) - 1
			}

			return tea.Batch(
				layout.ChangeContainerBorderForeground(cellIDs[previous], lipgloss.White),
				layout.ChangeContainerBorderForeground(cellIDs[selected], colour),
				game.KingdomSelectionChanged(player, selected),
			)
		}),

		layout.Container(
			layout.Grid(gridWidth, gridHeight, cells...),
			layout.ContainerMargin(0, 2, 0),
		),
	)
}
