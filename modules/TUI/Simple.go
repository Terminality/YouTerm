package TUI

import (
	"encoding/json"
	"fmt"
	"log"

	"dalton.dog/YouTerm/modules/Storage"
	"dalton.dog/YouTerm/resources"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

var listStyle = lipgloss.NewStyle().Padding(1, 2)

type errMsg error
type channelMsg *resources.Channel

type model struct {
	inputModel  textinput.Model
	listShowing bool
	listModel   list.Model
	user        *resources.User
	err         error
}

func newModel(user *resources.User) tea.Model {
	var input = textinput.New()
	input.Placeholder = "Northernlion"
	listItems := []list.Item{}
	for _, id := range user.GetList(resources.SUBBED_TO) {
		var channel *resources.Channel
		var err error
		bytes := Storage.LoadResource("Channels", id)
		if bytes == nil {
			channel, err = resources.NewChannel(id, "", "")
		} else {
			err = json.Unmarshal(bytes, &channel)
		}
		checkErr(err)
		listItems = append(listItems, channel)

	}
	list := list.New(listItems, list.NewDefaultDelegate(), 0, 0)
	list.Title = fmt.Sprintf("%v's Channel List", user.GetID())

	model := model{
		inputModel:  input,
		listShowing: true,
		listModel:   list,
		user:        user,
	}

	return model
}

func NewPromptProgram(user *resources.User) *tea.Program {
	return tea.NewProgram(newModel(user), tea.WithAltScreen())
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		h, v := listStyle.GetFrameSize()
		m.listModel.SetSize(msg.Width-h, msg.Height-v)
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlA:
			if !m.listShowing {
				break
			}
			m.listShowing = false
			m.inputModel.Focus()
			return m, nil
		case tea.KeyEnter:
			if m.listShowing {
				break
			}
			cmd = loadChannelFromAPI(m.inputModel.Value())
			m.inputModel.SetValue("")
			m.listShowing = true
			m.inputModel.Blur()
			return m, cmd
		}

	case errMsg:
		m.err = msg
		return m, nil
	case channelMsg:
		var channel *resources.Channel = msg
		if m.user.AddToList(resources.SUBBED_TO, channel.GetID()) {
			cmd = m.listModel.InsertItem(0, channel)
		}
		return m, cmd
	}

	m.inputModel, cmd = m.inputModel.Update(msg)
	cmds = append(cmds, cmd)
	m.listModel, cmd = m.listModel.Update(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	if m.listShowing {
		return listStyle.Render(m.listModel.View())
	} else {
		return m.inputModel.View()
	}
}

func loadChannelFromAPI(username string) tea.Cmd {
	return func() tea.Msg {
		channel, err := resources.NewChannel("", username, "")
		if err != nil {
			return errMsg(err)
		}
		return channelMsg(channel)
	}
}
