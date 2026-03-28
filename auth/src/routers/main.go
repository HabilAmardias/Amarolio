package routers

import (
	"amarolio-auth/src/middlewares"
	"amarolio-auth/src/users"

	"github.com/gofiber/fiber/v3"
)

type AppRouter struct {
	App *fiber.App
	Uh  *users.UserHandlerImpl
}

func (ar *AppRouter) Setup() {
	ar.publicSetup()
	ar.privateSetup()
}

func (ar *AppRouter) publicSetup() {
	v1 := ar.App.Group("/api/v1")
	v1.Post("/login", ar.Uh.Login)
	v1.Post("/login/callback", ar.Uh.LoginCallback)
}

func (ar *AppRouter) privateSetup() {
	v1 := ar.App.Group("/api/v1")
	v1.Use(middlewares.NewAuthMiddleware())
	v1.Post("/refresh", ar.Uh.RefreshAuth)
}
