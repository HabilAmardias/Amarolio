package routers

import (
	"amary/src/domain/url"
	visitrecords "amary/src/domain/visit_records"
	"amary/src/middlewares"

	"github.com/gin-gonic/gin"
)

type Logger interface {
	Infoln(args ...interface{})
	Errorln(args ...interface{})
}

type AppRouter struct {
	App                *gin.Engine
	ShortenURLHandler  *url.URLHandlerImpl
	VisitRecordHandler *visitrecords.VisitRecordHandlerImpl
	Logger             Logger
}

func (ar *AppRouter) Setup() {
	ar.App.Use(middlewares.NewLoggerMiddleware(ar.Logger))
	ar.App.Use(middlewares.NewErrorMiddleware(ar.Logger))
	ar.SetupPublicRoute()
	ar.SetupPrivateRoute()
}

func (ar *AppRouter) SetupPublicRoute() {
	v1 := ar.App.Group("/api/v1")

	v1.GET("/url/:id/redirect", ar.ShortenURLHandler.VisitOriginalURL)
	v1.GET("/url/:id/metadata", ar.ShortenURLHandler.FindOriginalURL)
}

func (ar *AppRouter) SetupPrivateRoute() {
	v1 := ar.App.Group("/api/v1")
	v1.Use(middlewares.NewAuthMiddleware())

	v1.GET("/me/url", ar.ShortenURLHandler.GetUserLinks)
	v1.POST("/url", ar.ShortenURLHandler.NewShortURL)
	v1.POST("/url/custom-code", ar.ShortenURLHandler.IsCustomURLExist)
	v1.GET("/url/:id/dashboard", ar.VisitRecordHandler.GetVisitRecordSummary)
}
