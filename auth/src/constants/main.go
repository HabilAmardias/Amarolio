package constants

import "time"

const (
	AUTH_AGE    = 15 * time.Minute
	REFRESH_AGE = 7 * 24 * time.Hour
)

const (
	ForAuth = iota + 1
	ForRefresh
)

const (
	PRODUCTION = "PRODUCTION"
)

const (
	AUTH_KEY = "auth_key"
)
