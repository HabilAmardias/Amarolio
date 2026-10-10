package otpchallenges

import (
	"amarolio-auth/src/customerrors"
	"amarolio-auth/src/db"
	"context"
	"database/sql"
	"errors"
	"time"
)

type OTPChallengeRepositoryImpl struct {
	dbtx db.DBTXItf
}

func NewOTPChallengeRepository(dbtx db.DBTXItf) *OTPChallengeRepositoryImpl {
	return &OTPChallengeRepositoryImpl{dbtx}
}

func (r *OTPChallengeRepositoryImpl) DeleteAllByUserID(ctx context.Context, userID string) error {
	query := `
	UPDATE otp_challenges
	SET deleted_at = NOW()
	WHERE user_id = $1 AND deleted_at IS NULL
	`
	if _, err := r.dbtx.ExecContext(
		ctx,
		query,
		userID,
	); err != nil {
		return customerrors.NewError(
			"something went wrong",
			err,
			customerrors.DatabaseExecutionErr,
		)
	}
	return nil
}

func (r *OTPChallengeRepositoryImpl) FindByID(ctx context.Context, challengeID string, challenge *OTPChallenge) error {
	query := `
	SELECT id, user_id, otp_hash, created_at, updated_at, otp_expired_at, deleted_at
	FROM otp_challenges
	WHERE id = $1 AND deleted_at IS NULL
	`
	if err := r.dbtx.QueryRowContext(
		ctx,
		query,
		challengeID,
	).Scan(
		&challenge.ID,
		&challenge.UserID,
		&challenge.OTPHash,
		&challenge.CreatedAt,
		&challenge.UpdatedAt,
		&challenge.OTPExpiredAt,
		&challenge.DeletedAt,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return customerrors.NewError(
				"Please login first",
				err,
				customerrors.ItemNotFound,
			)
		}
		return customerrors.NewError(
			"something went wrong",
			err,
			customerrors.DatabaseExecutionErr,
		)
	}
	return nil
}

func (r *OTPChallengeRepositoryImpl) CreateNewOTPChallenge(ctx context.Context, challengeID string, userID string, otpHash string, eat time.Time, challenge *OTPChallenge) error {
	query := `
	INSERT INTO otp_challenges (id, user_id, otp_hash, otp_expired_at)
	VALUES
	($1, $2, $3, $4)
	RETURNING id, user_id, otp_hash, created_at, updated_at, otp_expired_at, deleted_at
	`
	if err := r.dbtx.QueryRowContext(
		ctx,
		query,
		challengeID,
		userID,
		otpHash,
		eat,
	).Scan(
		&challenge.ID,
		&challenge.UserID,
		&challenge.OTPHash,
		&challenge.CreatedAt,
		&challenge.UpdatedAt,
		&challenge.OTPExpiredAt,
		&challenge.DeletedAt,
	); err != nil {
		return customerrors.NewError(
			"something went wrong",
			err,
			customerrors.DatabaseExecutionErr,
		)
	}
	return nil
}
