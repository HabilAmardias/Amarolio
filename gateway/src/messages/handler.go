package messages

import (
	"amarolio-gateway/src/constants"
	"amarolio-gateway/src/customerrors"
	"amarolio-gateway/src/dto"
	"amarolio-gateway/src/handlers"
	"errors"

	"github.com/gofiber/fiber/v3"
	"github.com/valyala/fasthttp"
)

type MessageServiceItf interface {
	SendMessage(userID string, chatroomID string, userMessage string) (string, error)
	GetMessages(userID string, chatroomID string, limit *int, lastID *int) ([]Message, error)
}

type MessageHandlerImpl struct {
	ms MessageServiceItf
}

func NewMessageHandler(ms MessageServiceItf) *MessageHandlerImpl {
	return &MessageHandlerImpl{ms}
}

func (mh *MessageHandlerImpl) GetMessages(ctx fiber.Ctx) error {
	claim, err := handlers.GetAuthPayload(ctx, constants.AUTH_CLAIM_KEY)
	if err != nil {
		return err
	}
	chatroomID := ctx.Params("id")
	if len(chatroomID) == 0 {
		return customerrors.NewError(
			"data not found",
			errors.New("chatroom id does not found"),
			customerrors.InvalidAction,
		)
	}
	req := new(GetMessagesReq)
	if err := ctx.Bind().Query(req); err != nil {
		return err
	}
	msgs, err := mh.ms.GetMessages(claim.Subject, chatroomID, req.Limit, req.LastID)
	res := []MessageRes{}
	for _, m := range msgs {
		res = append(res, MessageRes{
			ID:         m.ID,
			ChatroomID: m.ChatroomID,
			Role:       m.Role,
			Content:    m.Content,
		})
	}
	ctx.Status(fasthttp.StatusOK)
	return ctx.JSON(dto.ServerResponse[GetMessagesRes]{
		Success: true,
		Data: GetMessagesRes{
			Messages: res,
		},
	})
}

func (mh *MessageHandlerImpl) SendMessage(ctx fiber.Ctx) error {
	claim, err := handlers.GetAuthPayload(ctx, constants.AUTH_CLAIM_KEY)
	if err != nil {
		return err
	}
	chatroomID := ctx.Params("id")
	if len(chatroomID) == 0 {
		return customerrors.NewError(
			"data not found",
			errors.New("chatroom id does not found"),
			customerrors.InvalidAction,
		)
	}
	req := new(SendMessageReq)
	if err := ctx.Bind().Body(req); err != nil {
		return err
	}
	res, err := mh.ms.SendMessage(claim.Subject, chatroomID, req.UserMessage)
	if err != nil {
		return err
	}
	ctx.Status(fasthttp.StatusCreated)
	return ctx.JSON(dto.ServerResponse[SendMessageRes]{
		Success: true,
		Data: SendMessageRes{
			Message: res,
		},
	})
}
