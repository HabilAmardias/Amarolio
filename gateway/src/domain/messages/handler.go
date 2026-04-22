package messages

import (
	"amarolio-gateway/src/constants"
	"amarolio-gateway/src/customerrors"
	"amarolio-gateway/src/dto"
	"amarolio-gateway/src/handlers"
	"amarolio-gateway/src/logger"
	"amarolio-gateway/src/utils"
	"context"
	"errors"
	"sync"
	"time"

	"github.com/gofiber/contrib/v3/websocket"
	"github.com/gofiber/fiber/v3"
	"github.com/valyala/fasthttp"
)

type MessageServiceItf interface {
	GetMessages(userID string, chatroomID string, limit *int, lastID *int) ([]Message, error)
	Subscribe(ctx context.Context, chatroomID string, cc *ChatClient)
	Write(ctx context.Context, cc *ChatClient)
	Read(ctx context.Context, cc *ChatClient, userID string, chatroomID string)
}

type MessageHandlerImpl struct {
	ms MessageServiceItf
	lg logger.Logger
}

func NewMessageHandler(ms MessageServiceItf, lg logger.Logger) *MessageHandlerImpl {
	return &MessageHandlerImpl{ms, lg}
}

func (mh *MessageHandlerImpl) ConnectChatChannel(c *websocket.Conn) {
	claim, ok := c.Locals(constants.AUTH_CLAIM_KEY).(*utils.CustomClaim)
	if !ok {
		c.WriteMessage(websocket.TextMessage, []byte("unauthorized"))
		c.Close()
		return
	}

	chatroomID := c.Params("id")
	c.SetReadDeadline(time.Now().Add(time.Minute))

	ctx, cancel := context.WithCancel(context.Background())

	var wg sync.WaitGroup
	wg.Add(2)

	cc := NewChatClient(c, &wg, cancel, mh.lg)

	mh.ms.Subscribe(ctx, chatroomID, cc)

	go mh.ms.Write(ctx, cc)
	go mh.ms.Read(ctx, cc, claim.Subject, chatroomID)

	wg.Wait()
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
