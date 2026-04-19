package shortenurl

import "time"

type (
	ShortenURL struct {
		ID               int64
		UserID           *string
		EncryptedLongUrl string
		CreatedAt        time.Time
		UpdatedAt        time.Time
		DeletedAt        *time.Time
		ExpiredAt        *time.Time
	}
)
