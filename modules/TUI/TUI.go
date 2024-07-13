package TUI

import (
	tea "github.com/charmbracelet/bubbletea"
	//	gloss "github.com/charmbracelet/lipgloss"
	table "github.com/evertras/bubble-table/table"
)

func MakeProgram() tea.Program {
	return *tea.NewProgram(makeNewTableModel())
}

func makeNewTableModel() tea.Model {
	return VideoTableModel{
		videoTable: table.New([]table.Column{
			table.NewColumn("name", "Name", 15),
			table.NewColumn("animal", "Animal", 20),
		}).WithRows([]table.Row{
			makeRow("Jasper", "Dog"),
			makeRow("Dahlia", "Cat"),
			makeRow("Alphonse", "Cat"),
		}),
	}
}

func makeRow(name string, animal string) table.Row {
	return table.NewRow(table.RowData{
		"name":   name,
		"animal": animal,
	})
}

type VideoTableModel struct {
	videoTable table.Model
}

func (m VideoTableModel) Init() tea.Cmd {
	return nil
}

func (m VideoTableModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	m.videoTable, cmd = m.videoTable.Update(msg)
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

func (m VideoTableModel) View() string {
	return m.videoTable.View()
}
