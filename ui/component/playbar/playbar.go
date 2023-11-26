package playbar

import (
	"context"
	"strings"
	"time"

	"github.com/a-yee/spot/app"
	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	padding = 2
)

var helpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#626262")).Render

type Playbar struct {
	app           app.AppInfo
	progress      progress.Model
	trackProgress float64
	trackDuration float64
}

func NewPlaybar(ai app.AppInfo) Playbar {
	return Playbar{
		app: ai,
		progress: progress.New(
			progress.WithDefaultGradient(),
			//progress.WithoutPercentage(),
		),
	}
}

func (p Playbar) Init() tea.Cmd {
	return tickCmd()
}

func (p Playbar) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		p.progress.Width = msg.Width - padding*4 - 11
		return p, nil

	case tea.KeyMsg:
		return p, tea.Quit

	case tickMsg:
		cp, _ := p.app.API.PlayerCurrentlyPlaying(context.Background())

		p.trackProgress = float64(cp.Progress)
		p.trackDuration = float64(cp.Item.Duration)
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

func (p Playbar) View() string {
	pad := strings.Repeat(" ", padding)
	songProgress := (time.Duration(p.trackProgress) * time.Millisecond).
		Round(time.Second).
		String()
	songDuration := (time.Duration(p.trackDuration) * time.Millisecond).
		Round(time.Second).
		String()
	return "\n" +
		pad + p.progress.View() + pad +
		songProgress + " / " + songDuration + pad + "\n\n"
}

type tickMsg time.Time

func tickCmd() tea.Cmd {
	return tea.Tick(time.Millisecond*500, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}
