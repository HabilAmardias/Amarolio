package transactionschedulers

import (
	"amarfin/src/customerror"
	"amarfin/src/db"
	"context"
	"time"
)

type TransactionSchedulerRepository struct {
	handle db.DBTX
}

func NewTransactionSchedulerRepository(handle db.DBTX) *TransactionSchedulerRepository {
	return &TransactionSchedulerRepository{handle: handle}
}

func (tsr *TransactionSchedulerRepository) Delete(ctx context.Context, id string) error {
	query := `
	UPDATE transaction_schedulers
	SET deleted_at = $1
	WHERE id = $2 AND deleted_at IS NULL
	`
	_, err := tsr.handle.ExecContext(ctx, query, time.Now().String(), id)
	if err != nil {
		return customerror.NewError(
			"something went wrong",
			err,
			customerror.DatabaseExecutionErr,
		)
	}
	return nil
}

func (tsr *TransactionSchedulerRepository) Create(
	ctx context.Context,
	id string,
	frequency string,
	dateOfMonth int,
	timezone string,
	localTime string,
	startAt string,
	endAt *string,
	amount uint64,
	userID string,
	description string,
	nextRunAt string,
) (TransactionScheduler, error) {
	ts := TransactionScheduler{}
	now := time.Now().String()
	query := `
	INSERT INTO transaction_schedulers (id, frequency, date_of_month, timezone, local_time, start_at, end_at, amount, user_id, description, next_run_at, created_at, updated_at, deleted_at)
	VALUES
	($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)
	RETURNING id, frequency, date_of_month, timezone, local_time, start_at, end_at, amount, user_id, description, next_run_at, created_at, updated_at, deleted_at
	`
	if err := tsr.handle.QueryRowContext(
		ctx,
		query,
		id,
		frequency,
		dateOfMonth,
		timezone,
		localTime,
		startAt,
		endAt,
		amount,
		userID,
		description,
		nextRunAt,
		now,
		now,
		nil,
	).Scan(
		&ts.ID,
		&ts.Frequency,
		&ts.DateOfMonth,
		&ts.TimeZone,
		&ts.LocalTime,
		&ts.StartAt,
		&ts.EndAt,
		&ts.Amount,
		&ts.UserID,
		&ts.Description,
		&ts.NextRunAt,
		&ts.CreatedAt,
		&ts.UpdatedAt,
		&ts.DeletedAt,
	); err != nil {
		return ts, customerror.NewError(
			"something went wrong",
			err,
			customerror.DatabaseExecutionErr,
		)
	}
	return ts, nil
}
