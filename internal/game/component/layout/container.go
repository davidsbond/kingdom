package layout

import (
	"image/color"
	"io"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss/v2"
	"github.com/charmbracelet/log"

	"github.com/davidsbond/kingdom/internal/game/component"
)

type (
	container struct {
		id     string
		model  tea.Model
		style  lipgloss.Style
		logger *log.Logger
	}

	// The ContainerOption type is a function that modifies a container component.
	ContainerOption func(c *container)
)

// Container returns a tea.Model implementation that wraps another tea.Model implementation with the ability to apply
// styles to it as a whole.
func Container(model tea.Model, options ...ContainerOption) tea.Model {
	c := &container{
		model:  model,
		style:  lipgloss.NewStyle(),
		logger: log.New(io.Discard),
	}

	for _, option := range options {
		option(c)
	}

	return c
}

func (c *container) Init() tea.Cmd {
	return component.Init(c.model)
}

func (c *container) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	logger := c.logger.With("component", c.id)

	switch message := msg.(type) {
	case ContainerBorderForegroundChangeMessage:
		if message.ID != c.id {
			break
		}

		logger.Debug("changing container border foreground")
		c.style = c.style.BorderForeground(message.Foreground)
	}

	return c, component.Update(msg, c.model)
}

func (c *container) View() string {
	return c.style.Render(c.model.View())
}

// ContainerAlign is a ContainerOption that sets the alignment of the container.
func ContainerAlign(p lipgloss.Position) ContainerOption {
	return func(c *container) {
		c.style = c.style.Align(p)
	}
}

// ContainerMargin is a ContainerOption that sets the margin of the container.
func ContainerMargin(m ...int) ContainerOption {
	return func(c *container) {
		c.style = c.style.Margin(m...)
	}
}

// ContainerBorder is a ContainerOption that sets the border of the container.
func ContainerBorder(b lipgloss.Border, colour color.Color) ContainerOption {
	return func(c *container) {
		c.style = c.style.Border(b).BorderForeground(colour)
	}
}

// ContainerID is a ContainerOption that sets the id of the container. This is used to enable dynamic behaviour
// on the container.
func ContainerID(id string) ContainerOption {
	return func(c *container) {
		c.id = id
	}
}

func ContainerLogger(logger *log.Logger) ContainerOption {
	return func(c *container) {
		c.logger = logger
	}
}
