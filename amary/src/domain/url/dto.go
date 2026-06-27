package url

import (
	"amary/src/dto"
	"time"
)

type (
	FindCustomURLReq struct {
		CustomCode string `json:"custom_code" binding:"required"`
	}
	NewShortURLReq struct {
		URL        string  `json:"url" binding:"required,url"`
		Duration   *int    `json:"duration" binding:"omitempty,gt=0"`
		CustomCode *string `json:"custom_code" binding:"omitempty,min=1"`
	}
	NewShortURLRes struct {
		URL         string     `json:"url"`
		OriginalURL string     `json:"original_url"`
		ExpiredAt   *time.Time `json:"expired_at"`
	}
	VisitOriginalUrlRes struct {
		URL string `json:"url"`
	}
	GetUserLinksReq struct {
		dto.SeekPaginateReq
	}
	FindOriginalUrlRes struct {
		URL URLRes `json:"url"`
	}
	URLRes struct {
		ID        int64      `json:"id"`
		UserID    *string    `json:"user_id"`
		ShortURL  string     `json:"short_url"`
		Code      string     `json:"code"`
		LongURL   string     `json:"url"`
		CreatedAt time.Time  `json:"created_at"`
		ExpiredAt *time.Time `json:"expired_at"`
	}
)
