package chatrooms

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

type ChatroomServiceImpl struct {
	hs string
	pr string
}

func NewChatroomService(hs string, pr string) *ChatroomServiceImpl {
	return &ChatroomServiceImpl{hs, pr}
}

func (cs *ChatroomServiceImpl) callCreateChatroom(userID string, userMessage string) (*dto.ServerResponse[CreateChatroom], error) {
	b := CreateChatroomBody{
		Message: userMessage,
	}
	header := map[string]string{
		constants.X_USER_ID: userID,
	}
	path := "/api/v1/chatrooms"

	reqBody, err := json.Marshal(b)
	if err != nil {
		return nil, customerrors.NewError(
			"something went wrong",
			err,
			customerrors.CommonErr,
		)
	}

	return services.Call[CreateChatroom](cs.hs, cs.pr, path, fasthttp.MethodPost, fasthttp.StatusCreated, reqBody, nil, header)
}

func (cs *ChatroomServiceImpl) callDeleteChatroom(userID string, chatroomID string) (*dto.ServerResponse[DeleteChatroom], error) {
	path := fmt.Sprintf("/api/v1/chatrooms/%s", chatroomID)
	header := map[string]string{
		constants.X_USER_ID: userID,
	}

	return services.Call[DeleteChatroom](cs.hs, cs.pr, path, fasthttp.MethodDelete, fasthttp.StatusOK, nil, nil, header)
}

func (cs *ChatroomServiceImpl) callGetChatrooms(userID string, limit *int, page *int) (*dto.ServerResponse[GetChatrooms], error) {
	path := "/api/v1/chatrooms"
	header := map[string]string{
		constants.X_USER_ID: userID,
	}
	queries := make(map[string]string)
	if limit != nil {
		queries["limit"] = strconv.Itoa(*limit)
	}
	if page != nil {
		queries["page"] = strconv.Itoa(*page)
	}

	return services.Call[GetChatrooms](cs.hs, cs.pr, path, fasthttp.MethodGet, fasthttp.StatusOK, nil, queries, header)
}

func (cs *ChatroomServiceImpl) GetChatrooms(userID string, limit *int, page *int) ([]Chatroom, error) {
	res, err := cs.callGetChatrooms(userID, limit, page)
	if err != nil {
		return nil, err
	}
	return res.Data.Chatrooms, nil
}

func (cs *ChatroomServiceImpl) CreateChatroom(userID string, userMessage string) (string, error) {
	res, err := cs.callCreateChatroom(userID, userMessage)
	if err != nil {
		return "", err
	}
	return res.Data.ID, nil
}

func (cs *ChatroomServiceImpl) DeleteChatroom(userID string, chatroomID string) error {
	_, err := cs.callDeleteChatroom(userID, chatroomID)
	return err
}
