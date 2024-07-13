package API

import (
	"context"
	"errors"
	"fmt"
	"log"

	// "net/http"
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
	APIKey, err := GetKeyFromEnv()
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

func GetKeyFromEnv() (string, error) {
	api := os.Getenv("YOUTERM_API_KEY")
	if api != "" {
		// fmt.Println(fmt.Sprintf("Key loaded: %v", api))
		return api, nil
	}
	// TODO: Prompt user for API key, then save it to env variable (os.Setenv())

	return "", errors.New("Couldn't load YOUTERM_API_KEY from environment variable.")
}

// func getClientFromAPI(ctx context.Context, key string) *http.Client {
//
// }

func GetService(ctx context.Context) *youtube.Service {
	APIKey, err := GetKeyFromEnv()
	if err != nil {
		log.Fatalf("Error getting API Key: %v", err)
	}

	service, err := youtube.NewService(ctx, option.WithAPIKey(APIKey))
	if err != nil {
		log.Fatalf("Error creating new YouTube service: %v", err)
	}

	return service

}

// TODO: Get subscription list from user (apparently this is pretty hard, but haven't looked too deep into it)
// TODO: Add videos to user's Watch Later playlist

func getServiceFromAPI(ctx context.Context, APIkey string) *youtube.Service {
	service, err := youtube.NewService(ctx, option.WithAPIKey(APIkey))
	if err != nil {
		log.Fatalf("Error creating new YouTube service: %v", err)
	}

	return service
}

// TODO: Move all of this stuff to either Main or TUI.
// Context and Service shouldn't be created/maintained in here, since it needs to be passed in to the functions
// Alternatively, could make all of these struct methods of a custom Service wrapper. Something to mull over
func MainAPI() {
	ctx := context.Background()
	APIKey, err := GetKeyFromEnv()
	if err != nil {
		log.Fatalf("Error getting API Key: %v", err)
	}

	service := getServiceFromAPI(ctx, APIKey)

	videoIDs := GetUploadsForChannel(service, "", "Northernlion", "")

	fmt.Sprintln(videoIDs)

	for _, video := range videoIDs {
		PrintInfoForVideo(service, video)
	}

	// listInfoByChannelUsername(service, []string{"snippet", "contentDetails", "statistics"}, "Northernlion")
	// listInfoByChannelUsername(service, []string{"snippet", "contentDetails", "statistics"}, "CobaltStreak")
	// listInfoByChannelUsername(service, []string{"snippet", "contentDetails", "statistics"}, "FakeGoogleUsername")
}

func PrintInfoForVideo(service *youtube.Service, videoID string) {
	call := service.Videos.List([]string{"contentDetails", "snippet"}).Id(videoID)
	response, err := call.Do()
	if err != nil {
		log.Fatalf("Video information couldn't be obtained: %v", err)
	}

	for _, item := range response.Items {
		fmt.Print(fmt.Sprintf("%v -- %v in length -- Uploaded at %v\n", item.Snippet.Title, item.ContentDetails.Duration, item.Snippet.PublishedAt))
	}
}

func GetUploadsForChannel(service *youtube.Service, channelID string, channelUsername string, pageID string) []string {
	// fmt.Sprintf("Trying to fetch uploads for %v\n", channelUsername)
	// TODO: Check to see if we've already loaded this channel before, and have its uploadID stored
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

	// fmt.Sprintln(channelResponse)

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

func RequestChannelFromAPI(userID string, username string, handle string) *youtube.ChannelListResponse {
	call := masterAPI.service.Channels.List([]string{"snippet", "contentDetails", "statistics"})
	call = call.ForUsername(username)
	resp, err := call.Do()
	if err != nil {
		log.Fatalf("Error making API call: %v", err)
	}

	return resp
}

func GetInfoByUsername(username string) *youtube.ChannelListResponse {
	call := masterAPI.service.Channels.List([]string{"snippet", "contentDetails", "statistics"})
	call = call.ForUsername(username)
	resp, err := call.Do()
	if err != nil {
		log.Fatalf("Error making API call: %v", err)
	}

	return resp
}

func listInfoByChannelUsername(service *youtube.Service, thingsToLoad []string, username string) {
	call := service.Channels.List(thingsToLoad)
	call = call.ForUsername(username)
	resp, err := call.Do()
	if err != nil {
		log.Fatalf("Error making API call: %v", err)
	}

	if len(resp.Items) == 0 {
		fmt.Println("No channel found with username ", username)
		return
	}

	fmt.Println(fmt.Sprintf("This channel's ID is %s. Its title is '%s', "+
		"and it has %d views.",
		resp.Items[0].Id,
		resp.Items[0].Snippet.Title,
		resp.Items[0].Statistics.ViewCount))
}
