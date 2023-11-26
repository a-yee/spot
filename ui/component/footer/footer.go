package footer

import (
	"github.com/a-yee/spot/app"
	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
)

type Footer struct {
	app    app.AppInfo
	help   help.Model
	keymap help.KeyMap
}

func New(ai app.AppInfo, keymap help.KeyMap) *Footer {
	h := help.New()
	f := &Footer{
		app:    ai,
		help:   h,
		keymap: keymap,
	}
	f.SetSize(ai.Width, ai.Height)
	return f
}

func (f *Footer) SetSize(width, height int) {
	f.app.SetSize(width, height)
	f.help.Width = width -
		f.app.Style.Footer.GetHorizontalFrameSize()
}

func (f *Footer) Init() tea.Cmd {
	return nil
}

func (f *Footer) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case tea.KeyMsg:
		return f, tea.Quit
	}
	return f, nil
}

func (f *Footer) View() string {
	if f.keymap == nil {
		return ""
	}
	s := f.app.Style.Footer.Copy().Width(f.app.Width)
	helpView := f.help.View(f.keymap)
	return s.Render(helpView)
}
