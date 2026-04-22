package chatrooms

import "time"

type (
	Chatroom struct {
		ID        string     `json:"id"`
		UserID    string     `json:"user_id"`
		Title     string     `json:"title"`
		CreatedAt time.Time  `json:"created_at"`
		UpdatedAt time.Time  `json:"updated_at"`
		DeletedAt *time.Time `json:"deleted_at"`
	}
	GetChatrooms struct {
		Chatrooms []Chatroom `json:"chatrooms"`
	}
	CreateChatroomBody struct {
		Message string `json:"user_message"`
	}
	CreateChatroom struct {
		ID string `json:"id"`
	}
	DeleteChatroom struct {
		Message string `json:"message"`
	}
)
