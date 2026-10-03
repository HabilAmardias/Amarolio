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
		ResetPasswordToken          *string    `json:"reset_password_token"`
		CreatedAt                   time.Time  `json:"created_at"`
		UpdatedAt                   time.Time  `json:"updated_at"`
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
