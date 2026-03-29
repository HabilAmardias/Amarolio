package chatrooms

import (
	"amarolio-gateway/src/constants"
	"amarolio-gateway/src/customerrors"
	"amarolio-gateway/src/dto"
	"amarolio-gateway/src/handlers"

	"github.com/gofiber/fiber/v3"
)

type ChatroomServiceItf interface {
	GetChatrooms(userID string, limit *int, page *int) ([]Chatroom, error)
	CreateChatroom(userID string, userMessage string) (string, error)
	DeleteChatroom(userID string, chatroomID string) error
}

type ChatroomHandlerImpl struct {
	cs ChatroomServiceItf
}

func NewChatroomHandler(cs ChatroomServiceItf) *ChatroomHandlerImpl {
	return &ChatroomHandlerImpl{cs}
}

func (ch *ChatroomHandlerImpl) DeleteChatroom(ctx fiber.Ctx) error {
	claim, err := handlers.GetAuthPayload(ctx, constants.AUTH_CLAIM_KEY)
	if err != nil {
		return err
	}
	chatroomID := ctx.Params("id")
	if err := ch.cs.DeleteChatroom(claim.Subject, chatroomID); err != nil {
		return err
	}
	ctx.Status(fiber.StatusOK)
	return ctx.JSON(dto.ServerResponse[DeleteChatroomRes]{
		Success: true,
		Data: DeleteChatroomRes{
			ID: chatroomID,
		},
	})
}

func (ch *ChatroomHandlerImpl) CreateChatroom(ctx fiber.Ctx) error {
	claim, err := handlers.GetAuthPayload(ctx, constants.AUTH_CLAIM_KEY)
	if err != nil {
		return err
	}
	req := new(CreateChatroomReq)
	if err := ctx.Bind().JSON(req); err != nil {
		return customerrors.NewError(
			"something went wrong",
			err,
			customerrors.CommonErr,
		)
	}
	id, err := ch.cs.CreateChatroom(claim.Subject, req.UserMessage)
	if err != nil {
		return err
	}

	ctx.Status(fiber.StatusCreated)
	return ctx.JSON(dto.ServerResponse[CreateChatroomRes]{
		Success: true,
		Data: CreateChatroomRes{
			ID: id,
		},
	})
}

func (ch *ChatroomHandlerImpl) GetChatrooms(ctx fiber.Ctx) error {
	claim, err := handlers.GetAuthPayload(ctx, constants.AUTH_CLAIM_KEY)
	if err != nil {
		return err
	}
	req := new(GetChatroomsReq)
	if err := ctx.Bind().Query(req); err != nil {
		return customerrors.NewError(
			"something went wrong",
			err,
			customerrors.CommonErr,
		)
	}
	chatrooms, err := ch.cs.GetChatrooms(claim.Subject, req.Limit, req.Page)
	if err != nil {
		return err
	}
	res := []ChatroomRes{}
	for _, v := range chatrooms {
		res = append(res, ChatroomRes{
			ID:     v.ID,
			Title:  v.Title,
			UserID: v.UserID,
		})
	}
	ctx.Status(fiber.StatusOK)
	return ctx.JSON(dto.ServerResponse[GetChatroomsRes]{
		Success: true,
		Data: GetChatroomsRes{
			Chatrooms: res,
		},
	})
}
