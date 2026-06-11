package repository

import (
	"amary/src/customerror"
	"amary/src/db"
	"context"
	"database/sql"
)

type TxKey struct{}

func GetTransactionFromCtx(ctx context.Context) *sql.Tx {
	tx, ok := ctx.Value(TxKey{}).(*sql.Tx)
	if !ok {
		return nil
	}
	return tx
}

type TransactionManagerImpl struct {
	db *db.DBHandle
}

func NewTransactionManager(db *db.DBHandle) *TransactionManagerImpl {
	return &TransactionManagerImpl{db}
}

func (tr *TransactionManagerImpl) WithTransaction(ctx context.Context, fun func(c context.Context) error) error {
	tx, err := tr.db.Begin(ctx, &sql.TxOptions{})
	if err != nil {
		return customerror.NewError(
			"something went wrong",
			err,
			customerror.DatabaseExecutionErr,
		)
	}
	defer tx.Rollback()

	newCtx := context.WithValue(ctx, TxKey{}, tx)
	if err := fun(newCtx); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return customerror.NewError(
			"something went wrong",
			err,
			customerror.DatabaseExecutionErr,
		)
	}
	return nil
}
