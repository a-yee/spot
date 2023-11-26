package style

import "github.com/charmbracelet/lipgloss"

type Style struct {
	Footer lipgloss.Style
}

func DefaultStyles() *Style {
	s := new(Style)

	s.Footer = lipgloss.NewStyle().
		MarginTop(1).
		Padding(0, 1).
		Height(1)
	return s
}
