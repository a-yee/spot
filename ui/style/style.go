package style

import "github.com/charmbracelet/lipgloss"

type Style struct {
	Footer lipgloss.Style

	PlaybarTrack    lipgloss.Style
	PlaybarProgress lipgloss.Style
	PlaybarHelp     lipgloss.Style
	PlaybarPadding  lipgloss.Style
}

func DefaultStyles() *Style {
	s := new(Style)

	s.Footer = lipgloss.NewStyle().
		MarginTop(1).
		Padding(0, 1).
		Height(1)

	s.PlaybarTrack = lipgloss.NewStyle().
		Padding(0, 1).
		Background(lipgloss.Color("#0070C0")).
		Foreground(lipgloss.Color("#FFFFFF"))

	s.PlaybarHelp = lipgloss.NewStyle().
		Padding(0, 1).
		Background(lipgloss.Color("237")).
		Foreground(lipgloss.Color("243"))

	s.PlaybarProgress = lipgloss.NewStyle().
		Padding(0, 1).
		Background(lipgloss.Color("#0070C0")).
		Foreground(lipgloss.Color("#FFFFFF"))

	s.PlaybarPadding = lipgloss.NewStyle().
		Background(lipgloss.Color("235")).
		Foreground(lipgloss.Color("243"))
	return s
}
