package messages

import "time"

type (
	Message struct {
		ID         int        `json:"id"`
		ChatroomID string     `json:"chatroom_id"`
		Role       string     `json:"role"`
		Content    string     `json:"content"`
		CreatedAt  time.Time  `json:"created_at"`
		UpdatedAt  time.Time  `json:"updated_at"`
		DeletedAt  *time.Time `json:"deleted_at"`
	}
	SendMessage struct {
		Message string `json:"message"`
	}
	SendMessageBody struct {
		Message string `json:"user_message"`
	}
	GetMessages struct {
		Messages []Message `json:"messages"`
	}
)
