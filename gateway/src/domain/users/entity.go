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
	ResendVerificationBody struct {
		Email string `json:"email"`
	}
)
