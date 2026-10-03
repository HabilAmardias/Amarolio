package verifyuserchallenges

import "time"

type (
	VerifyUserChallenge struct {
		ID                           string     `json:"id"`
		UserID                       string     `json:"user_id"`
		CreatedAt                    time.Time  `json:"created_at"`
		UpdatedAt                    time.Time  `json:"updated_at"`
		VerifyUserChallengeExpiredAt time.Time  `json:"verify_user_challenge_expired_at"`
		DeletedAt                    *time.Time `json:"deleted_at"`
	}
)

func (v *VerifyUserChallenge) IsExpired() bool {
	return v.VerifyUserChallengeExpiredAt.Before(time.Now())
}
