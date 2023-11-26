package ui

import (
	"github.com/a-yee/spot/app"
	tea "github.com/charmbracelet/bubbletea"
)

type AppModel struct {
	quitting bool
	error    error
	app      app.AppInfo
	//footer
	//showFooter
}

func NewAppModel(ai app.AppInfo) AppModel {
	return AppModel{
		app: ai,
	}
}

func (m AppModel) Init() tea.Cmd {
	return nil
}

func (m AppModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return nil, nil
}

func (m AppModel) View() string {
	return ""
}
