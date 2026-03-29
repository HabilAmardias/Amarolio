package messages

type (
	SendMessageReq struct {
		UserMessage string `json:"user_message"`
	}
	SendMessageRes struct {
		Message string `json:"message"`
	}
	GetMessagesReq struct {
		Limit  *int `query:"limit"`
		LastID *int `query:"last_id"`
	}
	MessageRes struct {
		ID         int    `json:"id"`
		ChatroomID string `json:"chatroom_id"`
		Role       string `json:"role"`
		Content    string `json:"content"`
	}
	GetMessagesRes struct {
		Messages []MessageRes `json:"messages"`
	}
)
