package routers

import (
	"amarolio-gateway/src/chatrooms"
	"amarolio-gateway/src/constants"
	"amarolio-gateway/src/messages"
	"amarolio-gateway/src/middlewares"
	"amarolio-gateway/src/users"
	"amarolio-gateway/src/utils"

	"github.com/gofiber/contrib/v3/websocket"
	"github.com/gofiber/fiber/v3"
)

type AppRouter struct {
	App             *fiber.App
	UserHandler     *users.UserHandlerImpl
	MessageHandler  *messages.MessageHandlerImpl
	ChatroomHandler *chatrooms.ChatroomHandlerImpl
	JWTUtil         *utils.JWTUtil
}

func (ar *AppRouter) Setup() {
	ar.SetupPublicRoute()
	ar.SetupPrivateRoute()
}

func (ar *AppRouter) SetupPublicRoute() {
	v1 := ar.App.Group("/api/v1")
	v1.Post("/login", ar.UserHandler.Login)
	v1.Post(
		"/login/callback",
		middlewares.NewAuthMiddleware(
			ar.JWTUtil,
			constants.REFRESH_TOKEN,
			constants.ForRefresh,
			constants.REFRESH_CLAIM_KEY,
		),
		ar.UserHandler.LoginCallback)
}

func (ar *AppRouter) SetupPrivateRoute() {
	v1 := ar.App.Group("/api/v1")
	v1.Use(middlewares.NewAuthMiddleware(
		ar.JWTUtil,
		constants.AUTH_TOKEN,
		constants.ForAuth,
		constants.AUTH_CLAIM_KEY,
	))

	v1.Post("/refresh", ar.UserHandler.RefreshAuth)
	v1.Get("/chatrooms/me", ar.ChatroomHandler.GetChatrooms)
	v1.Post("/chatrooms", ar.ChatroomHandler.CreateChatroom)
	v1.Delete("/chatrooms/:id", ar.ChatroomHandler.DeleteChatroom)
	v1.Get("/chatrooms/:id/messages", ar.MessageHandler.GetMessages)

	v1.Get("/ws/chatrooms/:id", websocket.New(ar.MessageHandler.ConnectChatChannel))
}
