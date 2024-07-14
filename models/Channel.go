package models

import (
	"encoding/json"
	"errors"
	"fmt"

	"dalton.dog/YouTerm/modules/API"
	"dalton.dog/YouTerm/modules/Storage"
)

var (
	SUBBED_TO   string = "Subbed"
	WATCH_LATER string = "WatchLater"
)

type Channel struct {
	ID                string
	Bucket            string
	ChannelTitle      string
	Views             uint64
	Subscribers       uint64
	Videos            uint64
	UploadsPlaylistID string
}

// Implements bubbletea/list/listitem
func (c Channel) Title() string       { return c.ChannelTitle }
func (c Channel) Description() string { return fmt.Sprintf("ID: %s -- Views: %d", c.ID, c.Views) }
func (c Channel) FilterValue() string { return c.ChannelTitle }

// Impelements Storage.Resource
func (c *Channel) GetID() string                { return c.ID }
func (c *Channel) GetBucketName() string        { return c.Bucket }
func (c *Channel) MarshalData() ([]byte, error) { return json.Marshal(c) }

func (c *Channel) UnmarshalData(data []byte) *Channel {
	var output Channel
	json.Unmarshal(data, &output)
	return &output
}

func NewChannel(userID string, username string, userHandle string) (*Channel, error) {
	resp := API.RequestChannelFromAPI(userID, username, userHandle)

	if len(resp.Items) == 0 {
		return nil, errors.New("Couldn't load that channel!")
	}

	channelRsrc := resp.Items[0]

	var channel = Channel{
		ID:                channelRsrc.Id,
		Bucket:            Storage.CHANNELS,
		ChannelTitle:      channelRsrc.Snippet.Title,
		Views:             channelRsrc.Statistics.ViewCount,
		Subscribers:       channelRsrc.Statistics.SubscriberCount,
		Videos:            channelRsrc.Statistics.VideoCount,
		UploadsPlaylistID: channelRsrc.ContentDetails.RelatedPlaylists.Uploads,
	}
	channel.Save()
	return &channel, nil
}

func (c *Channel) Save() {
	Storage.SaveResource(c)
}

func (c *Channel) ToString() string { return "" }
