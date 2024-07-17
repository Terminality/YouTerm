package models

import (
	"dalton.dog/YouTerm/resources"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	mainMenuStyle = lipgloss.NewStyle()
)

type MainMenu struct {
	user   *resources.User
	width  int
	height int
}

func NewMainMenu(user *resources.User) *MainMenu {
	return &MainMenu{
		user: user,
	}
}

func (m *MainMenu) Init() tea.Cmd { return nil }

func (m *MainMenu) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case tea.KeyMsg:
		switch msg.String() {
		case "1":
			return GetMasterModel().ChangeModel(NewChannelList(m.user))
		case "0":
			return GetMasterModel().ChangeModel(NewAdminMenu(m.user))
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m *MainMenu) View() string {
	s := lipgloss.JoinHorizontal(lipgloss.Top, PlayButtonColor.Render(PlayButtonIcon), YouTermLogo)
	s += "\n\n Welcome, " + m.user.ID + "\n"
	s += "1. View User Channel List\n"
	s += "0. View Admin Menu\n"
	s += "\nPress q to Quit!"
	s = lipgloss.NewStyle().Align(lipgloss.Center).Render(s)
	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, s)
}

var PlayButtonColor = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF0000"))

const PlayButtonIcon = `
 ⠀⣠⣶⣶⢿⡿⣿⢿⡿⣿⢿⡿⣿⢿⡿⣿⢿⡿⣿⢿⡿⣶⣶⣄⠀⠀
 ⢰⣿⡽⣾⣻⣽⢯⣿⡽⣯⡿⣽⣯⢿⣽⢯⣿⡽⣯⡿⣽⡷⣯⢿⡇⠀
 ⣼⣯⣟⣷⢿⣽⣻⣞⣿⣳⠛⢿⣾⣻⡽⣿⣞⡿⣽⣻⢷⣟⡿⣯⣿⠀
 ⣿⣾⣿⣾⣿⣷⣿⣿⣾⣿⠀⠀⠘⠻⣿⣷⣿⣿⣿⣿⣿⣾⣿⣷⣿⠀
 ⣿⣷⣻⢷⣯⣟⣾⣽⢾⡿⠀⠀⠀⠀⠀⢉⣾⣯⢿⣳⣯⢿⣽⣻⢿⠀
 ⣿⣷⣻⣟⣾⣽⣳⣯⢿⣻⠀⠀⣀⣴⣾⣟⡿⣞⡿⣯⣟⣯⡿⣽⣿⠀
 ⢸⡷⣟⣾⣻⢾⣽⢯⣟⡿⣶⢿⣻⣟⣾⣽⣻⢯⡿⣷⢯⡿⣽⡷⡿⠀
 ⠘⣿⣻⢷⣻⣯⣟⣯⡿⣽⣻⣯⣟⣾⣽⢾⣯⣟⣿⣽⣻⣽⡷⣿⠇⠀
 ⠀⠈⠛⠻⠽⠾⠽⠷⠿⠿⠷⠿⠾⠿⠾⠿⠾⠽⠾⠷⠿⠷⠛⠁⠀⠀
`

const YouTermLogo = `
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⣿⣿⣿⣧⠀⠀⢀⣾⣿⣿⣿⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢰⣶⣶⣶⣶⣶⣶⣶⣶⣶⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠈⢿⣿⣿⣧⢀⣾⣿⣿⡿⠁⢀⣀⣤⣤⣤⣄⡀⠀⢠⣤⣤⣤⠀⢠⣤⣤⣤⠈⠉⠉⠉⣿⣿⡏⠉⠉⠉⠀⠀⢀⣀⣀⣀⠀⠀⠀⢀⣀⡀⢀⣀⡀⢀⣀⡀⢀⣀⣀⡀⠀⢀⣀⣀⡀⠀
⠀⠘⢿⣿⣿⣿⣿⣿⡿⠁⢰⣿⣿⣿⣿⣿⣿⣿⣆⢸⣿⣿⣿⠀⢸⣿⣿⣿⠀⠀⠀⠀⣿⣿⡇⠀⠀⠀⢀⣾⣿⣿⢿⣿⣿⣆⠀⢸⣿⣷⣿⣿⡇⢸⣿⣿⡿⢿⣿⣿⣾⡿⢿⣿⣿⡆
⠀⠀⠈⢿⣿⣿⣿⡿⠁⠀⣿⣿⣿⡏⠀⠹⣿⣿⣿⢸⣿⣿⣿⠀⢸⣿⣿⣿⠀⠀⠀⠀⣿⣿⡇⠀⠀⠀⣸⣿⣟⣀⣀⣘⣿⣿⠀⢸⣿⣟⠁⠀⠀⢸⣿⡏⠀⠀⢹⣿⡏⠀⠀⢹⣿⡧
⠀⠀⠀⢸⣿⣿⣿⠁⠀⠀⣿⣿⣿⣇⠀⣠⣿⣿⣿⢸⣿⣿⣿⠀⢸⣿⣿⣿⠀⠀⠀⠀⣿⣿⡇⠀⠀⠀⢻⣿⣟⠛⠛⠛⠛⠛⠀⢸⣿⣯⠀⠀⠀⢸⣿⡇⠀⠀⢸⣿⡇⠀⠀⢸⣿⡇
⠀⠀⠀⢸⣿⣿⣿⠀⠀⠀⠹⣿⣿⣿⣿⣿⣿⣿⠏⠘⣿⣿⣿⣿⣿⣿⣿⣿⠀⠀⠀⠀⣿⣿⡇⠀⠀⠀⠈⠿⣿⣶⣦⣶⣾⡦⠀⢸⣿⣷⠀⠀⠀⢸⣿⡇⠀⠀⢸⣿⣇⠀⠀⢸⣿⣇
⠀⠀⠀⠘⠛⠛⠛⠀⠀⠀⠀⠈⠙⠛⠛⠛⠋⠁⠀⠀⠈⠛⠛⠛⠛⠛⠛⠛⠀⠀⠀⠀⠉⠉⠁⠀⠀⠀⠀⠀⠈⠉⠉⠉⠁⠀⠀⠈⠉⠉⠀⠀⠀⠈⠉⠁⠀⠀⠈⠉⠉⠀⠀⠈⠉⠉
`
