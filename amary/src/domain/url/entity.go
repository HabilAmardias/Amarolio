package url

import "time"

type (
	URL struct {
		ID               int64
		UserID           *string
		EncryptedLongUrl string
		CreatedAt        time.Time
		UpdatedAt        time.Time
		DeletedAt        *time.Time
		ExpiredAt        *time.Time
	}
	DecryptedURL struct {
		ID        int64
		UserID    *string
		ShortURL  string
		LongURL   string
		CreatedAt time.Time
		ExpiredAt *time.Time
	}
)
