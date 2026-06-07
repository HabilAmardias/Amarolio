package url

import (
	"amary/src/customerror"
	"amary/src/db"
	"context"
	"database/sql"
	"errors"
	"time"
)

type URLRepoImpl struct {
	handle db.DBTX
}

func NewURLRepo(handle db.DBTX) *URLRepoImpl {
	return &URLRepoImpl{handle}
}

func (sur *URLRepoImpl) GetUserLinkCount(ctx context.Context, userID string, count *int64) error {
	query := `
	SELECT
		COUNT(id)
	FROM urls
	WHERE user_id = $1 AND deleted_at IS NULL
	`
	if err := sur.handle.QueryRowContext(ctx, query, userID).Scan(count); err != nil {
		return customerror.NewError(
			"something went wrong",
			err,
			customerror.DatabaseExecutionErr,
		)
	}
	return nil
}

func (sur *URLRepoImpl) FindUserLinks(ctx context.Context, userID string, lastID *int64, limit int64, links *[]URL) error {
	query := `
	SELECT
		id,
		user_id,
		encrypted_long_url,
		created_at,
		updated_at,
		deleted_at,
		expired_at
	FROM urls
	WHERE user_id = $1 
	AND id < COALESCE($2, 9223372036854775807)
	AND deleted_at IS NULL
	ORDER BY id DESC
	LIMIT $3
	`
	rows, err := sur.handle.QueryContext(ctx, query, userID, lastID, limit)
	if err != nil {
		return customerror.NewError(
			"something went wrong",
			err,
			customerror.DatabaseExecutionErr,
		)
	}
	defer rows.Close()

	for rows.Next() {
		var l URL
		if err := rows.Scan(
			&l.ID,
			&l.UserID,
			&l.EncryptedLongUrl,
			&l.CreatedAt,
			&l.UpdatedAt,
			&l.DeletedAt,
			&l.ExpiredAt,
		); err != nil {
			return customerror.NewError(
				"something went wrong",
				err,
				customerror.DatabaseExecutionErr,
			)
		}
		*links = append(*links, l)
	}

	if err := rows.Err(); err != nil {
		return customerror.NewError(
			"something went wrong",
			err,
			customerror.DatabaseExecutionErr,
		)
	}

	return nil
}

func (sur *URLRepoImpl) InsertNewURL(
	ctx context.Context,
	userID *string,
	encryptedLongURL string,
	expiredAt *time.Time,
	shortenURL *URL,
) error {
	query := `
	INSERT INTO urls (user_id, encrypted_long_url, expired_at)
	VALUES ($1, $2, $3)
	RETURNING id, user_id, encrypted_long_url, created_at, updated_at, deleted_at, expired_at
	`
	if err := sur.handle.QueryRowContext(ctx, query, userID, encryptedLongURL, expiredAt).Scan(
		&shortenURL.ID,
		&shortenURL.UserID,
		&shortenURL.EncryptedLongUrl,
		&shortenURL.CreatedAt,
		&shortenURL.UpdatedAt,
		&shortenURL.DeletedAt,
		&shortenURL.ExpiredAt,
	); err != nil {
		return customerror.NewError(
			"something went wrong",
			err,
			customerror.DatabaseExecutionErr,
		)
	}
	return nil
}

func (sur *URLRepoImpl) FindByID(ctx context.Context, id int64, url *URL) error {
	query := `
	SELECT
		id,
		user_id,
		encrypted_long_url,
		created_at,
		updated_at,
		deleted_at,
		expired_at
	FROM urls
	WHERE id = $1 AND deleted_at IS NULL
	`

	if err := sur.handle.QueryRowContext(ctx, query, id).Scan(
		&url.ID,
		&url.UserID,
		&url.EncryptedLongUrl,
		&url.CreatedAt,
		&url.UpdatedAt,
		&url.DeletedAt,
		&url.ExpiredAt,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return customerror.NewError(
				"invalid url",
				err,
				customerror.ItemNotFound,
			)
		}
		return customerror.NewError(
			"something went wrong",
			err,
			customerror.DatabaseExecutionErr,
		)
	}
	return nil
}
