package users

type (
	RefreshAuthRes struct {
		Message string `json:"message"`
	}
	GetProfileRes struct {
		Email string `json:"email"`
	}
)
