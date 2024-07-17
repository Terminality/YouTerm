package models

import (
	"dalton.dog/YouTerm/resources"
	tea "github.com/charmbracelet/bubbletea"
)

type MainMenu struct {
	user *resources.User
}

func NewMainMenu(user *resources.User) *MainMenu {
	return &MainMenu{
		user: user,
	}
}

func (m *MainMenu) Init() tea.Cmd { return nil }

func (m *MainMenu) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m *MainMenu) View() string {
	return "Main menu :)"
}
