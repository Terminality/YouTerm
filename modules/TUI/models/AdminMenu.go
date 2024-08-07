package models

import (
	"fmt"
	"io"
	"log"

	"strings"

	"dalton.dog/YouTerm/modules/Storage"
	"dalton.dog/YouTerm/resources"
	"dalton.dog/YouTerm/utils"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))
	paginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	helpStyle         = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
	quitTextStyle     = lipgloss.NewStyle().Margin(1, 0, 2, 4)
)

type item string

func (i item) FilterValue() string { return "" }

type itemDelegate struct{}

func (d itemDelegate) Height() int                             { return 1 }
func (d itemDelegate) Spacing() int                            { return 0 }
func (d itemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(item)
	if !ok {
		return
	}

	str := fmt.Sprintf("%d. %s", index+1, i)

	fn := itemStyle.Render
	if index == m.Index() {
		fn = func(s ...string) string {
			return selectedItemStyle.Render("> " + strings.Join(s, " "))
		}
	}

	fmt.Fprint(w, fn(str))
}

type AdminMenu struct {
	list      list.Model
	user      *resources.User
	status    string
	modelName string
}

func NewAdminMenu(user *resources.User) *AdminMenu {
	log.Printf("Initializing Admin Menu -- User: %v\n", user.GetID())
	items := []list.Item{
		item("Clear Channels Bucket"),
		item("Clear Videos Bucket"),
		item("Clear User Subbed List"),
	}

	w, h, err := utils.GetTerminalSize()
	utils.CheckErrFatal(err, "Couldn't get terminal size")

	list := list.New(items, itemDelegate{}, w, h)
	list.Title = "Hit a button to do the admin thing"
	// list.Styles.Title = titleStyle
	list.Styles.PaginationStyle = paginationStyle
	list.Styles.HelpStyle = helpStyle

	menu := &AdminMenu{
		list:      list,
		user:      user,
		modelName: "Admin Menu",
	}

	return menu
}

func (m AdminMenu) Init() tea.Cmd { return nil }

func (m AdminMenu) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "1":
			Storage.ClearBucket(Storage.CHANNELS)
			m.status = "Channels bucket cleared!"
			return m, nil
		case "2":
			Storage.ClearBucket(Storage.VIDEOS)
			m.status = "Videos bucket cleared!"
			return m, nil
		case "3":
			m.user.UserLists[resources.SUBBED_TO] = map[string]bool{}
			m.status = "Cleared user sub list"
		case "0", "q", "ctrl+c":
			return GetMasterModel().GoHome()
		}
	}
	return m, nil
}

func (m AdminMenu) View() string {
	return m.list.View() + "\n" + m.status
}
