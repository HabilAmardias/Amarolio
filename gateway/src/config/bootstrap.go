package config

import (
	"amarolio-gateway/src/domain/chatrooms"
	"amarolio-gateway/src/domain/messages"
	shortenurls "amarolio-gateway/src/domain/shorten_urls"
	"amarolio-gateway/src/domain/users"
	"amarolio-gateway/src/logger"
	"amarolio-gateway/src/routers"
	"amarolio-gateway/src/utils"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/redis/go-redis/v9"
)

func Bootstrap(rc *redis.Client, lg logger.Logger, app *fiber.App) {
	ju := utils.NewJWTUtil()

	mcr := messages.NewMessageChannelRepository(rc)

	sus := shortenurls.NewShortenURLService(os.Getenv("AMARY_SERVICE_HOST"), os.Getenv("SERVER_PORT"))
	us := users.NewUserService(os.Getenv("AUTH_SERVICE_HOST"), os.Getenv("SERVER_PORT"))
	ms := messages.NewMessageService(os.Getenv("AMARATH_SERVICE_HOST"), os.Getenv("SERVER_PORT"), mcr)
	cs := chatrooms.NewChatroomService(os.Getenv("AMARATH_SERVICE_HOST"), os.Getenv("SERVER_PORT"))

	ch := chatrooms.NewChatroomHandler(cs)
	mh := messages.NewMessageHandler(ms, lg)
	uh := users.NewUserHandler(us)
	suh := shortenurls.NewShortenURLHandler(sus)

	ar := &routers.AppRouter{
		App:               app,
		JWTUtil:           ju,
		ChatroomHandler:   ch,
		MessageHandler:    mh,
		UserHandler:       uh,
		ShortenURLHandler: suh,
	}
	ar.Setup()
}
