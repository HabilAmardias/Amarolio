package entity

type (
	Paginate[T any] struct {
		Entries  []T `json:"entries"`
		PageInfo struct {
			LastID   *int64 `json:"last_id"`
			Page     *int64 `json:"page"`
			Limit    int64  `json:"limit"`
			FilterBy []struct {
				Name  string `json:"name"`
				Value any    `json:"value"`
			} `json:"filter_by"`
			SortBy []struct {
				Name   string `json:"name"`
				Ascend bool   `json:"ascend"`
			} `json:"sort_by"`
		} `json:"page_info"`
	}
	PlainMessageResponse struct {
		Message string `json:"message"`
	}
)
