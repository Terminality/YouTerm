package API

import (
	"errors"
	"fmt"
)

type Channel struct {
	ID                string
	title             string
	totalViews        uint64
	username          string
	uploadsPlaylistID string
	uploadedVideoIDs  []string
}

func (c Channel) Title() string       { return c.title }
func (c Channel) Description() string { return fmt.Sprintf("ID: %s -- Views: %d", c.ID, c.totalViews) }
func (c Channel) FilterValue() string { return c.title }

func NewChannelByUsername(username string) (*Channel, error) {

	resp := GetInfoByUsername(username)

	if len(resp.Items) == 0 {
		return nil, errors.New("Couldn't load that channel!")
	}

	var channel = Channel{
		ID:         resp.Items[0].Id,
		title:      resp.Items[0].Snippet.Title,
		totalViews: resp.Items[0].Statistics.ViewCount,
		username:   username,
	}
	return &channel, nil
}

func (c *Channel) ToString() string {
	return fmt.Sprintf("This channel's ID is %s. Its title is '%s', "+
		"and it has %d views.",
		c.ID, c.title, c.totalViews)
}

func UpdateUploadList() {

}

func PrintRecentUploads() {

}
