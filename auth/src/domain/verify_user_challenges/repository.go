package verifyuserchallenges

import (
	"amarolio-auth/src/customerrors"
	"amarolio-auth/src/db"
	"context"
	"database/sql"
	"errors"
	"time"
)

type VerifyUserChallengeRepositoryImpl struct {
	dbtx db.DBTXItf
}

func NewVerifyUserChallengeRepository(dbtx db.DBTXItf) *VerifyUserChallengeRepositoryImpl {
	return &VerifyUserChallengeRepositoryImpl{dbtx}
}

func (r *VerifyUserChallengeRepositoryImpl) DeleteAllByUserID(ctx context.Context, userID string) error {
	query := `
	UPDATE verify_user_challenges
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

func (r *VerifyUserChallengeRepositoryImpl) FindByID(ctx context.Context, challengeID string, challenge *VerifyUserChallenge) error {
	query := `
	SELECT id, user_id, created_at, updated_at, verify_user_challenge_expired_at, deleted_at
	FROM verify_user_challenges
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
		&challenge.VerifyUserChallengeExpiredAt,
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

func (r *VerifyUserChallengeRepositoryImpl) CreateNewVerifyUserChallenge(ctx context.Context, challengeID string, userID string, eat time.Time, challenge *VerifyUserChallenge) error {
	query := `
	INSERT INTO verify_user_challenges (id, user_id, verify_user_challenge_expired_at)
	VALUES
	($1, $2, $3)
	RETURNING id, user_id, created_at, updated_at, verify_user_challenge_expired_at, deleted_at
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
		&challenge.VerifyUserChallengeExpiredAt,
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
