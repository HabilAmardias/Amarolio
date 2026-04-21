package services

import (
	"amary/src/db"
	"context"
	"time"
)

func WithTransaction(ctx context.Context, db *db.DBHandle, callback func(tx db.DBTX) error) error {
	tx, err := db.Begin(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	if err := callback(tx); err != nil {
		return err
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

func WithErrorRetry(ctx context.Context, fun func() error, delay time.Duration) {
	maxRetry := 3
	for attempt := 0; attempt <= maxRetry; attempt++ {
		if err := fun(); err == nil {
			return
		}
		if attempt == maxRetry {
			break
		}

		select {
		case <-ctx.Done():
			return
		case <-time.After(delay):
		}
	}
}
