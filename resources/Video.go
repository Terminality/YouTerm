package resources

import (
	"encoding/json"
	"errors"
	"log"

	"dalton.dog/YouTerm/modules/API"
	"dalton.dog/YouTerm/modules/Storage"
	"github.com/evertras/bubble-table/table"
)

type Video struct {
	ID     string
	Bucket string
	Title  string
	// Description  string
	PublishedAt  string
	ChannelID    string
	ChannelTitle string
	Duration     string

	ViewCount    uint64
	LikeCount    uint64
	DislikeCount uint64
	CommentCount uint64
}

// Impelements Storage.Resource
func (v *Video) GetID() string                { return v.ID }
func (v *Video) GetBucketName() string        { return v.Bucket }
func (v *Video) MarshalData() ([]byte, error) { return json.Marshal(v) }

func LoadOrCreateVideo(videoID string) (*Video, error) {
	bytes := Storage.LoadResource(Storage.VIDEOS, videoID)
	if bytes == nil {
		return NewVideo(videoID)
	}

	var video *Video
	json.Unmarshal(bytes, &video)
	return video, nil
}
func NewVideo(videoID string) (*Video, error) {
	log.Printf("Creating video -- ID: %v", videoID)
	resp, err := API.RequestVideo(videoID)

	if err != nil {
		return nil, err
	}

	if len(resp.Items) == 0 {
		return nil, errors.New("Couldn't load that video!")
	}

	videoRsrc := resp.Items[0]

	video := Video{
		ID:     videoID,
		Bucket: Storage.VIDEOS,
		Title:  videoRsrc.Snippet.Title,
		// Description:  videoRsrc.Snippet.Description,
		PublishedAt:  videoRsrc.Snippet.PublishedAt,
		ChannelID:    videoRsrc.Snippet.ChannelId,
		ChannelTitle: videoRsrc.Snippet.ChannelTitle,
		Duration:     videoRsrc.ContentDetails.Duration,

		ViewCount:    videoRsrc.Statistics.ViewCount,
		LikeCount:    videoRsrc.Statistics.LikeCount,
		DislikeCount: videoRsrc.Statistics.DislikeCount,
		CommentCount: videoRsrc.Statistics.CommentCount,
	}

	video.Save()
	return &video, nil
}

func (v *Video) ToRow() table.Row {
	return table.NewRow(table.RowData{})
}

func (v *Video) Save() {
	Storage.SaveResource(v)
}
