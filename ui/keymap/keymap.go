package keymap

import "github.com/charmbracelet/bubbles/key"

var (
	spacebar = " "
)

type KeyMap struct {
	Up   key.Binding
	Down key.Binding
	Quit key.Binding
	Help key.Binding

	// Player Commands
	TogglePlay    key.Binding
	NextTrack     key.Binding
	PreviousTrack key.Binding
	IncVolume     key.Binding
	DecVolume     key.Binding
	//Mute          key.Binding
	// volume
	// seek
	// repeat

	// PlayInfo Commands
	ToggleInfo key.Binding
}

func DefaultKeyMap() *KeyMap {
	return &KeyMap{
		Up: key.NewBinding(
			key.WithKeys("up", "k"),
			key.WithHelp("↑/k", "up"),
		),
		Down: key.NewBinding(
			key.WithKeys("down", "j"),
			key.WithHelp("↓/j", "down"),
		),
		Quit: key.NewBinding(
			key.WithKeys("q"),
			key.WithHelp("q", "quit"),
		),
		Help: key.NewBinding(
			key.WithKeys("?"),
			key.WithHelp("?", "toggle help"),
		),
		TogglePlay: key.NewBinding(
			key.WithKeys(spacebar),
			key.WithHelp("spacebar", "toggle play"),
		),
		NextTrack: key.NewBinding(
			key.WithKeys(">"),
			key.WithHelp(">", "next"),
		),
		PreviousTrack: key.NewBinding(
			key.WithKeys("<"),
			key.WithHelp("<", "previous"),
		),
		IncVolume: key.NewBinding(
			key.WithKeys("+"),
			key.WithHelp("+", "+volume"),
		),
		DecVolume: key.NewBinding(
			key.WithKeys("-"),
			key.WithHelp("-", "-volume"),
		),
		ToggleInfo: key.NewBinding(
			key.WithKeys("i"),
			key.WithHelp("i", "toggle info"),
		),
	}
}
