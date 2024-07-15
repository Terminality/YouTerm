package TUI

import (
	"encoding/json"

	"dalton.dog/YouTerm/modules/TUI/models"
	tea "github.com/charmbracelet/bubbletea"
	//	gloss "github.com/charmbracelet/lipgloss"
	"dalton.dog/YouTerm/modules/Storage"
	"dalton.dog/YouTerm/resources"
	table "github.com/evertras/bubble-table/table"
	//"github.com/charmbracelet/log"
)

type ChannelModel struct {
	table table.Model
	user  *resources.User
}

func MakeNewChannelModel(user *resources.User) tea.Model {
	var tableRows []table.Row
	for _, id := range user.GetList(resources.SUBBED_TO) {
		var channel *resources.Channel
		var err error
		bytes := Storage.LoadResource("Channels", id)
		if bytes == nil {
			channel, err = resources.NewChannel(id, "", "")
		} else {
			err = json.Unmarshal(bytes, &channel)
		}
		checkErr(err)
		tableRows = append(tableRows, channel.MakeRow())
	}
	return ChannelModel{
		table: resources.MakeChannelTable().WithRows(tableRows),
		user:  user,
	}
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
