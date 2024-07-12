package TUI

import (
	"fmt"

	//"dalton.dog/YouTerm/API"
	"dalton.dog/YouTerm/API"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type errMsg error

type InputModel struct {
	textInput     textinput.Model
	channelToShow *API.Channel
	err           error
}

func newModel() tea.Model {
	var input = textinput.New()
	input.Placeholder = "Northernlion"
	input.Focus()

	return InputModel{
		textInput:     input,
		channelToShow: nil,
	}
}

func NewPromptProgram() *tea.Program {
	return tea.NewProgram(newModel())
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
