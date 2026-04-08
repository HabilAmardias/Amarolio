package messages

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type PubSubKey struct{}

type MessageChannelRepositoryImpl struct {
	rc *redis.Client
}

func (mcr *MessageChannelRepositoryImpl) PublishMessage(ctx context.Context, userID string, chatroomID string, message string) {
	args := &redis.XAddArgs{
		Stream: "event",
		Approx: true,
		Values: map[string]any{
			"action":   "send",
			"user":     userID,
			"chatroom": chatroomID,
			"message":  message,
			"time":     time.Now().Format(time.RFC3339),
		},
	}
	mcr.rc.XAdd(ctx, args)
}

func (mcr *MessageChannelRepositoryImpl) Subscribe(ctx context.Context, chatroomID string) *redis.PubSub {
	rps := mcr.rc.Subscribe(ctx, fmt.Sprintf("chatrooms:%s:messages", chatroomID))
	return rps
}
