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

func InitializeManager() {
	ctx := context.Background()
	APIKey, err := getKeyFromEnv()
	if err != nil {
		log.Fatalf("Error getting API Key: %v", err)
	}

	service := getServiceFromAPI(ctx, APIKey)

	masterAPI = &apiManager{
		ctx:     ctx,
		apiKey:  APIKey,
		service: service,
	}
}

// TODO: Remove this
func GetIDFromAPI(username string) string {
	channelCall := masterAPI.service.Channels.List([]string{"contentDetails"})
	if username != "" {
		channelCall.ForUsername(username)
	} else {
		log.Fatalln("No channel identifier passed into request!")
	}
	channelResponse, err := channelCall.Do()
	if err != nil {
		log.Fatalf("Channel content information couldn't be obtained: %v", err)
	}

	return channelResponse.Items[0].Id
}

func getKeyFromEnv() (string, error) {
	api := os.Getenv("YOUTERM_API_KEY")
	if api != "" {
		return api, nil
	}
	// TODO: Prompt user for API key, then save it to env variable (os.Setenv())

	return "", errors.New("Couldn't load YOUTERM_API_KEY from environment variable.")
}

func getServiceFromAPI(ctx context.Context, APIkey string) *youtube.Service {
	service, err := youtube.NewService(ctx, option.WithAPIKey(APIkey))
	if err != nil {
		log.Fatalf("Error creating new YouTube service: %v", err)
	}

	return service
}

// TODO: Reimplement this properly
func GetUploadsForChannel(service *youtube.Service, channelID string, channelUsername string, pageID string) []string {
	channelCall := service.Channels.List([]string{"contentDetails"})
	if channelID != "" {
		channelCall.Id(channelID)
	} else if channelUsername != "" {
		channelCall.ForUsername(channelUsername)
	} else {
		log.Fatalln("No channel identifier passed into request!")
	}
	channelResponse, err := channelCall.Do()
	if err != nil {
		log.Fatalf("Channel content information couldn't be obtained: %v", err)
	}

	// TODO: Implement a check for successive pages using the pageID

	uploadsPlaylist := channelResponse.Items[0].ContentDetails.RelatedPlaylists.Uploads
	uploadsCall := service.PlaylistItems.List([]string{"contentDetails"}).PlaylistId(uploadsPlaylist).MaxResults(20)
	uploadsResponse, err := uploadsCall.Do()
	if err != nil {
		log.Fatalf("Uploads couldn't be obtained from uploads playlist: %v", err)
	}

	var videoIDs []string

	for _, item := range uploadsResponse.Items {
		videoIDs = append(videoIDs, item.ContentDetails.VideoId)
	}

	return videoIDs
}

func RequestVideo(videoID string) (*youtube.VideoListResponse, error) {
	call := masterAPI.service.Videos.List([]string{"snippet", "statistics"})
	call = call.Id(videoID)
	resp, err := call.Do()
	if err != nil {
		return nil, errors.New("Unable to load video")
	}

	return resp, nil
}

func RequestChannel(userID string, username string, handle string) *youtube.ChannelListResponse {
	call := masterAPI.service.Channels.List([]string{"snippet", "contentDetails", "statistics"})
	call = call.ForUsername(username)
	resp, err := call.Do()
	if err != nil {
		log.Fatalf("Error making API call: %v", err)
	}

	return resp
}
