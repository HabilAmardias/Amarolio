package transactions

import (
	"amarfin/src/dto"
	"amarfin/src/handler"
	"context"
	"net/http"

	"github.com/google/uuid"
)

type TransactionUseCaseItf interface {
	Delete(
		ctx context.Context,
		id string,
		userID string,
	) error
	Create(
		ctx context.Context,
		amount uint64,
		userID string,
		description string,
		schedulerID *string,
		scheduledFor *string,
	) (Transaction, error)
}

type TransactionHandler struct {
	tus TransactionUseCaseItf
}

func NewTransactionHandler(tus TransactionUseCaseItf) *TransactionHandler {
	return &TransactionHandler{tus}
}

func (th *TransactionHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	_, err := uuid.Parse(id)
	if err != nil {
		r = handler.Error(err, r)
		return
	}

	uid, err := handler.GetAuth(r)
	if err != nil {
		r = handler.Error(err, r)
		return
	}

	if err := th.tus.Delete(r.Context(), id, uid); err != nil {
		r = handler.Error(err, r)
		return
	}
	res := dto.ServerResponse[IDRes]{
		Success: true,
		Data:    IDRes{ID: id},
	}
	if err := handler.WriteJSON(res, http.StatusOK, w); err != nil {
		r = handler.Error(err, r)
	}
}

func (th *TransactionHandler) Create(w http.ResponseWriter, r *http.Request) {
	req := new(CreateReq)
	if err := handler.ReadJSON(req, r); err != nil {
		r = handler.Error(err, r)
		return
	}

	uid, err := handler.GetAuth(r)
	if err != nil {
		r = handler.Error(err, r)
		return
	}

	t, err := th.tus.Create(r.Context(), req.Amount, uid, req.Description, nil, nil)
	if err != nil {
		r = handler.Error(err, r)
		return
	}

	res := dto.ServerResponse[CreateRes]{
		Success: true,
		Data: CreateRes{
			ID:          t.ID,
			Amount:      t.Amount,
			Description: t.Description,
		},
	}

	if err := handler.WriteJSON(res, http.StatusCreated, w); err != nil {
		r = handler.Error(err, r)
	}
}
