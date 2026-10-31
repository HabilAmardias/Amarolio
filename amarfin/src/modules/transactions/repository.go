package transactions

import (
	"amarfin/src/customerror"
	"amarfin/src/db"
	"context"
	"database/sql"
	"errors"
	"time"
)

type TransactionRepository struct {
	handle db.DBTX
}

func NewTransactionRepository(handle db.DBTX) *TransactionRepository {
	return &TransactionRepository{handle}
}

func (tr *TransactionRepository) FindByID(ctx context.Context, id string) (Transaction, error) {
	t := Transaction{}
	query := `
	SELECT
		id,
		amount,
		user_id,
		description,
		scheduler_id,
		scheduled_for,
		created_at,
		updated_at,
		deleted_at
	FROM transactions
	WHERE id = $1 AND deleted_at IS NULL
	`
	if err := tr.handle.QueryRowContext(ctx, query, id).Scan(
		&t.ID,
		&t.Amount,
		&t.UserID,
		&t.Description,
		&t.SchedulerID,
		&t.ScheduledFor,
		&t.CreatedAt,
		&t.UpdatedAt,
		&t.DeletedAt,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return t, customerror.NewError(
				"transaction not found",
				err,
				customerror.ItemNotFound,
			)
		}
		return t, customerror.NewError(
			"something went wrong",
			err,
			customerror.DatabaseExecutionErr,
		)
	}
	return t, nil
}

func (tr *TransactionRepository) Delete(ctx context.Context, id string) error {
	now := time.Now().String()
	query := `
	UPDATE transactions
	SET deleted_at = $1
	WHERE id = $2 AND deleted_at IS NULL
	`
	_, err := tr.handle.ExecContext(ctx, query, now, id)
	if err != nil {
		return customerror.NewError(
			"something went wrong",
			err,
			customerror.DatabaseExecutionErr,
		)
	}
	return nil
}

func (tr *TransactionRepository) Create(
	ctx context.Context,
	id string,
	amount uint64,
	userID string,
	description string,
	schedulerID *string,
	scheduledFor *string,
) (Transaction, error) {
	t := Transaction{}
	now := time.Now().String()

	query := `
	INSERT INTO transactions (id, amount, user_id, description, scheduler_id, scheduled_for, created_at, updated_at, deleted_at)
	VALUES
	($1, $2, $3, $4, $5, $6, $7, $8, $9)
	RETURNING id, amount, user_id, description, scheduler_id, scheduled_for, created_at, updated_at, deleted_at
	`

	if err := tr.handle.QueryRowContext(
		ctx,
		query,
		id,
		amount,
		userID,
		description,
		schedulerID,
		scheduledFor,
		now,
		now,
		nil,
	).Scan(
		&t.ID,
		&t.Amount,
		&t.UserID,
		&t.Description,
		&t.SchedulerID,
		&t.ScheduledFor,
		&t.CreatedAt,
		&t.UpdatedAt,
		&t.DeletedAt,
	); err != nil {
		return t, customerror.NewError(
			"something went wrong",
			err,
			customerror.DatabaseExecutionErr,
		)
	}
	return t, nil
}
