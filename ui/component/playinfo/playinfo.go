package playinfo

import (
	"context"
	"fmt"
	"strings"

	"github.com/a-yee/spot/app"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

type ShuffleState bool

func (s ShuffleState) String() string {
	if s {
		return "on"
	}
	return "off"
}

type RepeatState int

const (
	RepeatOff RepeatState = iota
	RepeatOn
	RepeatTrack
)

func (r RepeatState) String() string {
	switch r {
	case RepeatOff:
		return "off"
	case RepeatOn:
		return "on"
	case RepeatTrack:
		return "track"
	default:
		return ""
	}
}

func NewRepeat(state string) RepeatState {
	switch state {
	case "off":
		return RepeatOff
	case "context":
		return RepeatOn
	case "track":
		return RepeatTrack
	default:
		return RepeatOff
	}
}

type PlayInfo struct {
	app     app.AppInfo
	Track   string
	Artists []string
	Shuffle ShuffleState
	Repeat  RepeatState
	Volume  int
}

func NewPlayInfo(ai app.AppInfo) *PlayInfo {
	return &PlayInfo{
		app:    ai,
		Repeat: RepeatOff,
	}
}

func (i *PlayInfo) SetSize(width, height int) {
	i.app.SetSize(width, height)
}

func (i *PlayInfo) Height() int {
	return lipgloss.Height(i.View())
}

func (i *PlayInfo) SetPlayInfo(track string) {
}

func (i *PlayInfo) Init() tea.Cmd {
	player, _ := i.app.API.PlayerState(context.Background())
	i.Track = player.Item.Name
	i.Shuffle = ShuffleState(player.ShuffleState)
	i.Repeat = NewRepeat(player.RepeatState)
	i.Volume = player.Device.Volume
	i.Artists = make([]string, len(player.Item.Artists))
	for j, artist := range player.Item.Artists {
		i.Artists[j] = artist.Name
	}
	return nil
}

func (i *PlayInfo) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return i, nil
}

func (i *PlayInfo) View() string {

	infoBar := table.New().Width(i.app.Width)

	w := lipgloss.Width
	s := i.app.Style
	section := s.PlayInfoSection

	track := s.PlayInfoTrack.Render("Playing: " + i.Track)
	sep := s.PlayInfoSep.Render("|")
	controls := lipgloss.JoinHorizontal(
		lipgloss.Top,
		section.Render("shuffle: "+i.Shuffle.String()),
		sep,
		section.Render("repeat: "+i.Repeat.String()),
		sep,
		section.Render("volume: "+fmt.Sprintf("%d", i.Volume)+"%"),
	)
	padding1 := s.PlayInfoPadding.Render(
		strings.Repeat(" ", i.app.Width-w(track)-w(controls)-8),
	)
	topLine := lipgloss.JoinHorizontal(
		lipgloss.Top,
		track,
		padding1,
		controls,
	)

	artists := s.PlayInfoTrack.Render("By: " + strings.Join(i.Artists, ", "))
	bottomLine := lipgloss.JoinHorizontal(
		lipgloss.Top,
		artists,
	)

	infoBar.Row(topLine)
	infoBar.Row(bottomLine)

	return infoBar.Render()
}

// TODO: might be able to remove later
//type TogglePlayInfoMsg struct{}
