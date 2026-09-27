package routers

import (
	"amarolio-gateway/src/constants"
	"amarolio-gateway/src/domain/chatrooms"
	"amarolio-gateway/src/domain/messages"
	shortenurls "amarolio-gateway/src/domain/shorten_urls"
	"amarolio-gateway/src/domain/users"
	"amarolio-gateway/src/middlewares"
	"amarolio-gateway/src/utils"
	"os"
	"slices"
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/valyala/fasthttp"
)

type Logger interface {
	Infoln(args ...interface{})
	Errorln(args ...interface{})
}

type AppRouter struct {
	App               *fiber.App
	UserHandler       *users.UserHandlerImpl
	MessageHandler    *messages.MessageHandlerImpl
	ChatroomHandler   *chatrooms.ChatroomHandlerImpl
	ShortenURLHandler *shortenurls.ShortenURLHandlerImpl
	JWTUtil           *utils.JWTUtil
	TurnstileUtil     *utils.TurnstileUtil
	Logger            Logger
}

func (ar *AppRouter) Setup() {
	ar.App.Use(cors.New(cors.Config{
		AllowOriginsFunc: func(origin string) bool {
			allowed := os.Getenv("ALLOWED_ORIGIN")
			origins := strings.Split(allowed, ";")
			return slices.Contains(origins, origin)
		},
		AllowCredentials: true,
		AllowMethods:     []string{fasthttp.MethodGet, fasthttp.MethodPost, fasthttp.MethodDelete, fasthttp.MethodPatch, fasthttp.MethodPut},
	}))
	ar.App.Use(middlewares.NewLoggerMiddleware(ar.Logger))
	ar.SetupPublicRoute()
	ar.SetupPrivateRoute()
}

func (ar *AppRouter) SetupPublicRoute() {
	v1 := ar.App.Group("/api/v1")
	v1.Use(middlewares.NewRateLimiterMiddleware(constants.AUTH_CLAIM_KEY))
	v1.Post("/prelogin", ar.UserHandler.PreLogin)
	v1.Post("/verify/send", ar.UserHandler.ResendVerification)
	v1.Post("/logout", ar.UserHandler.LogOut)
	v1.Post("/login", middlewares.NewAuthMiddleware(
		ar.JWTUtil,
		constants.AUTH_TOKEN,
		constants.ForOTP,
		constants.AUTH_CLAIM_KEY,
		false,
	), ar.UserHandler.Login)
	v1.Get("/otp/send", middlewares.NewAuthMiddleware(
		ar.JWTUtil,
		constants.AUTH_TOKEN,
		constants.ForOTP,
		constants.AUTH_CLAIM_KEY,
		false,
	), ar.UserHandler.ResendOTP)
	v1.Post("/register", ar.UserHandler.Register)
	v1.Post("/verify", ar.UserHandler.Verify)
	v1.Post("/refresh", middlewares.NewAuthMiddleware(
		ar.JWTUtil,
		constants.REFRESH_TOKEN,
		constants.ForRefresh,
		constants.REFRESH_CLAIM_KEY,
		false,
	), ar.UserHandler.RefreshAuth)
	v1.Post("/url", middlewares.NewTurnstileMiddleware(ar.TurnstileUtil), middlewares.NewAuthMiddleware(
		ar.JWTUtil,
		constants.AUTH_TOKEN,
		constants.ForAuth,
		constants.AUTH_CLAIM_KEY,
		true,
	), ar.ShortenURLHandler.NewShortURL)
	v1.Get("/url/:id/redirect", middlewares.NewTurnstileMiddleware(ar.TurnstileUtil), ar.ShortenURLHandler.RedirectToURL)
	v1.Get("/url/:id/metadata", ar.ShortenURLHandler.FindOriginalURL)
}

func (ar *AppRouter) SetupPrivateRoute() {
	v1 := ar.App.Group("/api/v1")
	v1.Use(middlewares.NewAuthMiddleware(
		ar.JWTUtil,
		constants.AUTH_TOKEN,
		constants.ForAuth,
		constants.AUTH_CLAIM_KEY,
		false,
	))
	v1.Use(middlewares.NewRateLimiterMiddleware(constants.AUTH_CLAIM_KEY))
	v1.Get("/me", ar.UserHandler.GetProfile)
	v1.Get("/me/url", ar.ShortenURLHandler.GetUserLinks)
	v1.Post("/url/custom-code", ar.ShortenURLHandler.IsCustomURLAvailable)
	v1.Get("/url/:id/dashboard", ar.ShortenURLHandler.GetVisitSummary)
}
