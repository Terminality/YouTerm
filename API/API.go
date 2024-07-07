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

func GetKeyFromEnv() (string, error) {
	api := os.Getenv("YOUTERM_API_KEY")
	if api != "" {
		return api, nil
	}
	// TODO: Prompt user for API key, then save it to env variable (os.Setenv())

	return "", errors.New("Couldn't load YOUTERM_API_KEY from environment variable.")
}

// func getClientFromAPI(ctx context.Context, key string) *http.Client {
//
// }

func getServiceFromAPI(ctx context.Context, APIkey string) *youtube.Service {
	service, err := youtube.NewService(ctx, option.WithAPIKey(APIkey))
	if err != nil {
		log.Fatalf("Error creating new YouTube service: %v", err)
	}

	return service
}

func MainAPI() {
	ctx := context.Background()
	APIKey, err := GetKeyFromEnv()
	if err != nil {
		log.Fatalf("Error getting API Key: %v", err)
	}

	service := getServiceFromAPI(ctx, APIKey)

	listInfoByChannelUsername(service, []string{"snippet", "contentDetails", "statistics"}, "Northernlion")
}

func listInfoByChannelUsername(service *youtube.Service, thingsToLoad []string, username string) {
	call := service.Channels.List(thingsToLoad)
	call = call.ForUsername(username)
	resp, err := call.Do()
	if err != nil {
		log.Fatalf("Error making API call: %v", err)
	}

	fmt.Println(fmt.Sprintf("This channel's ID is %s. Its title is '%s', "+
		"and it has %d views.",
		resp.Items[0].Id,
		resp.Items[0].Snippet.Title,
		resp.Items[0].Statistics.ViewCount))
}
