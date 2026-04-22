package messages

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type MessageBridgeRepository struct {
	rc *redis.Client
}

func NewMessageChannelRepository(rc *redis.Client) *MessageBridgeRepository {
	return &MessageBridgeRepository{rc}
}

func (mcr *MessageBridgeRepository) PublishMessage(ctx context.Context, userID string, chatroomID string, message string) {
	args := &redis.XAddArgs{
		Stream: "streams:chatrooms",
		Approx: true,
		Values: map[string]string{
			"action":   "send_message",
			"user":     userID,
			"chatroom": chatroomID,
			"message":  message,
			"time":     time.Now().Format(time.RFC3339),
		},
	}
	mcr.rc.XAdd(ctx, args)
}

func (mcr *MessageBridgeRepository) Subscribe(ctx context.Context, chatroomID string) *redis.PubSub {
	rps := mcr.rc.Subscribe(ctx, fmt.Sprintf("chatrooms:%s:messages", chatroomID))
	return rps
}
