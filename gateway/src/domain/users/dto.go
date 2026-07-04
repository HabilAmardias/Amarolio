package users

type (
	RefreshAuthRes struct {
		Message string `json:"message"`
	}
	GetProfileRes struct {
		Username string `json:"username"`
	}
	LoginReq struct {
		RedirectURI string `json:"redirect_uri" validate:"required"`
	}
	LogoutReq struct {
		RedirectURI string `json:"redirect_uri" validate:"required"`
	}
	LoginRes struct {
		RedirectURI string `json:"redirect_uri"`
	}
	LogoutRes struct {
		RedirectURI string `json:"redirect_uri"`
	}
)
