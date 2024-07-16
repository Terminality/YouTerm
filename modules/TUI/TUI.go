package TUI

import (
	"dalton.dog/YouTerm/modules/TUI/models"
	tea "github.com/charmbracelet/bubbletea"
	//	gloss "github.com/charmbracelet/lipgloss"
	"dalton.dog/YouTerm/resources"
	table "github.com/evertras/bubble-table/table"
	//"github.com/charmbracelet/log"
)

type ChannelModel struct {
	table table.Model
	user  *resources.User
}

func MakeNewProgram(user *resources.User) tea.Program {
	return *tea.NewProgram(models.NewChannelList(user))
}

func (m ChannelModel) Init() tea.Cmd { return nil }

func (m ChannelModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	m.table, cmd = m.table.Update(msg)
	cmds = append(cmds, cmd)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc", "q":
			cmds = append(cmds, tea.Quit)
		}
	}
	return m, tea.Batch(cmds...)
}

func (m ChannelModel) View() string { return m.table.View() }
