package shortenurls

import "time"

type (
	NewShortenURLBody struct {
		URL      string `json:"url"`
		Duration *int   `json:"duration"`
	}
	NewShortenURL struct {
		URL         string     `json:"url"`
		OriginalURL string     `json:"original_url"`
		ExpiredAt   *time.Time `json:"expired_at"`
	}
	FindLongURL struct {
		URL string `json:"url"`
	}
	UserLink struct {
		ID        int64      `json:"id"`
		UserID    *string    `json:"user_id"`
		ShortURL  string     `json:"short_url"`
		LongURL   string     `json:"url"`
		CreatedAt time.Time  `json:"created_at"`
		ExpiredAt *time.Time `json:"expired_at"`
	}
)
