package messages

type (
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
