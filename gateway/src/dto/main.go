package dto

import "fmt"

type (
	ServerResponse[T interface{}] struct {
		Success bool `json:"success"`
		Data    T    `json:"data"`
	}
	ErrorResponse struct {
		Detail string `json:"detail"`
	}
	DetailsError struct {
		Field   string `json:"field"`
		Message string `json:"message"`
	}
)

func (de *DetailsError) ToString() string {
	return fmt.Sprintf("%s:%s", de.Field, de.Message)
}
