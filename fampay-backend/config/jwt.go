package config

import (
	"time"
)

const (
	ACCESS_TOKEN_TTL  = 15 * time.Minute
	REFRESH_TOKEN_TTL = 14 * 24 * time.Hour
)
