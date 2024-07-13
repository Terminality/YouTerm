package models

import (
	"dalton.dog/YouTerm/modules/API"
	"errors"
	"fmt"
)

type Channel struct {
	ID                string
	bucket            string
	title             string
	views             uint64
	subscribers       uint64
	videos            uint64
	uploadsPlaylistID string
}

func (c Channel) Title() string       { return c.title }
func (c Channel) Description() string { return fmt.Sprintf("ID: %s -- Views: %d", c.ID, c.views) }
func (c Channel) FilterValue() string { return c.title }

func NewChannel(userID string, username string, userHandle string) (*Channel, error) {
	resp := API.RequestChannelFromAPI(userID, username, userHandle)

	if len(resp.Items) == 0 {
		return nil, errors.New("Couldn't load that channel!")
	}

	channelRsrc := resp.Items[0]

	var channel = Channel{
		ID:                channelRsrc.Id,
		bucket:            "Channels",
		title:             channelRsrc.Snippet.Title,
		views:             channelRsrc.Statistics.ViewCount,
		subscribers:       channelRsrc.Statistics.SubscriberCount,
		videos:            channelRsrc.Statistics.VideoCount,
		uploadsPlaylistID: channelRsrc.ContentDetails.RelatedPlaylists.Uploads,
	}
	return &channel, nil
}

func (c *Channel) Save() {

}

func (c *Channel) ToString() string { return "" }
