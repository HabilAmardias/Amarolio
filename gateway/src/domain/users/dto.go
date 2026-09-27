package users

type (
	GetProfileRes struct {
		Username string `json:"username"`
	}
	LoginReq struct {
		OTP string `json:"otp"`
	}
	VerifyReq struct {
		Token  string `json:"token"`
		UserID string `json:"user_id"`
	}
	CredentialsReq struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	OTPRes struct {
		OTPToken string `json:"otp_token"`
	}
	ResendVerificationReq struct {
		Email string `json:"email"`
	}
)
