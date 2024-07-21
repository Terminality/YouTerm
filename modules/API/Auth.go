package API

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"dalton.dog/YouTerm/utils"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

// Have instruction page w/ links to Google Dev Console
// Prompt for API key or OAuth2
// API Key
//	Try to load from env var. If can't, prompt for it
// OAuth2
//	Prompt for path to load JSON from. Store path to user
//		Not sure if it's good/possible/stupid to store the JSON data itself in the user
//	Get config from google.ConfigFromJSON()
//	Try loading token from file. If can't...
//		Launch auth URL, get auth code
//		Exchange for token, save token

func GetAuthenticatedService(ctx context.Context) (*youtube.Service, error) {
	service, oauthErr := getServiceFromOAuth(ctx)
	if service != nil {
		return service, nil
	}

	service, apiErr := getServiceFromAPI(ctx)
	if service != nil {
		return service, nil
	}

	errorMsg := "Couldn't authenticate.\n"
	errorMsg += fmt.Sprintf("OAuth error: %v", oauthErr)
	errorMsg += fmt.Sprintf("API Key error: %v", apiErr)

	return nil, errors.New(errorMsg)
}

func getServiceFromOAuth(ctx context.Context) (*youtube.Service, error) {
	authConfig, err := getConfig(getCredentialsFilepath())
	if err != nil {
		return nil, err
	}
	client := getClientFromConfig(ctx, authConfig)

	service, err := youtube.NewService(ctx, option.WithHTTPClient(client))
	return service, err
}

func getConfig(path string) (*oauth2.Config, error) {
	jsonBytes, err := os.ReadFile(path)
	utils.CheckErrFatal(err, fmt.Sprintf("Couldn't load credentials JSON file from %v", path))

	return google.ConfigFromJSON(jsonBytes, youtube.YoutubeReadonlyScope)
}

func getClientFromConfig(ctx context.Context, config *oauth2.Config) *http.Client {
	token, err := getToken(config)
	if err != nil {
		return nil
	}
	return config.Client(ctx, token)
}

func getToken(config *oauth2.Config) (*oauth2.Token, error) {
	tokenPath := getTokenFilePath()
	tokenFile, err := os.Open(tokenPath)
	if err == nil {
		defer tokenFile.Close()
		token := &oauth2.Token{}
		err = json.NewDecoder(tokenFile).Decode(token)
		return token, err
	}

	// Couldn't load token from file, so getting it from web instead
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	utils.LaunchURL(authURL)

	fmt.Println("Paste your auth code: ")
	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		return nil, errors.New(fmt.Sprintf("Unable to read authorization code: %v", err))
	}

	token, err := config.Exchange(context.Background(), authCode)
	if err != nil {
		return nil, err
	}

	tokenFile, err = os.Create(tokenPath)
	if err != nil {
		return nil, err
	}

	json.NewEncoder(tokenFile).Encode(token)
	tokenFile.Close()

	return token, nil
}

func getCredentialsFilepath() string {
	homedir, err := os.UserHomeDir()
	utils.CheckErrFatal(err, "No user home dir to load")

	return filepath.Join(homedir, ".terminality", "youterm", "credentials.json")
}

func getTokenFilePath() string {
	homedir, err := os.UserHomeDir()
	utils.CheckErrFatal(err, "No user home dir to load")

	return filepath.Join(homedir, ".terminality", "youterm", "token.json")
}

func getKeyFromEnv() (string, error) {
	log.Println("Loading API key from env var")
	api := os.Getenv("YOUTERM_API_KEY")
	if api != "" {
		return api, nil
	}

	return "", errors.New("Couldn't load YOUTERM_API_KEY from environment variable.")
}

// Attempts to create an authenticated YouTube service from an API key
func getServiceFromAPI(ctx context.Context) (*youtube.Service, error) {
	log.Println("Creating service from API key")
	APIkey, err := getKeyFromEnv()
	if err != nil {
		return nil, err
	}

	// TODO: Prompt for API key, store in envVar
	// os.Setenv("YOUTERM_API_KEY", APIkey)

	service, err := youtube.NewService(ctx, option.WithAPIKey(APIkey))
	if err != nil {
		return nil, err
	}

	return service, nil
}
