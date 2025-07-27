package text

import (
	"image/color"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss/v2"

	"github.com/davidsbond/kingdom/internal/game/component"
)

type (
	text struct {
		component.NoUpdate
		component.NoInit

		style   lipgloss.Style
		content string
	}

	// The Option type is a function used to modify the text instance.
	Option func(txt *text)
)

// Text returns a tea.Model implementation that displays some static text.
func Text(content string, options ...Option) tea.Model {
	txt := &text{
		content: content,
		style:   lipgloss.NewStyle(),
	}

	for _, option := range options {
		option(txt)
	}

	return txt
}

func (t *text) View() string {
	return t.style.Render(t.content)
}

// Width is an Option that modifies the width of a piece of text.
func Width(w int) Option {
	return func(txt *text) {
		txt.style = txt.style.Width(w)
	}
}

// Align is an Option that modifies the alignment of a piece of text.
func Align(p lipgloss.Position) Option {
	return func(txt *text) {
		txt.style = txt.style.Align(p)
	}
}

// Foreground is an Option that modifies the foreground colour of a piece of text.
func Foreground(c color.Color) Option {
	return func(txt *text) {
		txt.style = txt.style.Foreground(c)
	}
}

// Background is an Option that modifies the background colour of a piece of text.
func Background(c color.Color) Option {
	return func(txt *text) {
		txt.style = txt.style.Background(c)
	}
}

// Padding is an Option that modifies the padding of a piece of text.
func Padding(p ...int) Option {
	return func(txt *text) {
		txt.style = txt.style.Padding(p...)
	}
}
