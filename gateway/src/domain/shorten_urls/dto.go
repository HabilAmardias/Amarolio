package shortenurls

import "time"

type (
	NewShortenURLReq struct {
		URL      string `json:"url"`
		Duration *int   `json:"duration"`
	}
	NewShortenURLRes struct {
		URL         string     `json:"url"`
		OriginalURL string     `json:"original_url"`
		ExpiredAt   *time.Time `json:"expired_at"`
	}
)
