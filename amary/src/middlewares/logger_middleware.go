package middlewares

import (
	"html"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func NewLoggerMiddleware(lg Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()

		ctx.Next()

		elapsed := time.Since(start)
		path := html.EscapeString(ctx.Request.URL.Path)
		path = strings.ReplaceAll(path, "\n", "")
		path = strings.ReplaceAll(path, "\r", "")

		lg.Infoln(ctx.Request.Method, path, ctx.Writer.Status(), elapsed.Abs().String())
	}
}
