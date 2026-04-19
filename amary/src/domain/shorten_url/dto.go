package shortenurl

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
)
