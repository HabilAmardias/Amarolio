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
)
