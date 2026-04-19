package middlewares

import (
	"time"

	"github.com/gin-gonic/gin"
)

func NewLoggerMiddleware(lg Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()

		ctx.Next()

		elapsed := time.Since(start)
		lg.Infoln(ctx.Request.Method, ctx.Request.URL.Path, ctx.Writer.Status(), elapsed.Abs().String())
	}
}
