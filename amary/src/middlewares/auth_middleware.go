package middlewares

import (
	"amary/src/constant"

	"github.com/gin-gonic/gin"
)

func NewAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		headers := ctx.Request.Header
		userID := headers.Get("x-user-id")
		ctx.Set(constant.AUTH_KEY, userID)

		ctx.Next()
	}
}
