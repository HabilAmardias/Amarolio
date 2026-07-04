package users

type (
	RefreshAuthRes struct {
		Message string `json:"message"`
	}
	GetProfileRes struct {
		Email string `json:"email"`
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
