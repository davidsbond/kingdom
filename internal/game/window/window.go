package window

import (
	tea "github.com/charmbracelet/bubbletea"
)

type (
	window struct {
		width  int
		height int
	}
)

// Window returns a tea.Model implementation that handles per-client window size tracking. When using wish, it seems
// that calling tea.WindowSize doesn't produce a tea.WindowSizeMsg. This means we need a model that lasts the lifetime
// of the client as when we switch from scene to scene any components that rely on the window size will no longer
// have it.
//
// This model handles the tea.WindowSizeMsg should any client change their window size, but also handles its own custom
// GetSizeMessage which will produce a Size message for any other models that want it.
func Window(w, h int) tea.Model {
	return &window{
		width:  w,
		height: h,
	}
}

func (w *window) Init() tea.Cmd {
	return size(w.width, w.height)
}

func (w *window) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch message := msg.(type) {
	case tea.WindowSizeMsg:
		w.width = message.Width
		w.height = message.Height

		return w, size(w.width, w.height)
	case GetSizeMessage:
		return w, size(w.width, w.height)
	}

	return w, nil
}

func (w *window) View() string {
	return ""
}
