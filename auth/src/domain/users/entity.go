package users

import (
	"time"
)

type (
	User struct {
		ID                          string     `json:"id"`
		Email                       string     `json:"email"`
		Password                    string     `json:"password"`
		Verified                    bool       `json:"verified"`
		VerificationToken           *string    `json:"verification_token"`
		ResetPasswordToken          *string    `json:"reset_password_token"`
		CreatedAt                   time.Time  `json:"created_at"`
		UpdatedAt                   time.Time  `json:"updated_at"`
		VerificationTokenExpiredAt  *time.Time `json:"verification_token_expired_at"`
		ResetPasswordTokenExpiredAt *time.Time `json:"reset_password_token_expired_at"`
		DeletedAt                   *time.Time `json:"deleted_at"`
	}
	SendEmailParams struct {
		Receiver  string
		Subject   string
		EmailBody string
	}
)

func (u *User) IsResetPasswordTokenMatched(token string) bool {
	if u.ResetPasswordToken == nil {
		return false
	}
	return *u.ResetPasswordToken == token
}

func (u *User) IsResetPasswordTokenExpired() bool {
	if u.ResetPasswordTokenExpiredAt == nil {
		return true
	}
	return u.ResetPasswordTokenExpiredAt.Before(time.Now())
}

func (u *User) IsVerificationTokenExpired() bool {
	if u.VerificationTokenExpiredAt == nil {
		return true
	}
	return u.VerificationTokenExpiredAt.Before(time.Now())
}

func (u *User) IsVerificationTokenMatches(token string) bool {
	if u.VerificationToken == nil {
		return false
	}
	return *u.VerificationToken == token
}
