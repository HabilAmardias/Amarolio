package url

import "time"

type (
	URL struct {
		ID        int64
		UserID    *string
		LongURL   string
		ShortCode *string
		CreatedAt time.Time
		UpdatedAt time.Time
		DeletedAt *time.Time
		ExpiredAt *time.Time
	}
	ParsedURL struct {
		ID        int64
		UserID    *string
		ShortURL  string
		Code      string
		LongURL   string
		CreatedAt time.Time
		ExpiredAt *time.Time
	}
)
