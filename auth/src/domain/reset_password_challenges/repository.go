package resetpasswordchallenges

import (
	"amarolio-auth/src/customerrors"
	"amarolio-auth/src/db"
	"context"
	"database/sql"
	"errors"
	"time"
)

type ResetPasswordChallengeRepositoryImpl struct {
	dbtx db.DBTXItf
}

func NewResetPasswordChallengeRepository(dbtx db.DBTXItf) *ResetPasswordChallengeRepositoryImpl {
	return &ResetPasswordChallengeRepositoryImpl{dbtx}
}

func (r *ResetPasswordChallengeRepositoryImpl) DeleteAllByUserID(ctx context.Context, userID string) error {
	query := `
	UPDATE reset_password_challenges
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

func (r *ResetPasswordChallengeRepositoryImpl) FindByID(ctx context.Context, challengeID string, challenge *ResetPasswordChallenge) error {
	query := `
	SELECT id, user_id, created_at, updated_at, reset_password_challenge_expired_at, deleted_at
	FROM reset_password_challenges
	WHERE id = $1 AND deleted_at IS NULL
	`
	if err := r.dbtx.QueryRowContext(
		ctx,
		query,
		challengeID,
	).Scan(
		&challenge.ID,
		&challenge.UserID,
		&challenge.CreatedAt,
		&challenge.UpdatedAt,
		&challenge.ResetPasswordChallengeExpiredAt,
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

func (r *ResetPasswordChallengeRepositoryImpl) CreateNewResetPasswordChallenge(ctx context.Context, challengeID string, userID string, eat time.Time, challenge *ResetPasswordChallenge) error {
	query := `
	INSERT INTO reset_password_challenges (id, user_id, reset_password_challenge_expired_at)
	VALUES
	($1, $2, $3)
	RETURNING id, user_id, created_at, updated_at, reset_password_challenge_expired_at, deleted_at
	`
	if err := r.dbtx.QueryRowContext(
		ctx,
		query,
		challengeID,
		userID,
		eat,
	).Scan(
		&challenge.ID,
		&challenge.UserID,
		&challenge.CreatedAt,
		&challenge.UpdatedAt,
		&challenge.ResetPasswordChallengeExpiredAt,
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
