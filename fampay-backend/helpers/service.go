package helpers

import (
	"context"
	"log"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

var YoutubeService *youtube.Service

func InitializeYoutubeService() {
	apiKey := os.Getenv("YOUTUBE_API_KEY")
	if apiKey == "" {
		log.Fatal("YOUTUBE_API_KEY environment variable not set")
	}

	// Create a new HTTP client using the provided API key
	client := &http.Client{
		Transport: &oauth2.Transport{
			Base:   http.DefaultTransport,
			Source: oauth2.StaticTokenSource(&oauth2.Token{AccessToken: apiKey}),
		},
	}

	// Create a new YouTube client with the obtained HTTP client
	service, err := youtube.NewService(context.Background(), option.WithHTTPClient(client))
	if err != nil {
		log.Fatal("Error occurred while connecting to client. ", err)
	}

	YoutubeService = service
}
