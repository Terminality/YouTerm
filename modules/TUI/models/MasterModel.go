package models

import (
	"fmt"
	"log"

	"dalton.dog/YouTerm/resources"
	tea "github.com/charmbracelet/bubbletea"
	//"github.com/charmbracelet/log"
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

func (m *MasterModel) ChangeModel(newModel tea.Model) (tea.Model, tea.Cmd) {
	log.Printf("Changing model from %T to %T\n", m.curModel, newModel)
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
	log.Print("MasterModel Update", "current", fmt.Sprintf("%T", m.curModel))
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
