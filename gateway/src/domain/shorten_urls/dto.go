package shortenurls

type (
	NewShortenURLReq struct {
		URL      string `json:"url"`
		Duration *int   `json:"duration"`
	}
	NewShortenURLRes struct {
		URL string `json:"url"`
	}
)
