package messages

import (
	"amarolio-gateway/src/constants"
	"amarolio-gateway/src/customerrors"
	"amarolio-gateway/src/dto"
	"amarolio-gateway/src/services"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/valyala/fasthttp"
)

type MessageServiceImpl struct {
	hs string
	pr string
}

func NewMessageService(hs string, pr string) *MessageServiceImpl {
	return &MessageServiceImpl{hs, pr}
}

func (ms *MessageServiceImpl) callSendMessage(userID string, chatroomID string, userMessage string) (*dto.ServerResponse[SendMessage], error) {
	path := fmt.Sprintf("/api/v1/chatrooms/%s/messages", chatroomID)
	reqBody, err := json.Marshal(SendMessageBody{
		Message: userMessage,
	})
	if err != nil {
		return nil, customerrors.NewError(
			"something went wrong",
			err,
			customerrors.CommonErr,
		)
	}
	headers := map[string]string{
		constants.X_USER_ID: userID,
	}
	return services.Call[SendMessage](ms.hs, ms.pr, path, fasthttp.MethodPost, fasthttp.StatusCreated, reqBody, nil, headers)
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

func (ms *MessageServiceImpl) SendMessage(userID string, chatroomID string, userMessage string) (string, error) {
	res, err := ms.callSendMessage(userID, chatroomID, userMessage)
	if err != nil {
		return "", err
	}
	return res.Data.Message, nil
}
