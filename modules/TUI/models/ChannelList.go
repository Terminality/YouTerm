package TUI

import (
	"encoding/json"
	"fmt"
	"log"

	"dalton.dog/YouTerm/modules/Storage"
	"dalton.dog/YouTerm/resources"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	listHeight = 15
	listWidth  = 20
)

var (
	listStyle  = lipgloss.NewStyle().Padding(1, 2)
	titleStyle = lipgloss.NewStyle().MarginLeft(2)
)

type errMsg error
type channelMsg *resources.Channel
type listKeyMap struct {
	addItem    key.Binding
	removeItem key.Binding
	selectItem key.Binding
}

type ChannelListModel struct {
	listModel   list.Model
	active      bool
	keys        *listKeyMap
	user        *resources.User
	err         error
	inputModel  textinput.Model
	listShowing bool
}

func NewChannelList(user *resources.User) *ChannelListModel {

	var input = textinput.New()

	listItems := []list.Item{}
	for _, id := range user.GetList(resources.SUBBED_TO) {
		var channel *resources.Channel
		var err error
		bytes := Storage.LoadResource(Storage.CHANNELS, id)
		if bytes == nil {
			channel, err = resources.NewChannel(id, "", "")
		} else {
			err = json.Unmarshal(bytes, &channel)
		}
		checkErr(err)
		listItems = append(listItems, channel)

	}
	newList := list.New(listItems, list.NewDefaultDelegate(), 0, 0)
	newList.Title = fmt.Sprintf("%v's Channel List", user.GetID())
	newList.Styles.Title = titleStyle
	newList.SetStatusBarItemName("channel", "channels")

	newModel := &ChannelListModel{
		user:        user,
		active:      false,
		keys:        newKeyMap(),
		listModel:   newList,
		listShowing: true,
		inputModel:  input,
	}
	return newModel
}

func (m ChannelListModel) Init() tea.Cmd { return nil }

func (m ChannelListModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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

func (m ChannelListModel) View() string {
	if m.listShowing {
		return listStyle.Render(m.listModel.View())
	} else {
		return m.inputModel.View()
	}
}

func newKeyMap() *listKeyMap {
	return &listKeyMap{
		addItem: key.NewBinding(
			key.WithKeys("a"),
			key.WithHelp("a", "add item"),
		),
		removeItem: key.NewBinding(
			key.WithKeys("x"),
			key.WithHelp("x", "delete item"),
		),
		selectItem: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "choose"),
		),
	}
}

// TODO: Set up username->ID mapping (or just another username->Channel mapping?)
func loadChannelFromAPI(username string) tea.Cmd {
	return func() tea.Msg {
		channel, err := resources.NewChannel("", username, "")
		if err != nil {
			return errMsg(err)
		}
		return channelMsg(channel)
	}
}

// func loadChannel(username string) tea.Cmd {
// 	return func() tea.Msg {
// 		channel, err := Storage.LoadResource(Storage.CHANNELS, )
// 	}
// }

func checkErr(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
