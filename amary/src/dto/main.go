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
		Detail    string `json:"detail"`
		ErrorCode int    `json:"error_code"`
	}
	PlainMessageResponse struct {
		Message string `json:"message"`
	}
	PaginateOffsetReq struct {
		Page  int64 `form:"page" binding:"required"`
		Limit int64 `form:"limit" binding:"required,lte=26"`
	}
	SeekPaginateReq struct {
		LastID *int64 `form:"last_id"`
		Limit  *int64 `form:"limit" binding:"omitempty,lte=26"`
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
)

func (de *DetailsError) ToString() string {
	return fmt.Sprintf("%s:%s", de.Field, de.Message)
}
