package models

import (
	"dalton.dog/YouTerm/resources"
	tea "github.com/charmbracelet/bubbletea"
)

var TrueMasterModel MasterModel

type MasterModel struct {
	mainMenu tea.Model
	curModel tea.Model
	user     *resources.User
}

func GetMasterModel() *MasterModel { return &TrueMasterModel }

func NewMasterModel(user *resources.User) {
	mainMenu := NewMainMenu(user)
	TrueMasterModel = MasterModel{
		mainMenu: mainMenu,
		curModel: mainMenu,
		user:     user,
	}
}

// TODO: Make it so that if this is passed nil, default to a main menu (if that's possible?)
// If that's not possible, I guess just make a ChangeToMainModel func instead
func (m *MasterModel) ChangeModel(newModel tea.Model) (tea.Model, tea.Cmd) {
	m.curModel = newModel
	return m.curModel, m.curModel.Init()
}

func (m *MasterModel) GoHome() (tea.Model, tea.Cmd) {
	m.curModel = m.mainMenu
	return m.curModel, m.curModel.Init()
}

func (m *MasterModel) Init() tea.Cmd { return m.curModel.Init() }

func (m *MasterModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		}
	}
	return m.curModel.Update(msg)
}

func (m *MasterModel) View() string { return m.curModel.View() }
