package playbar

import (
	"context"
	"strings"
	"time"

	"github.com/a-yee/spot/app"
	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/reflow/truncate"
)

var (
	padding       = 8
	maxTrackWidth = 20
)

var helpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#626262")).Render

type Playbar struct {
	app           app.AppInfo
	progress      progress.Model
	trackProgress float64
	trackDuration float64
	track         string
}

func NewPlaybar(ai app.AppInfo) *Playbar {
	return &Playbar{
		app: ai,
		progress: progress.New(
			progress.WithDefaultGradient(),
			//progress.WithoutPercentage(),
		),
	}
}

func (p *Playbar) SetSize(width, height int) {
	p.app.SetSize(width, height)
}

func (p *Playbar) Init() tea.Cmd {
	return tickCmd()
}

func (p *Playbar) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		p.SetSize(msg.Width, 0)
		return p, nil

	case tea.KeyMsg:
		return p, tea.Quit

	case tickMsg:
		cp, _ := p.app.API.PlayerCurrentlyPlaying(context.Background())

		p.trackProgress = 0.0
		p.trackDuration = 0.0
		if cp.Item != nil {
			p.track = cp.Item.Name
			p.trackProgress = float64(cp.Progress)
			p.trackDuration = float64(cp.Item.Duration)
		}
		cmd := p.progress.SetPercent(p.trackProgress / p.trackDuration)
		return p, tea.Batch(tickCmd(), cmd)

	case progress.FrameMsg:
		progressModel, cmd := p.progress.Update(msg)
		p.progress = progressModel.(progress.Model)
		return p, cmd

	default:
		return p, nil
	}
}

func (p *Playbar) View() string {
	s := p.app.Style
	w := lipgloss.Width

	//pad := s.PlaybarPadding.Render(strings.Repeat(" ", padding))
	pad := strings.Repeat("-", padding)
	// TODO: fix the padding around short track names
	t := truncate.StringWithTail(
		p.track,
		uint(maxTrackWidth-s.PlaybarTrack.GetHorizontalFrameSize()),
		"…",
	)
	track := s.PlaybarTrack.
		Width(maxTrackWidth).
		Render(t)
	trackProgress := (time.Duration(p.trackProgress) * time.Millisecond).
		Round(time.Second).
		String()
	trackDuration := (time.Duration(p.trackDuration) * time.Millisecond).
		Round(time.Second).
		String()
	playProgress := s.PlaybarProgress.Render(
		trackProgress + " / " + trackDuration)

	help := s.PlaybarHelp.Render("? help")
	// TODO: fix the values to prevent bar from yoyo-ing in size
	maxWidth := p.app.Width - w(track) - w(playProgress) - w(help) - 4*padding
	p.progress.Width = maxWidth

	return lipgloss.NewStyle().MaxWidth(p.app.Width).
		Render(
			lipgloss.JoinHorizontal(lipgloss.Top,
				track,
				pad,
				pad,
				p.progress.View(),
				pad,
				pad,
				playProgress,
				help,
			),
		)
}

type tickMsg time.Time

func tickCmd() tea.Cmd {
	return tea.Tick(time.Millisecond*500, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}
