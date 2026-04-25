package dto

import "fmt"

type (
	ServerResponse[T any] struct {
		Success bool `json:"success"`
		Data    T    `json:"data"`
	}
	DetailsError struct {
		Field   string `json:"field"`
		Message string `json:"message"`
	}
	ErrorResponse struct {
		Detail string `json:"detail"`
	}
	PaginateOffsetReq struct {
		Page  int64 `form:"page" binding:"required"`
		Limit int64 `form:"limit" binding:"required,lte=25"`
	}
)

func (de *DetailsError) ToString() string {
	return fmt.Sprintf("%s:%s", de.Field, de.Message)
}
