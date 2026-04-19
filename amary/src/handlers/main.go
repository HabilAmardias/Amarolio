package handlers

import (
	"amary/src/constant"

	"github.com/gin-gonic/gin"
)

func GetAuthenticationPayload(ctx *gin.Context) string {
	val, exist := ctx.Get(constant.AUTH_KEY)
	if !exist {
		return ""
	}
	userID, ok := val.(string)
	if !ok {
		return ""
	}
	return userID
}
