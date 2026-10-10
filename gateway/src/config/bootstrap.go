package config

import (
	shortenurls "amarolio-gateway/src/domain/shorten_urls"
	"amarolio-gateway/src/domain/users"
	"amarolio-gateway/src/logger"
	"amarolio-gateway/src/routers"
	"amarolio-gateway/src/utils"
	"os"

	"github.com/gofiber/fiber/v3"
)

func Bootstrap(lg logger.Logger, app *fiber.App) {
	ju := utils.NewJWTUtil()
	tu := utils.NewTurnstileUtil()

	sus := shortenurls.NewShortenURLService(os.Getenv("AMARY_SERVICE_HOST"), os.Getenv("SERVER_PORT"))
	us := users.NewUserService(os.Getenv("AUTH_SERVICE_HOST"), os.Getenv("SERVER_PORT"))

	uh := users.NewUserHandler(us)
	suh := shortenurls.NewShortenURLHandler(sus)

	ar := &routers.AppRouter{
		App:               app,
		JWTUtil:           ju,
		UserHandler:       uh,
		ShortenURLHandler: suh,
		TurnstileUtil:     tu,
		Logger:            lg,
	}
	ar.Setup()
}
