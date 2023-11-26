package app

import (
	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
)

// Component extends the tea.Model struct with help keys and SetSize
type Component interface {
	tea.Model
	help.KeyMap
	SetSize(width, height int)
}
