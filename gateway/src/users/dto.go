package users

type (
	LoginData struct {
		URL   string `json:"url"`
		State string `json:"oauthstate"`
	}
	AuthData struct {
		Token string `json:"auth_token"`
	}
	RefreshAuthRes struct {
		Message string `json:"message"`
	}
)
