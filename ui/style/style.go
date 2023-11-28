package style

import "github.com/charmbracelet/lipgloss"

type Style struct {
	App lipgloss.Style

	PlayInfoSection lipgloss.Style
	PlayInfoTrack   lipgloss.Style
	PlayInfoPadding lipgloss.Style

	PlaybarTrack    lipgloss.Style
	PlaybarProgress lipgloss.Style
	PlaybarHelp     lipgloss.Style
	PlaybarPadding  lipgloss.Style

	Footer lipgloss.Style
}

func DefaultStyles() *Style {
	s := new(Style)

	s.App = lipgloss.NewStyle().
		Padding(1, 2)

	s.PlayInfoSection = lipgloss.NewStyle().
		Padding(0, 1)

	s.PlayInfoTrack = lipgloss.NewStyle().
		Padding(0, 1).
		Background(lipgloss.Color("#0070C0")).
		Foreground(lipgloss.Color("#FFFFFF"))

	s.PlayInfoPadding = lipgloss.NewStyle().
		Background(lipgloss.Color("235")).
		Foreground(lipgloss.Color("243"))

	s.PlaybarTrack = lipgloss.NewStyle().
		Padding(0, 1)
		//Background(lipgloss.Color("#0070C0")).
		//Foreground(lipgloss.Color("#FFFFFF"))

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

	s.Footer = lipgloss.NewStyle().
		MarginTop(1).
		Padding(0, 1).
		Height(1)

	return s
}
