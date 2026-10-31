package transactions

import (
	"amarfin/src/customerror"
	"amarfin/src/db"
	"context"
	"errors"

	"github.com/google/uuid"
)

type TransactionUseCase struct {
	dbHandle *db.DBHandle
}

func NewTransactionUseCase(db *db.DBHandle) *TransactionUseCase {
	return &TransactionUseCase{dbHandle: db}
}

func (tus *TransactionUseCase) Delete(
	ctx context.Context,
	id string,
	userID string,
) error {
	_, err := uuid.Parse(id)
	if err != nil {
		return customerror.NewError(
			"invalid transaction",
			err,
			customerror.InvalidAction,
		)
	}

	_, err = uuid.Parse(userID)
	if err != nil {
		return customerror.NewError(
			"please login first",
			err,
			customerror.Unauthenticate,
		)
	}

	tr := NewTransactionRepository(tus.dbHandle)

	t, err := tr.FindByID(ctx, id)
	if err != nil {
		return err
	}
	if t.UserID != userID {
		return customerror.NewError(
			"unauthorized access",
			errors.New("transaction does not belong to the user"),
			customerror.ForbiddenAction,
		)
	}

	return tr.Delete(ctx, id)
}

func (tus *TransactionUseCase) Create(
	ctx context.Context,
	amount uint64,
	userID string,
	description string,
	schedulerID *string,
	scheduledFor *string,
) (Transaction, error) {

	id, err := uuid.NewRandom()
	if err != nil {
		return Transaction{}, customerror.NewError(
			"something went wrong",
			err,
			customerror.CommonErr,
		)
	}
	idParsed := id.String()

	_, err = uuid.Parse(userID)
	if err != nil {
		return Transaction{}, customerror.NewError(
			"invalid user",
			err,
			customerror.Unauthenticate,
		)
	}

	tr := NewTransactionRepository(tus.dbHandle)

	return tr.Create(
		ctx,
		idParsed,
		amount,
		userID,
		description,
		schedulerID,
		scheduledFor,
	)
}
