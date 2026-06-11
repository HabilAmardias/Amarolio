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
	TurnstileReq struct {
		Token string `query:"token" validate:"required"`
	}
	SeekPaginateReq struct {
		LastID *int64 `query:"last_id"`
		Limit  *int64 `query:"limit" validate:"omitempty,lte=26"`
	}
	PaginateOffsetReq struct {
		Page  int64 `query:"page" validate:"required"`
		Limit int64 `query:"limit" validate:"required,lte=26"`
	}
	PaginateRes[T any] struct {
		Entries  []T `json:"entries"`
		PageInfo struct {
			LastID   *int64 `json:"last_id,omitempty"`
			Page     *int64 `json:"page,omitempty"`
			Limit    int64  `json:"limit"`
			FilterBy []struct {
				Name  string `json:"name"`
				Value any    `json:"value"`
			} `json:"filter_by,omitempty"`
			SortBy []struct {
				Name   string `json:"name"`
				Ascend bool   `json:"ascend"`
			} `json:"sort_by,omitempty"`
		} `json:"page_info"`
	}
	PlainMessageRes struct {
		Message string `json:"message"`
	}
)

func (de *DetailsError) ToString() string {
	return fmt.Sprintf("%s:%s", de.Field, de.Message)
}
