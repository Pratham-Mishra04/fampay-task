package helpers

import (
	"context"
	"log"
	"time"

	"github.com/Pratham-Mishra04/fampay/fampay-backend/initializers"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

var (
	apiKeys      []string
	currentIndex int
)

var Service *youtube.Service

func InitializeService() {
	if len(apiKeys) == 0 {
		apiKeys = initializers.CONFIG.YOUTUBE_API_KEYS
		currentIndex = 0
	}

	err := retry(len(apiKeys), 2*time.Second, func() error {
		apiKey := apiKeys[currentIndex]

		youtubeService, err := youtube.NewService(context.Background(), option.WithAPIKey(apiKey))
		if err != nil {
			if apiErr, ok := err.(*googleapi.Error); ok {
				// Check if the error is due to an expired or invalid API key
				if apiErr.Code == 403 && apiErr.Errors[0].Reason == "accessNotConfigured" {
					rotateAPIKey()
				}
			}
			return err
		}

		Service = youtubeService

		log.Println("Connected to the service!")
		return nil
	})

	if err != nil {
		log.Fatal("Error occurred while connecting to the service. ", err)
	}
}

func rotateAPIKey() {
	currentIndex = (currentIndex + 1) % len(apiKeys)
	log.Printf("\nSwitched to API key %d.", currentIndex+1)
}

func retry(maxAttempts int, delay time.Duration, fn func() error) error {
	var err error
	for attempt := 1; attempt <= maxAttempts; attempt++ {
		err = fn()
		if err == nil {
			break
		}
		log.Printf("\nError connecting to the service, retrying attempt %d: %v", attempt, err)
		time.Sleep(delay)
	}
	return err
}
