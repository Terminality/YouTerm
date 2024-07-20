package models

// TODO: Figure out how to make the terminal get the size and set list dimensions on launch, not just resize

import (
	"fmt"
	"log"

	"dalton.dog/YouTerm/resources"
	"dalton.dog/YouTerm/utils"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
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

	newModel := &ChannelListModel{
		user:      user,
		keys:      keyMap,
		listModel: newList,
		modelName: "Channel List",
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
		case key.Matches(msg, m.keys.addItem):
			return m, nil
		case key.Matches(msg, m.keys.launchItem):
			if channel, ok := selectedItem.(resources.Channel); ok {
				utils.LaunchURL(fmt.Sprintf("https://youtube.com/channel/%v", channel.ID))
			}

			return m, nil
		case key.Matches(msg, m.keys.removeItem):
			i := m.listModel.Index()
			m.listModel.RemoveItem(i)
			if channel, ok := selectedItem.(resources.Channel); ok {
				m.user.RemoveFromList(resources.SUBBED_TO, channel.ID)
			}
			return m, nil
		case key.Matches(msg, m.keys.selectItem):
			return m, nil
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

	m.listModel, cmd = m.listModel.Update(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m ChannelListModel) View() string {
	return m.listModel.View()
}

func newKeyMap() *listKeyMap {
	return &listKeyMap{
		addItem: key.NewBinding(
			key.WithKeys("a"),
			key.WithHelp("a", "add item"),
		),
		launchItem: key.NewBinding(
			key.WithKeys("l"),
			key.WithHelp("l", "launch"),
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
