package shortenurls

import "time"

type (
	IsCustomURLAvailableBody struct {
		CustomCode string `json:"custom_code"`
	}
	NewShortenURLBody struct {
		URL        string  `json:"url"`
		Duration   *int    `json:"duration"`
		CustomCode *string `json:"custom_code"`
	}
	NewShortenURL struct {
		URL         string     `json:"url"`
		OriginalURL string     `json:"original_url"`
		ExpiredAt   *time.Time `json:"expired_at"`
	}
	OriginalURL struct {
		URL string `json:"url"`
	}
	URL struct {
		ID        int64      `json:"id"`
		UserID    *string    `json:"user_id"`
		ShortURL  string     `json:"short_url"`
		LongURL   string     `json:"url"`
		CreatedAt time.Time  `json:"created_at"`
		ExpiredAt *time.Time `json:"expired_at"`
	}
	FindOriginalURL struct {
		URL URL `json:"url"`
	}
)
