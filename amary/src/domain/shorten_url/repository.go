package shortenurl

import (
	"amary/src/customerror"
	"amary/src/db"
	"context"
	"database/sql"
	"errors"
	"time"
)

type ShortenURLRepoImpl struct {
	handle db.DBTX
}

func NewShortenURLRepo(handle db.DBTX) *ShortenURLRepoImpl {
	return &ShortenURLRepoImpl{handle}
}

func (sur *ShortenURLRepoImpl) InsertNewURL(
	ctx context.Context,
	userID *string,
	encryptedLongURL string,
	expiredAt *time.Time,
	shortenURL *ShortenURL,
) error {
	query := `
	INSERT INTO shorten_urls (user_id, encrypted_long_url, expired_at)
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

func (sur *ShortenURLRepoImpl) FindByID(ctx context.Context, id int64, url *ShortenURL) error {
	query := `
	SELECT
		id,
		user_id,
		encrypted_long_url,
		created_at,
		updated_at,
		deleted_at,
		expired_at
	FROM shorten_urls
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
