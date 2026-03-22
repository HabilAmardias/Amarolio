package dto

type (
	ServerResponse struct {
		Success bool        `json:"success"`
		Data    interface{} `json:"data"`
	}
	ErrorResponse struct {
		Detail interface{} `json:"detail"`
	}
	DetailsError struct {
		Field   string `json:"field"`
		Message string `json:"message"`
	}
)
