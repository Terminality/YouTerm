package resources

import (
	"encoding/json"
	"errors"
	"fmt"

	"dalton.dog/YouTerm/modules/API"
	"dalton.dog/YouTerm/modules/Storage"
	"github.com/evertras/bubble-table/table"
)

const (
	SUBBED_TO   string = "Subbed"
	WATCH_LATER string = "WatchLater"

	keyChannelID     string = "channel_ID"
	keyChannelTitle  string = "channel_title"
	keyChannelViews  string = "channel_views"
	keyChannelSubs   string = "channel_subs"
	keyChannelVideos string = "channel_videos"
)

func MakeChannelTable() table.Model {
	model := table.New([]table.Column{
		table.NewColumn(keyChannelTitle, "Title", 10),
		table.NewColumn(keyChannelViews, "Views", 7),
		table.NewColumn(keyChannelSubs, "Subs", 7),
		table.NewColumn(keyChannelVideos, "Videos", 8),
		table.NewColumn(keyChannelID, "ID", 20),
	})

	return model
}

type Channel struct {
	ID                string
	Bucket            string
	ChannelTitle      string
	TotalViewCount    uint64
	SubCount          uint64
	VideoCount        uint64
	UploadsPlaylistID string
	LoadedUploadIDs   []string
}

// Implements bubbletea list item
func (c Channel) Title() string { return c.ChannelTitle }
func (c Channel) Description() string {
	return fmt.Sprintf("ID: %s -- Views: %d", c.ID, c.TotalViewCount)
}
func (c Channel) FilterValue() string { return c.ChannelTitle }

// Impelements Storage.Resource
func (c *Channel) GetID() string                { return c.ID }
func (c *Channel) GetBucketName() string        { return c.Bucket }
func (c *Channel) MarshalData() ([]byte, error) { return json.Marshal(c) }

func LoadOrCreateChannel(userID string, username string, userHandle string) (*Channel, error) {
	bytes := Storage.LoadResource(Storage.CHANNELS, userID)

	if bytes == nil {
		return NewChannel(userID, username, userHandle)
	}

	var channel *Channel
	json.Unmarshal(bytes, &channel)
	return channel, nil

}

func NewChannel(userID string, username string, userHandle string) (*Channel, error) {
	resp := API.RequestChannel(userID, username, userHandle)

	if len(resp.Items) == 0 {
		return nil, errors.New("Couldn't load that channel!")
	}

	channelRsrc := resp.Items[0]

	var channel = Channel{
		ID:                channelRsrc.Id,
		Bucket:            Storage.CHANNELS,
		ChannelTitle:      channelRsrc.Snippet.Title,
		TotalViewCount:    channelRsrc.Statistics.ViewCount,
		SubCount:          channelRsrc.Statistics.SubscriberCount,
		VideoCount:        channelRsrc.Statistics.VideoCount,
		UploadsPlaylistID: channelRsrc.ContentDetails.RelatedPlaylists.Uploads,
	}
	channel.Save()
	return &channel, nil
}

func (c *Channel) Save() { Storage.SaveResource(c) }

func (c *Channel) MakeRow() table.Row {
	return table.NewRow(table.RowData{
		keyChannelID:     c.ID,
		keyChannelSubs:   c.SubCount,
		keyChannelTitle:  c.ChannelTitle,
		keyChannelViews:  c.TotalViewCount,
		keyChannelVideos: c.VideoCount,
	})
}

func (c *Channel) ToString() string { return "" }
