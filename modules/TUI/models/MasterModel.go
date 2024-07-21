package models

import (
	"log"

	"dalton.dog/YouTerm/resources"
	tea "github.com/charmbracelet/bubbletea"
	//"github.com/charmbracelet/log"
)

var TrueMasterModel MasterModel

type MasterModel struct {
	history  []tea.Model
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

func (m *MasterModel) PushToHistory(modelPtr *tea.Model) {
	newModel := *modelPtr
	log.Printf("Pushing to history: %T", newModel)
	m.history = append(m.history, newModel)
}

func (m *MasterModel) popFromHistory() tea.Model {
	l := len(m.history)
	if l <= 0 {
		return nil
	}
	outModel := m.history[l-1]
	log.Printf("Popping from history: %T", outModel)
	m.history[l-1] = nil
	m.history = m.history[:l]
	return outModel
}

func (m *MasterModel) GoBack(msg tea.Msg) (tea.Model, tea.Cmd) {
	prevModel := m.popFromHistory()
	if prevModel == nil {
		return m.GoHome()
	}
	log.Printf("Going back to previous model: %T", prevModel)
	m.curModel = prevModel
	return m, passMsgCmd(msg)
}

func passMsgCmd(msg tea.Msg) tea.Cmd {
	return func() tea.Msg {
		return msg
	}
}

func (m *MasterModel) ChangeModel(newModel tea.Model) (tea.Model, tea.Cmd) {
	log.Printf("Changing model from %T to %T\n", m.curModel, newModel)
	m.PushToHistory(&m.curModel)
	m.curModel = newModel
	return m.curModel, m.curModel.Init()
}

func (m *MasterModel) GoHome() (tea.Model, tea.Cmd) {
	log.Println("Going back to main menu")
	m.curModel = m.mainMenu
	return m.curModel, m.curModel.Init()
}

func (m *MasterModel) Init() tea.Cmd { return m.curModel.Init() }

func (m *MasterModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	log.Printf("MasterModel Update: %T", m.curModel)
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		}
	}

	m.curModel, cmd = m.curModel.Update(msg)
	return m, cmd
}

func (m *MasterModel) View() string {
	if m.curModel != nil {
		return m.curModel.View()
	}
	return ""
}
