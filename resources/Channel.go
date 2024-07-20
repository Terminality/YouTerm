package resources

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

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

type Channel struct {
	ID                string
	Bucket            string
	ChannelTitle      string
	TotalViewCount    uint64
	SubCount          uint64
	VideoCount        uint64
	UploadsPlaylistID string
	LoadedUploadIDs   map[string]bool
	LastUploadPageID  string
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
	channel.LoadUploads()
	return channel, nil

}

func NewChannel(userID string, username string, userHandle string) (*Channel, error) {
	log.Printf("Making new Channel")
	resp, err := API.RequestChannel(userID, username, userHandle)

	if err != nil {
		return nil, err
	}

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
		LoadedUploadIDs:   map[string]bool{},
		LastUploadPageID:  "",
	}
	channel.Save()
	go channel.LoadUploads()
	return &channel, nil
}

func (c *Channel) LoadUploads() {
	resp, err := API.RequestPlaylistContents(c.UploadsPlaylistID, c.LastUploadPageID)
	if err != nil {
		log.Fatalf("Couldn't load playlist contents: %v", err)
	}

	for _, playlistItem := range resp.Items {
		videoID := playlistItem.ContentDetails.VideoId
		_, exists := c.LoadedUploadIDs[videoID]
		if !exists {
			c.LoadedUploadIDs[videoID] = true
			_, err := NewVideo(videoID)
			if err != nil {
				log.Fatalf("Couldn't load video: %v", err)
			}
		}
	}
	c.Save()
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
