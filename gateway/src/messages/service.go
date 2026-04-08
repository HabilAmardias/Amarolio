package messages

import (
	"amarolio-gateway/src/constants"
	"amarolio-gateway/src/dto"
	"amarolio-gateway/src/services"
	"context"
	"fmt"
	"strconv"

	"github.com/redis/go-redis/v9"
	"github.com/valyala/fasthttp"
)

type MessageChannelRepositoryItf interface {
	Subscribe(ctx context.Context, chatroomID string) *redis.PubSub
	PublishMessage(ctx context.Context, userID string, chatroomID string, message string)
}
type MessageServiceImpl struct {
	hs  string
	pr  string
	mcr MessageChannelRepositoryItf
}

func NewMessageService(hs string, pr string, mcr MessageChannelRepositoryItf) *MessageServiceImpl {
	return &MessageServiceImpl{hs, pr, mcr}
}

func (ms *MessageServiceImpl) callGetMessages(userID string, chatroomID string, limit *int, lastID *int) (*dto.ServerResponse[GetMessages], error) {
	path := fmt.Sprintf("/api/v1/chatrooms/%s/messages", chatroomID)
	headers := map[string]string{
		constants.X_USER_ID: userID,
	}
	queries := make(map[string]string)
	if limit != nil {
		queries["limit"] = strconv.Itoa(*limit)
	}
	if lastID != nil {
		queries["last_id"] = strconv.Itoa(*lastID)
	}
	return services.Call[GetMessages](ms.hs, ms.pr, path, fasthttp.MethodGet, fasthttp.StatusOK, nil, nil, headers)
}

func (ms *MessageServiceImpl) GetMessages(userID string, chatroomID string, limit *int, lastID *int) ([]Message, error) {
	res, err := ms.callGetMessages(userID, chatroomID, limit, lastID)
	if err != nil {
		return nil, err
	}
	return res.Data.Messages, nil
}

func (ms *MessageServiceImpl) Subscribe(ctx context.Context, chatroomID string, cc *ChatClient) {
	rps := ms.mcr.Subscribe(ctx, chatroomID)
	cc.Subscribe(rps)
}

func (ms *MessageServiceImpl) Write(ctx context.Context, cc *ChatClient) {
	cc.Write(ctx)
}

func (ms *MessageServiceImpl) Read(ctx context.Context, cc *ChatClient, userID string, chatroomID string) {
	cc.Read(ctx, userID, chatroomID, ms.mcr.PublishMessage)
}
