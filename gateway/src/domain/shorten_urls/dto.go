package shortenurls

import (
	"amarolio-gateway/src/dto"
	"time"
)

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
	GetUserLinksReq struct {
		dto.SeekPaginateReq
	}
	GetUserLinkRes struct {
		ID        int64      `json:"id"`
		UserID    *string    `json:"user_id"`
		ShortURL  string     `json:"short_url"`
		LongURL   string     `json:"url"`
		CreatedAt time.Time  `json:"created_at"`
		ExpiredAt *time.Time `json:"expired_at"`
	}
)
