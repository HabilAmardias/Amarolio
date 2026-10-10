package services

import (
	"amarolio-auth/src/customerrors"
	"amarolio-auth/src/db"
)

func WithTransaction(db *db.DBHandle, fn func(tx db.DBTXItf) error) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if err := fn(tx); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return customerrors.NewError(
			"something went wrong",
			err,
			customerrors.DatabaseExecutionErr,
		)
	}
	return nil
}
