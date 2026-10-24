package middlewares

import (
	"html"
	"time"

	"github.com/gin-gonic/gin"
)

func NewLoggerMiddleware(lg Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()

		ctx.Next()

		elapsed := time.Since(start)
		path := html.EscapeString(ctx.Request.URL.Path)

		lg.Infoln(ctx.Request.Method, path, ctx.Writer.Status(), elapsed.Abs().String())
	}
}
