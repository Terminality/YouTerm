package TUI

import (
	"fmt"

	//"dalton.dog/YouTerm/API"
	"dalton.dog/YouTerm/modules/API"
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

type InputModel struct {
	textInput     textinput.Model
	channelToShow *API.Channel
	err           error
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

type item struct {
	username string
}

func (i item) Title() string       { return i.username }
func (i item) Description() string { return "" }
func (i item) FilterValue() string { return i.username }

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
		var channel *API.Channel = msg
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

func (im InputModel) Init() tea.Cmd {
	return textinput.Blink
}

func (im InputModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return im, tea.Quit
		case tea.KeyDown:
			im.textInput.SetValue("Oops!")
			return im, nil
		case tea.KeyEnter:
			if im.channelToShow != nil {
				im.channelToShow = nil
			} else {
				cmd = loadChannelFromAPI(im.textInput.Value())
				im.textInput.SetValue("")
			}
			return im, cmd
		}
		im.textInput, cmd = im.textInput.Update(msg)
		return im, cmd
	case errMsg:
		im.err = error(msg)
		return im, nil
	case channelMsg:
		im.channelToShow = msg
		return im, nil
	}

	return im, nil
}

func loadChannelFromAPI(username string) tea.Cmd {
	return func() tea.Msg {
		channel, err := API.NewChannelByUsername(username)
		if err != nil {
			return errMsg(err)
		}
		return channelMsg(channel)
	}
}

type channelMsg *API.Channel

func (im InputModel) View() string {
	if im.err != nil {
		return fmt.Sprintln("Uh oh, ran into an error!")
		//var channel = API.GetChannelStruct(im.usernameToShow)
		// return API.GetInfoByUsername(im.usernameToShow)
	} else if im.channelToShow != nil {
		return im.channelToShow.ToString()
	} else {
		return fmt.Sprintf(
			"What channel do you want to query the API for?\n\n%s\n\n%s",
			im.textInput.View(),
			"(esc to quit)",
		) + "\n"
	}
}
