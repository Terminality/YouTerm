package TUI

import (
	"encoding/json"

	tea "github.com/charmbracelet/bubbletea"
	//	gloss "github.com/charmbracelet/lipgloss"
	"dalton.dog/YouTerm/modules/Storage"
	"dalton.dog/YouTerm/resources"
	table "github.com/evertras/bubble-table/table"
	//"github.com/charmbracelet/log"
)

// NOTE: Consider moving `models/ChannelList.go` to TUI package

type ChannelModel struct {
	table table.Model
	user  *models.User
}

func MakeNewChannelModel(user *models.User) tea.Model {
	var tableRows []table.Row
	for _, id := range user.GetList(models.SUBBED_TO) {
		var channel *models.Channel
		var err error
		bytes := Storage.LoadResource("Channels", id)
		if bytes == nil {
			channel, err = models.NewChannel(id, "", "")
		} else {
			err = json.Unmarshal(bytes, &channel)
		}
		checkErr(err)
		tableRows = append(tableRows, channel.MakeRow())
	}
	return ChannelModel{
		table: models.MakeChannelTable().WithRows(tableRows),
		user:  user,
	}
}

func MakeNewProgram(user *models.User) tea.Program {
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
