package constants

import "time"

const (
	PRODUCTION        = "PRODUCTION"
	REFRESH_CLAIM_KEY = "refresh_claim_key"
	AUTH_TOKEN        = "auth_token"
	AUTH_AGE          = 15 * time.Minute
)

const (
	ForAuth = iota + 1
	ForRefresh
)
