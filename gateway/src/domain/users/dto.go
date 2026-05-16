package users

type (
	RefreshAuthRes struct {
		Message string `json:"message"`
	}
	GetProfileRes struct {
		Email string `json:"email"`
	}
	LoginReq struct {
		RedirectURI string `query:"redirect_uri" validate:"required"`
	}
	LogoutReq struct {
		RedirectURI string `query:"redirect_uri" validate:"required"`
	}
)
