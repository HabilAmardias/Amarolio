package users

type (
	LoginRes struct {
		URL   string `json:"url"`
		State string `json:"oauthstate"`
	}
	LoginCallbackReq struct {
		State string `json:"oauthstate" validate:"required"`
	}
	LoginCallbackRes struct {
		AuthToken    string `json:"auth_token"`
		RefreshToken string `json:"refresh_token"`
	}
	RefreshAuthRes struct {
		AuthToken string `json:"auth_token"`
	}
)
