package users

type (
	LoginData struct {
		URL   string `json:"url"`
		State string `json:"oauthstate"`
	}
)
