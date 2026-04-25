package routers

import (
	"amary/src/domain/url"
	"amary/src/middlewares"

	"github.com/gin-gonic/gin"
)

type Logger interface {
	Infoln(args ...interface{})
	Errorln(args ...interface{})
}

type AppRouter struct {
	App               *gin.Engine
	ShortenURLHandler *url.URLHandlerImpl
	Logger            Logger
}

func (ar *AppRouter) Setup() {
	ar.App.Use(middlewares.NewLoggerMiddleware(ar.Logger))
	ar.App.Use(middlewares.NewErrorMiddleware(ar.Logger))
	ar.SetupPublicRoute()
	ar.SetupPrivateRoute()
}

func (ar *AppRouter) SetupPublicRoute() {
	v1 := ar.App.Group("/api/v1")

	v1.GET("/url/:id", ar.ShortenURLHandler.FindLongURL)
}

func (ar *AppRouter) SetupPrivateRoute() {
	v1 := ar.App.Group("/api/v1")
	v1.Use(middlewares.NewAuthMiddleware())

	v1.GET("/url/me", ar.ShortenURLHandler.GetUserLinks)
	v1.POST("/url", ar.ShortenURLHandler.NewShortURL)
}
