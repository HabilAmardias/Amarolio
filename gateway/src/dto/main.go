package dto

type (
	ServerResponse[T interface{}] struct {
		Success bool `json:"success"`
		Data    T    `json:"data"`
	}
	ErrorResponse struct {
		Detail string `json:"detail"`
	}
)
