package ui

import (
	"context"
	"strings"

	"github.com/a-yee/spot/app"
	"github.com/a-yee/spot/ui/component/footer"
	"github.com/a-yee/spot/ui/component/playbar"
	"github.com/a-yee/spot/ui/component/playinfo"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	debug           string
	volumeIncrement = 5
)

type AppModel struct {
	//quitting   bool
	//error      error
	app          app.AppInfo
	footer       *footer.Footer
	showFooter   bool
	playbar      *playbar.Playbar
	playinfo     *playinfo.PlayInfo
	showPlayInfo bool
	debug        string
}

func NewAppModel(ai app.AppInfo) *AppModel {
	return &AppModel{
		app:      ai,
		playbar:  playbar.NewPlaybar(ai),
		playinfo: playinfo.NewPlayInfo(ai),
	}
}

func (m *AppModel) SetFooter(f *footer.Footer) {
	m.footer = f
}

// Implement the help.KeyMap interface
func (m AppModel) ShortHelp() []key.Binding {
	return []key.Binding{
		m.app.KeyMap.Help,
		m.app.KeyMap.Quit,
	}
}

// Implement the help.KeyMap interface
func (m AppModel) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{
			m.app.KeyMap.Help,
			m.app.KeyMap.Quit,
			m.app.KeyMap.ToggleInfo,
		},
		{
			m.app.KeyMap.TogglePlay,
			m.app.KeyMap.NextTrack,
			m.app.KeyMap.PreviousTrack,
		},
		{
			m.app.KeyMap.IncVolume,
			m.app.KeyMap.DecVolume,
		},
	}
}

func (m *AppModel) SetSize(width, height int) {
	m.app.SetSize(width, height)
	s := m.app.Style.App.Copy()
	appBorderWidth := s.GetHorizontalFrameSize()
	appBorderHeight := s.GetVerticalFrameSize()

	if m.showPlayInfo {
		appBorderHeight += m.playinfo.Height()
	}

	if m.showFooter {
		appBorderHeight += m.footer.Height()
	}

	m.playinfo.SetSize(width-appBorderWidth, height-appBorderHeight)
	m.playbar.SetSize(width-appBorderWidth, height-appBorderHeight)
	m.footer.SetSize(width-appBorderWidth, height-appBorderHeight)
}

func (m *AppModel) Init() tea.Cmd {
	return tea.Batch(
		m.playinfo.Init(),
		m.playbar.Init(),
		m.footer.Init(),
	)
}

func (m *AppModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	cmds := make([]tea.Cmd, 0)

	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.SetSize(msg.Width, msg.Height)

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.app.KeyMap.Help):
			cmds = append(cmds, footer.ToggleFooterCmd)
		case key.Matches(msg, m.app.KeyMap.Quit):
			m.app.API.Pause(context.Background())
			cmds = append(cmds, tea.Quit)
		case key.Matches(msg, m.app.KeyMap.TogglePlay):
			if m.playbar.IsPlaying() {
				m.app.API.Pause(context.Background())
				m.playbar.SetPlaying(false)
			} else {
				m.app.API.Play(context.Background())
				m.playbar.SetPlaying(true)
			}
		case key.Matches(msg, m.app.KeyMap.NextTrack):
			m.app.API.Next(context.Background())
		case key.Matches(msg, m.app.KeyMap.PreviousTrack):
			m.app.API.Previous(context.Background())

		// TODO: parametrize volume increments + set sensible default
		case key.Matches(msg, m.app.KeyMap.IncVolume):
			playerState, _ := m.app.API.PlayerState(context.Background())
			currentVolume := playerState.Device.Volume
			volumeSet := 100
			if currentVolume+volumeIncrement < 100 {
				volumeSet = currentVolume + volumeIncrement
			}
			m.app.API.Volume(context.Background(), volumeSet)
		case key.Matches(msg, m.app.KeyMap.DecVolume):
			playerState, _ := m.app.API.PlayerState(context.Background())
			currentVolume := playerState.Device.Volume
			volumeSet := 0
			if currentVolume-volumeIncrement > 0 {
				volumeSet = currentVolume - volumeIncrement
			}
			m.app.API.Volume(context.Background(), volumeSet)

		case key.Matches(msg, m.app.KeyMap.ToggleInfo):
			m.showPlayInfo = !m.showPlayInfo
		}

		//case playinfo.TogglePlayInfoMsg:
		//m.showPlayInfo = !m.showPlayInfo

	case footer.ToggleFooterMsg:
		m.showFooter = !m.showFooter
		m.footer.SetShowAll(m.showFooter)
	}

	i, cmd := m.playinfo.Update(msg)
	m.playinfo = i.(*playinfo.PlayInfo)
	if cmd != nil {
		cmds = append(cmds, cmd)
	}

	f, cmd := m.footer.Update(msg)
	m.footer = f.(*footer.Footer)
	if cmd != nil {
		cmds = append(cmds, cmd)
	}

	p, cmd := m.playbar.Update(msg)
	m.playbar = p.(*playbar.Playbar)
	if cmd != nil {
		cmds = append(cmds, cmd)
	}

	m.SetSize(m.app.Width, m.app.Height)

	return m, tea.Batch(cmds...)
}

func (m *AppModel) View() string {
	var view string

	placeholderLines := lipgloss.Height(m.playbar.View()) + m.app.Style.App.GetVerticalFrameSize()

	if m.showPlayInfo {
		placeholderLines += m.playinfo.Height()
		view = lipgloss.JoinVertical(lipgloss.Top, view, m.playinfo.View())
	}

	view = lipgloss.JoinVertical(lipgloss.Top, view, m.playbar.View())

	if m.showFooter {
		placeholderLines += m.footer.Height()
		view = lipgloss.JoinVertical(lipgloss.Top, view, m.footer.View())
	}
	if m.app.Height > 0 {
		return strings.Repeat("\n", m.app.Height-placeholderLines) + view
	}
	return view
}
