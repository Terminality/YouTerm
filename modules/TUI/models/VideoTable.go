package models

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/evertras/bubble-table/table"
)

var ()

const ()

// Messages

type VideoTableModel struct {
	table table.Model
}

// TODO: Make this actually... load something
func MakeNewVideoTable() VideoTableModel {
	return VideoTableModel{}
}

func (m *VideoTableModel) Init() tea.Cmd { return nil }

func (m *VideoTableModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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

func (m *VideoTableModel) View() string { return m.table.View() }
