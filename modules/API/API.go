package API

import (
	"context"
	"errors"
	"log"
	"os"

	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

var masterAPI *apiManager

type apiManager struct {
	ctx     context.Context
	apiKey  string
	service *youtube.Service
}

func InitializeManager() error {
	ctx := context.Background()
	APIKey, err := getKeyFromEnv()

	if err != nil {
		return err
	}

	service, err := getServiceFromAPI(ctx, APIKey)

	if err != nil {
		return err
	}

	masterAPI = &apiManager{
		ctx:     ctx,
		apiKey:  APIKey,
		service: service,
	}

	return nil
}

func getKeyFromEnv() (string, error) {
	api := os.Getenv("YOUTERM_API_KEY")
	if api != "" {
		return api, nil
	}
	// TODO: Prompt user for API key, then save it to env variable (os.Setenv())

	return "", errors.New("Couldn't load YOUTERM_API_KEY from environment variable.")
}

func getServiceFromAPI(ctx context.Context, APIkey string) (*youtube.Service, error) {
	service, err := youtube.NewService(ctx, option.WithAPIKey(APIkey))
	if err != nil {
		return nil, errors.New("Error creating new YouTube service")
	}

	return service, nil
}

func RequestVideo(videoID string) (*youtube.VideoListResponse, error) {
	call := masterAPI.service.Videos.List([]string{"snippet", "contentDetails", "statistics"})
	call = call.Id(videoID)
	resp, err := call.Do()
	if err != nil {
		return nil, errors.New("Unable to load video")
	}

	return resp, nil
}

// TODO: Make this return an error like the others
func RequestChannel(userID string, username string, handle string) *youtube.ChannelListResponse {
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
		log.Fatalf("Error making API call: %v", err)
	}

	return resp
}

func RequestPlaylistContents(playlistID string, pageID string) ([]string, error) {
	// TODO: Make this actually account for pageID to load additional results
	uploadsCall := masterAPI.service.PlaylistItems.List([]string{"contentDetails"}).PlaylistId(playlistID).MaxResults(20)
	uploadsResponse, err := uploadsCall.Do()
	if err != nil {
		return nil, errors.New("Uploads couldn't be obtained from uploads playlist")
	}

	var videoIDs []string

	for _, item := range uploadsResponse.Items {
		videoIDs = append(videoIDs, item.ContentDetails.VideoId)
	}

	return videoIDs, nil
}
