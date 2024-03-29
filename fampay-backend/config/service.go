package config

import (
	"github.com/Pratham-Mishra04/fampay/fampay-backend/initializers"
	"google.golang.org/api/youtube/v3"
)

var (
	OAuthScopes      = []string{youtube.YoutubeForceSslScope}
	OAuthRedirectURI = initializers.CONFIG.BACKEND_URL + "/oauth2callback"
	ServiceQuery     = "football"
	RepeaterDelay    = 10 //in seconds
)
