package TUI

import (
	"dalton.dog/YouTerm/models"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var listStyle = lipgloss.NewStyle().Padding(1, 2)

type errMsg error

type model struct {
	inputModel  textinput.Model
	listShowing bool
	listModel   list.Model
	err         error
}

func newModel() tea.Model {
	var input = textinput.New()
	input.Placeholder = "Northernlion"
	list := list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 0)
	// list.Title = "Subscribed Channels"

	model := model{
		inputModel:  input,
		listShowing: true,
		listModel:   list,
	}

	return model
}

func NewPromptProgram() *tea.Program {
	return tea.NewProgram(newModel(), tea.WithAltScreen())
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
		var channel *models.Channel = msg
		cmd = m.listModel.InsertItem(0, channel)
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
		channel, err := models.NewChannel("", username, "")
		if err != nil {
			return errMsg(err)
		}
		return channelMsg(channel)
	}
}

type channelMsg *models.Channel
