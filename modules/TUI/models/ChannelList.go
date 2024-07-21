package models

// TODO: Figure out how to make the terminal get the size and set list dimensions on launch, not just resize

import (
	"fmt"
	"log"

	"dalton.dog/YouTerm/resources"
	"dalton.dog/YouTerm/utils"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	listHeight = 50
	listWidth  = 80
)

var (
	listStyle = lipgloss.NewStyle().Padding(1, 2)
)

type errMsg error
type channelMsg *resources.Channel
type listKeyMap struct {
	addItem    key.Binding
	removeItem key.Binding
	selectItem key.Binding
	launchItem key.Binding
}

type ChannelListModel struct {
	inputModel  textinput.Model
	inputActive bool

	listModel list.Model
	keys      *listKeyMap
	user      *resources.User
	err       error
	modelName string
}

func NewChannelList(user *resources.User) *ChannelListModel {
	log.Printf("Initializing Channel List -- User: %v\n", user.GetID())

	listItems := []list.Item{}
	for id := range user.GetList(resources.SUBBED_TO) {
		channel, err := resources.LoadOrCreateChannel(id, "", "")
		checkErr(err)
		listItems = append(listItems, channel)
	}

	newList := list.New(listItems, list.NewDefaultDelegate(), listWidth, listHeight)
	newList.Title = fmt.Sprintf("%v's Channel List", user.GetID())
	// newList.Styles.Title = titleStyle
	newList.SetStatusBarItemName("channel", "channels")
	keyMap := newKeyMap()

	newList.AdditionalShortHelpKeys = func() []key.Binding {
		return []key.Binding{
			keyMap.addItem,
			keyMap.removeItem,
			keyMap.selectItem,
			keyMap.launchItem,
		}
	}
	inputModel := textinput.New()
	inputModel.Placeholder = "Channel to sub to"

	newModel := &ChannelListModel{
		user:        user,
		keys:        keyMap,
		listModel:   newList,
		modelName:   "Channel List",
		inputActive: false,
		inputModel:  inputModel,
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
		selectedItem := m.listModel.SelectedItem()
		switch {
		case key.Matches(msg, m.keys.addItem) && !m.inputActive:
			m.inputActive = true
			m.inputModel.Focus()
			return m, nil
		case key.Matches(msg, m.keys.launchItem) && !m.inputActive:
			channel := selectedItem.(*resources.Channel)
			utils.LaunchURL(fmt.Sprintf("https://youtube.com/channel/%v", channel.ID))

			return m, nil

		case key.Matches(msg, m.keys.removeItem) && !m.inputActive:
			i := m.listModel.Index()
			m.listModel.RemoveItem(i)
			channel := selectedItem.(*resources.Channel)
			m.user.RemoveFromList(resources.SUBBED_TO, channel.ID)
			return m, nil

		case key.Matches(msg, m.keys.selectItem):
			if m.inputActive {
				value := m.inputModel.Value()
				m.inputActive = false
				m.inputModel.Blur()
				if value == "" {
					return m, nil
				}
				cmd = loadChannelFromAPI(value)
				m.inputModel.SetValue("")
				return m, cmd
			}
			return m, nil
		}

	case errMsg:
		m.err = msg
		return m, nil

	case channelMsg:
		var channel *resources.Channel = msg
		log.Printf("Channel message received for %v\n", channel.ChannelTitle)
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
	if m.inputActive {
		return m.inputModel.View()
	} else {
		return m.listModel.View()
	}
}

func newKeyMap() *listKeyMap {
	return &listKeyMap{
		addItem: key.NewBinding(
			key.WithKeys("a"),
			key.WithHelp("a", "add item"),
		),
		launchItem: key.NewBinding(
			key.WithKeys("b"),
			key.WithHelp("b", "browser"),
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
		channel, err := resources.LoadOrCreateChannel("", username, "")
		if err != nil {
			return errMsg(err)
		}
		log.Printf("Loaded channel (%v), returning as channelMsg\n", channel.ChannelTitle)
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
