package transactions

type (
	CreateReq struct {
		Amount      uint64 `json:"amount" validate:"gte=1"`
		Description string `json:"description" validate:"min=1"`
	}
	CreateRes struct {
		ID          string `json:"id"`
		Amount      uint64 `json:"amount"`
		Description string `json:"description"`
	}
	IDRes struct {
		ID string `json:"id"`
	}
)
