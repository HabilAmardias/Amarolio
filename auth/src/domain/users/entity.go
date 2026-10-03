package users

import (
	"time"
)

type (
	User struct {
		ID        string     `json:"id"`
		Email     string     `json:"email"`
		Password  string     `json:"password"`
		Verified  bool       `json:"verified"`
		CreatedAt time.Time  `json:"created_at"`
		UpdatedAt time.Time  `json:"updated_at"`
		DeletedAt *time.Time `json:"deleted_at"`
	}
	SendEmailParams struct {
		Receiver  string
		Subject   string
		EmailBody string
	}
)
