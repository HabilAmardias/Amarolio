package services

import (
	"amary/src/db"
	"context"
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
