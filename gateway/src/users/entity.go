package users

type (
	LoginCallbackData struct {
		AuthToken    string `json:"auth_token"`
		RefreshToken string `json:"refresh_token"`
	}
	AuthRefreshData struct {
		AuthToken    string `json:"auth_token"`
		RefreshToken string `json:"refresh_token"`
	}
	LoginData struct {
		URL   string `json:"url"`
		State string `json:"oauthstate"`
	}
	AuthData struct {
		Token string `json:"auth_token"`
	}
)
