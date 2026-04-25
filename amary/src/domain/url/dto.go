package url

import (
	"amary/src/dto"
	"time"
)

type (
	NewShortURLReq struct {
		URL      string `json:"url" binding:"required,url"`
		Duration *int   `json:"duration" binding:"omitempty,gt=0"`
	}
	NewShortURLRes struct {
		URL string `json:"url"`
	}
	FindLongUrlRes struct {
		URL string `json:"url"`
	}
	GetUserLinksReq struct {
		dto.PaginateOffsetReq
	}
	UserLinkRes struct {
		ID        int64      `json:"id"`
		UserID    *string    `json:"user_id"`
		ShortURL  string     `json:"short_url"`
		LongURL   string     `json:"url"`
		CreatedAt time.Time  `json:"created_at"`
		ExpiredAt *time.Time `json:"expired_at"`
	}
)
