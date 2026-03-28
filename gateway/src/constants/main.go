package constants

import "time"

const (
	PRODUCTION        = "PRODUCTION"
	REFRESH_CLAIM_KEY = "refresh_claim_key"
	AUTH_TOKEN        = "auth_token"
	REFRESH_TOKEN     = "refresh_token"
	AUTH_AGE          = 15 * time.Minute
	REFRESH_AGE       = 7 * 24 * time.Hour
)

const (
	ForAuth = iota + 1
	ForRefresh
)
