package models

import tea "github.com/charmbracelet/bubbletea"

type MasterModel struct {
	curModel tea.Model
}

func NewMasterModel(startingModel tea.Model) *MasterModel {
	return &MasterModel{
		curModel: startingModel,
	}
}

func (m *MasterModel) ChangeCurModel(newModel tea.Model) (tea.Model, tea.Cmd) {
	m.curModel = newModel
	return m.curModel, m.curModel.Init()
}

func (m *MasterModel) Init() tea.Cmd { return m.curModel.Init() }

func (m *MasterModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) { return m.curModel.Update(msg) }

func (m *MasterModel) View() string { return m.curModel.View() }
