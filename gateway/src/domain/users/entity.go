package users

type (
	Login struct {
		AuthToken    string `json:"auth_token"`
		RefreshToken string `json:"refresh_token"`
	}
	LoginBody struct {
		OTP string `json:"otp"`
	}
	RefreshAuth struct {
		Token string `json:"auth_token"`
	}
	GetProfile struct {
		Username string `json:"username"`
	}
	Text struct {
		Message string `json:"message"`
	}
	CredentialsBody struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	OTP struct {
		OTPToken string `json:"otp_token"`
	}
	SendEmailBody struct {
		Email string `json:"email"`
	}
	VerificationBody struct {
		UserID string `json:"user_id"`
		Token  string `json:"token"`
	}
	ResetPasswordBody struct {
		UserID      string `json:"user_id"`
		Token       string `json:"token"`
		NewPassword string `json:"new_password"`
	}
)
