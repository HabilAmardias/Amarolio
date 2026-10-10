package otpchallenges

import "time"

type (
	OTPChallenge struct {
		ID           string     `json:"id"`
		UserID       string     `json:"user_id"`
		OTPHash      string     `json:"otp_hash"`
		CreatedAt    time.Time  `json:"created_at"`
		UpdatedAt    time.Time  `json:"updated_at"`
		OTPExpiredAt time.Time  `json:"otp_expired_at"`
		DeletedAt    *time.Time `json:"deleted_at"`
	}
)

func (o *OTPChallenge) IsExpired() bool {
	return o.OTPExpiredAt.Before(time.Now())
}

func (o *OTPChallenge) VerifyOTP(otp string, hasher HasherItf) (bool, error) {
	return hasher.Validate(o.OTPHash, otp)
}
