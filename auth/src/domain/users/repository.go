package users

import (
	"amarolio-auth/src/constants"
	"amarolio-auth/src/customerrors"
	"amarolio-auth/src/db"
	"context"
	"database/sql"
	"errors"
	"time"
)

type UserRepositoryImpl struct {
	dbtx db.DBTXItf
}

func NewUserRepository(dbtx db.DBTXItf) *UserRepositoryImpl {
	return &UserRepositoryImpl{dbtx}
}

func (ur *UserRepositoryImpl) UpdateUserVerificationStatus(ctx context.Context, userID string, verified bool, user *User) error {
	query := `
	UPDATE users
	SET verified = $1, updated_at = NOW()
	WHERE id = $2 AND deleted_at IS NULL
	RETURNING id, email, password, otp, verified, verification_token, created_at, updated_at, otp_expired_at, verification_token_expired_at, deleted_at
	`
	if err := ur.dbtx.QueryRowContext(ctx, query, verified, userID).Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.OTP,
		&user.Verified,
		&user.VerificationToken,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.OTPExpiredAt,
		&user.VerificationTokenExpiredAt,
		&user.DeletedAt,
	); err != nil {
		return customerrors.NewError(
			"something went wrong",
			err,
			customerrors.DatabaseExecutionErr,
		)
	}
	return nil
}

func (ur *UserRepositoryImpl) UpdateVerificationToken(ctx context.Context, userID string, token string, eat time.Time, user *User) error {
	query := `
	UPDATE users
	SET verification_token = $1, verification_token_expired_at = $2, updated_at = NOW()
	WHERE id = $3 AND deleted_at IS NULL
	RETURNING id, email, password, otp, verified, verification_token, created_at, updated_at, otp_expired_at, verification_token_expired_at, deleted_at
	`

	if err := ur.dbtx.QueryRowContext(ctx, query, token, eat, userID).Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.OTP,
		&user.Verified,
		&user.VerificationToken,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.OTPExpiredAt,
		&user.VerificationTokenExpiredAt,
		&user.DeletedAt,
	); err != nil {
		return customerrors.NewError(
			"something went wrong",
			err,
			customerrors.DatabaseExecutionErr,
		)
	}
	return nil
}

func (ur *UserRepositoryImpl) UpdateOTP(ctx context.Context, userID string, otp *string, user *User) error {
	query := `
	UPDATE users
	SET otp = $1, otp_expired_at = $2, updated_at = NOW()
	WHERE id = $3 AND deleted_at IS NULL
	RETURNING id, email, password, otp, verified, verification_token, created_at, updated_at, otp_expired_at, verification_token_expired_at, deleted_at
	`

	if err := ur.dbtx.QueryRowContext(ctx, query, otp, time.Now().Add(constants.OTP_AGE), userID).Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.OTP,
		&user.Verified,
		&user.VerificationToken,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.OTPExpiredAt,
		&user.VerificationTokenExpiredAt,
		&user.DeletedAt,
	); err != nil {
		return customerrors.NewError(
			"something went wrong",
			err,
			customerrors.DatabaseExecutionErr,
		)
	}
	return nil
}

func (ur *UserRepositoryImpl) FindByID(ctx context.Context, userID string, user *User) error {
	query := `
	SELECT
		id,
		email,
		password,
		otp,
		verified,
		verification_token,
		created_at,
		updated_at,
		otp_expired_at,
		verification_token_expired_at,
		deleted_at
	FROM users
	WHERE id = $1 AND deleted_at IS NULL
	`
	if err := ur.dbtx.QueryRowContext(ctx, query, userID).Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.OTP,
		&user.Verified,
		&user.VerificationToken,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.OTPExpiredAt,
		&user.VerificationTokenExpiredAt,
		&user.DeletedAt,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return customerrors.NewError(
				"user not found",
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

func (ur *UserRepositoryImpl) FindByEmail(ctx context.Context, email string, user *User) error {
	query := `
	SELECT
		id,
		email,
		password,
		otp,
		verified,
		verification_token,
		created_at,
		updated_at,
		otp_expired_at,
		verification_token_expired_at,
		deleted_at
	FROM users
	WHERE email = $1 AND deleted_at IS NULL
	`
	if err := ur.dbtx.QueryRowContext(ctx, query, email).Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.OTP,
		&user.Verified,
		&user.VerificationToken,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.OTPExpiredAt,
		&user.VerificationTokenExpiredAt,
		&user.DeletedAt,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return customerrors.NewError(
				"user not found",
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

func (ur *UserRepositoryImpl) AddNewUser(ctx context.Context, email, hashedPassword string, user *User) error {
	query := `
	INSERT INTO users (email, password)
	VALUES
	($1, $2)
	RETURNING id, email, password, otp, verified, verification_token, created_at, updated_at, otp_expired_at, verification_token_expired_at, deleted_at
	`
	if err := ur.dbtx.QueryRowContext(
		ctx,
		query,
		email,
		hashedPassword,
	).Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.OTP,
		&user.Verified,
		&user.VerificationToken,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.OTPExpiredAt,
		&user.VerificationTokenExpiredAt,
		&user.DeletedAt,
	); err != nil {
		return customerrors.NewError(
			"something went wrong",
			err,
			customerrors.DatabaseExecutionErr,
		)
	}
	return nil
}
