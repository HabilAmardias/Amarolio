package users

type (
	LoginCallback struct {
		AuthToken    string `json:"auth_token"`
		RefreshToken string `json:"refresh_token"`
	}
	LoginCallbackBody struct {
		State string `json:"oauthstate"`
	}
	Login struct {
		URL   string `json:"url"`
		State string `json:"oauthstate"`
	}
	RefreshAuth struct {
		Token string `json:"auth_token"`
	}
)
