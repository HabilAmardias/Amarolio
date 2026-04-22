package chatrooms

type (
	ChatroomRes struct {
		ID     string `json:"id"`
		UserID string `json:"user_id"`
		Title  string `json:"title"`
	}
	GetChatroomsReq struct {
		Limit *int `query:"limit"`
		Page  *int `query:"page"`
	}
	GetChatroomsRes struct {
		Chatrooms []ChatroomRes `json:"chatrooms"`
	}
	CreateChatroomReq struct {
		UserMessage string `json:"user_message"`
	}
	CreateChatroomRes struct {
		ID string `json:"id"`
	}
	DeleteChatroomRes struct {
		ID string `json:"id"`
	}
)
