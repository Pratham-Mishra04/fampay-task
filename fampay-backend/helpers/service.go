package helpers

import (
	"context"
	"log"

	"github.com/Pratham-Mishra04/fampay/fampay-backend/config"
	"github.com/Pratham-Mishra04/fampay/fampay-backend/initializers"
	"github.com/Pratham-Mishra04/fampay/fampay-backend/utils"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

var Service *youtube.Service
var OAuthConfig *oauth2.Config

func InitializeService() {
	OAuthConfig = &oauth2.Config{
		ClientID:     initializers.CONFIG.OAUTH_CLIENT_ID,
		ClientSecret: initializers.CONFIG.OAUTH_CLIENT_SECRET,
		RedirectURL:  config.OAuthRedirectURI,
		Scopes:       config.OAuthScopes,
		Endpoint:     google.Endpoint,
	}

	// Retrieve the user's access token from storage
	token, err := utils.LoadToken()
	if err != nil {
		log.Fatal("Error occurred while fetching the token. ", err)
	}

	// Create YouTube API service using NewService
	youtubeService, err := youtube.NewService(context.Background(), option.WithTokenSource(OAuthConfig.TokenSource(context.Background(), token)))
	if err != nil {
		log.Fatal("Error occurred while connecting to the service. ", err)
	}

	Service = youtubeService
}
