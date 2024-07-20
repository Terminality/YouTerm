package models

import (
	"log"

	"dalton.dog/YouTerm/resources"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/evertras/bubble-table/table"
)

var ()

const (
	keyVideoID       string = "video_ID"
	keyVideoTitle    string = "video_title"
	keyVideoDesc     string = "video_description"
	keyVideoPubAt    string = "video_published_at"
	keyVideoChannel  string = "video_channel"
	keyVideoDuration string = "video_duration"
	keyVideoViews    string = "video_views"
	keyVideoLikes    string = "video_likes"
	keyVideoDislikes string = "video_dislikes"
	keyVideoComments string = "video_comments"
)

func makeVideoRow(video *resources.Video) table.Row {
	return table.NewRow(table.RowData{
		keyVideoID:    video.ID,
		keyVideoTitle: video.Title,
		// keyVideoDesc:     video.Description,
		keyVideoPubAt:    video.PublishedAt,
		keyVideoChannel:  video.ChannelTitle,
		keyVideoDuration: video.Duration,
		keyVideoViews:    video.ViewCount,
		keyVideoLikes:    video.LikeCount,
		keyVideoDislikes: video.DislikeCount,
		keyVideoComments: video.CommentCount,
	})
}

// Messages

type VideoTableModel struct {
	videoIDs  []string
	table     table.Model
	modelName string
}

// TODO: Make this actually... load something
func MakeNewVideoTable(videoIDs []string) *VideoTableModel {
	log.Printf("Initializing Video Table -- Input List: %v\n", videoIDs)
	newTable := table.New([]table.Column{
		table.NewColumn(keyVideoTitle, "Title", 30),
		table.NewColumn(keyVideoChannel, "Channel", 15),
		table.NewColumn(keyVideoDuration, "Duration", 10),
		table.NewColumn(keyVideoPubAt, "Published At", 15),
		table.NewColumn(keyVideoViews, "View Count", 10),
		table.NewColumn(keyVideoLikes, "Like Count", 10),
		table.NewColumn(keyVideoDislikes, "Dislike Count", 10),
	})

	var tableRows []table.Row
	for _, id := range videoIDs {
		video, err := resources.LoadOrCreateVideo(id)
		if err != nil {
			continue
		}
		tableRows = append(tableRows, makeVideoRow(video))
	}
	return &VideoTableModel{
		table:     newTable,
		videoIDs:  videoIDs,
		modelName: "Video Table",
	}
}

func (m *VideoTableModel) Init() tea.Cmd { return nil }

func (m *VideoTableModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	log.Println("Video Table Update")
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	m.table, cmd = m.table.Update(msg)
	cmds = append(cmds, cmd)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc", "q":
			cmds = append(cmds, tea.Quit)
		}
	}

	return m, tea.Batch(cmds...)
}

func (m *VideoTableModel) View() string { return m.table.View() }
