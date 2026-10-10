package users

type (
	LoginReq struct {
		OTP string `json:"otp" validate:"required,numeric,min=6,max=6"`
	}
	LoginRes struct {
		AuthToken    string `json:"auth_token"`
		RefreshToken string `json:"refresh_token"`
	}
	RefreshAuthRes struct {
		AuthToken string `json:"auth_token"`
	}
	GetProfileRes struct {
		Username string `json:"username"`
	}
	VerifyReq struct {
		Token string `json:"token" validate:"required"`
	}
	CredentialsReq struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=8,max=13,alphanum"`
	}
	OTPRes struct {
		OTPToken string `json:"otp_token"`
	}
	SendEmailReq struct {
		Email string `json:"email" validate:"required,email"`
	}
	ResetPasswordReq struct {
		Token       string `json:"token" validate:"required"`
		NewPassword string `json:"new_password" validate:"required,min=8,max=13,alphanum"`
	}
)
