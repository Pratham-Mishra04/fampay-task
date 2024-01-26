package helpers

import (
	"context"
	"log"

	"github.com/Pratham-Mishra04/fampay/fampay-backend/initializers"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

var (
	apiKeys      []string
	currentIndex int
)

var Service *youtube.Service

func InitializeService() {
	apiKeys = initializers.CONFIG.YOUTUBE_API_KEYS
	currentIndex = 0

	youtubeService, err := youtube.NewService(context.Background(), option.WithAPIKey(apiKeys[0]))
	if err != nil {
		log.Fatal("Error occurred while connecting to the service. ", err)
	}

	Service = youtubeService
	log.Println("Connected to the service!")
}

func rotateAPIKey() {
	if currentIndex == len(apiKeys)-1 {
		log.Fatal("All KEYS have been exhausted!")
	}

	currentIndex = (currentIndex + 1) % len(apiKeys)
	log.Printf("\nSwitched to API key %d.", currentIndex+1)
	log.Println(len(apiKeys))
}

func UpdateService() {
	rotateAPIKey()

	apiKey := apiKeys[currentIndex]

	youtubeService, err := youtube.NewService(context.Background(), option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatal("Error occurred while connecting to the service. ", err)
	}

	Service = youtubeService
}
