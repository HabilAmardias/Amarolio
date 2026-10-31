package dto

import "fmt"

type (
	ServerResponse[T any] struct {
		Success bool `json:"success"`
		Data    T    `json:"data"`
	}
	ErrorResponse struct {
		Detail    string `json:"detail"`
		ErrorCode int    `json:"error_code"`
	}
	DetailsError struct {
		Field   string `json:"field"`
		Message string `json:"message"`
	}
)

func (de *DetailsError) ToString() string {
	return fmt.Sprintf("%s:%s", de.Field, de.Message)
}
