package users

import (
	"time"
)

type (
	User struct {
		ID                         string     `json:"id"`
		Email                      string     `json:"email"`
		Password                   string     `json:"password"`
		OTP                        *string    `json:"otp"`
		Verified                   bool       `json:"verified"`
		VerificationToken          *string    `json:"verification_token"`
		CreatedAt                  time.Time  `json:"created_at"`
		UpdatedAt                  time.Time  `json:"updated_at"`
		OTPExpiredAt               *time.Time `json:"otp_updated_at"`
		VerificationTokenExpiredAt *time.Time `json:"verification_token_expired_at"`
		DeletedAt                  *time.Time `json:"deleted_at"`
	}
	SendEmailParams struct {
		Receiver  string
		Subject   string
		EmailBody string
	}
)

func (u *User) IsVerificationTokenExpired() bool {
	if u.VerificationTokenExpiredAt == nil {
		return true
	}
	return u.VerificationTokenExpiredAt.Before(time.Now())
}

func (u *User) IsOTPExpired() bool {
	if u.OTPExpiredAt == nil {
		return true
	}
	return u.OTPExpiredAt.Before(time.Now())
}

func (u *User) IsVerificationTokenMatches(token string) bool {
	if u.VerificationToken == nil {
		return false
	}
	return *u.VerificationToken == token
}

func (u *User) IsOTPMatches(otp string) bool {
	if u.OTP == nil {
		return false
	}
	return *u.OTP == otp
}
