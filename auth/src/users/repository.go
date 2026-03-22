package users

import (
	"amarolio-auth/src/customerrors"
	"context"
	"database/sql"
	"errors"
)

type DBTXItf interface {
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
}

type UserRepositoryImpl struct {
	dbtx DBTXItf
}

func NewUserRepository(dbtx DBTXItf) *UserRepositoryImpl {
	return &UserRepositoryImpl{dbtx}
}

func (ur *UserRepositoryImpl) FindByID(ctx context.Context, userID string, user *User) error {
	query := `
	SELECT
		id,
		email,
		created_at,
		updated_at,
		deleted_at
	FROM users
	WHERE id = $1 AND deleted_at IS NULL
	`
	if err := ur.dbtx.QueryRowContext(ctx, query, userID).Scan(
		&user.ID,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
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
		created_at,
		updated_at,
		deleted_at
	FROM users
	WHERE email = $1 AND deleted_at IS NULL
	`
	if err := ur.dbtx.QueryRowContext(ctx, query, email).Scan(
		&user.ID,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
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

func (ur *UserRepositoryImpl) AddNewUser(ctx context.Context, email string, user *User) error {
	query := `
	INSERT INTO users (email)
	VALUES
	($1)
	RETURNING id, email, created_at, updated_at, deleted_at
	`
	if err := ur.dbtx.QueryRowContext(ctx, query, email).Scan(
		&user.ID,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
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
