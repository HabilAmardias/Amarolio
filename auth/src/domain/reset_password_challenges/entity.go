package resetpasswordchallenges

import "time"

type (
	ResetPasswordChallenge struct {
		ID                              string     `json:"id"`
		UserID                          string     `json:"user_id"`
		CreatedAt                       time.Time  `json:"created_at"`
		UpdatedAt                       time.Time  `json:"updated_at"`
		ResetPasswordChallengeExpiredAt time.Time  `json:"reset_password_challenge_expired_at"`
		DeletedAt                       *time.Time `json:"deleted_at"`
	}
)

func (v *ResetPasswordChallenge) IsExpired() bool {
	return v.ResetPasswordChallengeExpiredAt.Before(time.Now())
}
