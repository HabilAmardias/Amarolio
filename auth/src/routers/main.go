package routers

import (
	"amarolio-auth/src/domain/users"
	"amarolio-auth/src/middlewares"

	"github.com/gofiber/fiber/v3"
	"go.uber.org/zap"
)

type AppRouter struct {
	App *fiber.App
	Lg  *zap.SugaredLogger
	Uh  *users.UserHandlerImpl
}

func (ar *AppRouter) Setup() {
	ar.App.Use(middlewares.NewLoggerMiddleware(ar.Lg))
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
	v1.Get("/me", ar.Uh.GetProfile)
}
