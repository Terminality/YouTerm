package API

import (
	"context"
	"errors"
	"fmt"
	"log"

	"google.golang.org/api/youtube/v3"
)

var masterAPI *apiManager

type apiManager struct {
	apiKey  string
	service *youtube.Service
}

func InitializeManager() error {
	service, err := GetAuthenticatedService(context.Background())
	if err != nil {
		return err
	}

	masterAPI = &apiManager{
		service: service,
	}

	log.Println("Initialized API manager")

	return nil
}

func RequestVideo(videoID string) (*youtube.VideoListResponse, error) {
	log.Printf("API Request (Video) -- ID: %v\n", videoID)

	call := masterAPI.service.Videos.List([]string{"snippet", "contentDetails", "statistics"})
	call = call.Id(videoID)
	resp, err := call.Do()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error making API call to load video: %v", err))
	}

	return resp, nil
}

func RequestChannel(userID string, username string, handle string) (*youtube.ChannelListResponse, error) {
	log.Printf("API Request (Channel) -- ID: %v -- Username: %v -- Handle: %v\n", userID, username, handle)

	call := masterAPI.service.Channels.List([]string{"snippet", "contentDetails", "statistics"})

	if userID != "" {
		call = call.Id(userID)
	} else if username != "" {
		call = call.ForUsername(username)
	} else if handle != "" {
		call = call.ForHandle(handle)
	}

	resp, err := call.Do()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error making API call: %v", err))
	}

	return resp, nil
}

// TODO: Make this actually account for pageID to load additional results
func RequestPlaylistContents(playlistID string, pageID string) (*youtube.PlaylistItemListResponse, error) {
	log.Printf("API Request (Playlist) -- ID: %v -- Page ID: %v\n", playlistID, pageID)
	uploadsCall := masterAPI.service.PlaylistItems.List([]string{"contentDetails"}).PlaylistId(playlistID).MaxResults(50)
	uploadsResponse, err := uploadsCall.Do()
	if err != nil {
		return nil, errors.New("Uploads couldn't be obtained from uploads playlist")
	}

	return uploadsResponse, nil
}
