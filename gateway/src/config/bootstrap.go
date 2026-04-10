package config

import (
	"amarolio-gateway/src/chatrooms"
	"amarolio-gateway/src/logger"
	"amarolio-gateway/src/messages"
	"amarolio-gateway/src/routers"
	"amarolio-gateway/src/users"
	"amarolio-gateway/src/utils"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/redis/go-redis/v9"
)

func Bootstrap(rc *redis.Client, lg logger.Logger, app *fiber.App) {
	ju := utils.NewJWTUtil()

	mcr := messages.NewMessageChannelRepository(rc)

	us := users.NewUserService(os.Getenv("USER_SERVICE_HOST"), os.Getenv("SERVER_PORT"))
	ms := messages.NewMessageService(os.Getenv("MESSAGE_SERVICE_HOST"), os.Getenv("SERVER_PORT"), mcr)
	cs := chatrooms.NewChatroomService(os.Getenv("CHATROOM_SERVICE_HOST"), os.Getenv("SERVER_PORT"))

	ch := chatrooms.NewChatroomHandler(cs)
	mh := messages.NewMessageHandler(ms, lg)
	uh := users.NewUserHandler(us)

	ar := &routers.AppRouter{
		App:             app,
		JWTUtil:         ju,
		ChatroomHandler: ch,
		MessageHandler:  mh,
		UserHandler:     uh,
	}
	ar.Setup()
}
