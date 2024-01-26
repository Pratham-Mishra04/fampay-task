package helpers

import (
	"context"
	"log"

	"github.com/Pratham-Mishra04/fampay/fampay-backend/initializers"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

var Service *youtube.Service

func InitializeService() {
	youtubeService, err := youtube.NewService(context.Background(), option.WithAPIKey(initializers.CONFIG.YOUTUBE_API_KEY))
	if err != nil {
		log.Fatal("Error occurred while connecting to the service. ", err)
	}

	Service = youtubeService
}
