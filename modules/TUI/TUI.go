package TUI

import (
	"dalton.dog/YouTerm/modules/TUI/models"
	tea "github.com/charmbracelet/bubbletea"
	//	gloss "github.com/charmbracelet/lipgloss"
	"dalton.dog/YouTerm/resources"
	//"github.com/charmbracelet/log"
)

func makeNewProgram(user *resources.User) tea.Program {
	// return *tea.NewProgram(models.NewChannelList(user))
	models.NewMasterModel(user)
	return *tea.NewProgram(models.GetMasterModel(), tea.WithAltScreen())
}

func RunProgram(user *resources.User) error {
	program := makeNewProgram(user)
	if _, err := program.Run(); err != nil {
		return err
	}
	return nil
}
