package shortenurls

type (
	NewShortenURLBody struct {
		URL      string `json:"url"`
		Duration *int   `json:"duration"`
	}
	NewShortenURL struct {
		URL string `json:"url"`
	}
	FindLongURL struct {
		URL string `json:"url"`
	}
)
