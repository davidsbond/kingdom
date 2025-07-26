package window

import (
	tea "github.com/charmbracelet/bubbletea"
)

type (
	// The Window type is used to manage the window of a single player.
	Window struct {
		width  int
		height int
	}
)

// New returns a Window instance that handles per-client window size tracking. When using wish, it seems
// that calling tea.WindowSize doesn't produce a tea.WindowSizeMsg. This means we need a model that lasts the lifetime
// of the client as when we switch from scene to scene any components that rely on the window size will no longer
// have it.
//
// This model handles the tea.WindowSizeMsg should any client change their window size, but also handles its own custom
// GetSizeMessage which will produce a Size message for any other models that want it.
func New(w, h int) *Window {
	return &Window{
		width:  w,
		height: h,
	}
}

func (w *Window) Init() tea.Cmd {
	return size(w.width, w.height)
}

func (w *Window) Update(msg tea.Msg) tea.Cmd {
	switch message := msg.(type) {
	case tea.WindowSizeMsg:
		w.width = message.Width
		w.height = message.Height

		return size(w.width, w.height)
	case GetSizeMessage:
		return size(w.width, w.height)
	}

	return nil
}
