package url

import (
	"amary/src/customerror"
	"amary/src/db"
	"amary/src/repository"
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

func (sur *URLRepoImpl) FindByCode(ctx context.Context, shortCode string, url *URL) error {
	var driver db.DBTX = sur.handle
	if tx := repository.GetTransactionFromCtx(ctx); tx != nil {
		driver = tx
	}
	query := `
	SELECT
		id,
		user_id,
		encrypted_long_url,
		short_code,
		created_at,
		updated_at,
		deleted_at,
		expired_at
	FROM urls
	WHERE short_code = $1 AND deleted_at IS NULL
	`
	if err := driver.QueryRowContext(ctx, query, shortCode).Scan(
		&url.ID,
		&url.UserID,
		&url.EncryptedLongUrl,
		&url.ShortCode,
		&url.CreatedAt,
		&url.UpdatedAt,
		&url.DeletedAt,
		&url.ExpiredAt,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return customerror.NewError(
				"url not found",
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

func (sur *URLRepoImpl) UpdateShortCode(ctx context.Context, id int64, shortCode string, url *URL) error {
	query := `
	UPDATE urls
	SET short_code = $1, updated_at = CURRENT_TIMESTAMP
	WHERE id = $2 AND deleted_at IS NULL
	RETURNING id, user_id, encrypted_long_url, short_code, created_at, updated_at, deleted_at, expired_at
	`
	var driver db.DBTX = sur.handle
	if tx := repository.GetTransactionFromCtx(ctx); tx != nil {
		driver = tx
	}
	if err := driver.QueryRowContext(ctx, query, shortCode, id).Scan(
		&url.ID,
		&url.UserID,
		&url.EncryptedLongUrl,
		&url.ShortCode,
		&url.CreatedAt,
		&url.UpdatedAt,
		&url.DeletedAt,
		&url.ExpiredAt,
	); err != nil {
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
		short_code,
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
	var driver db.DBTX = sur.handle
	if tx := repository.GetTransactionFromCtx(ctx); tx != nil {
		driver = tx
	}
	rows, err := driver.QueryContext(ctx, query, userID, lastID, limit)
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
			&l.ShortCode,
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
	var driver db.DBTX = sur.handle
	if tx := repository.GetTransactionFromCtx(ctx); tx != nil {
		driver = tx
	}
	if err := driver.QueryRowContext(ctx, query, userID, encryptedLongURL, expiredAt).Scan(
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
		short_code,
		created_at,
		updated_at,
		deleted_at,
		expired_at
	FROM urls
	WHERE id = $1 AND deleted_at IS NULL
	`
	var driver db.DBTX = sur.handle
	if tx := repository.GetTransactionFromCtx(ctx); tx != nil {
		driver = tx
	}
	if err := driver.QueryRowContext(ctx, query, id).Scan(
		&url.ID,
		&url.UserID,
		&url.EncryptedLongUrl,
		&url.ShortCode,
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
