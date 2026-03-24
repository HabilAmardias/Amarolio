package middlewares

import (
	"amarolio-gateway/src/utils"

	"github.com/gofiber/fiber/v3"
)

type JWTUtilItf interface {
	VerifyJWT(tokenStr string, usedFor int) (*utils.CustomClaim, error)
}

func NewAuthMiddleware(tokenUtil JWTUtilItf, cookieKey string, usedFor int, ctxKey string) fiber.Handler {
	return func(ctx fiber.Ctx) error {
		val := ctx.Cookies(cookieKey)
		claim, err := tokenUtil.VerifyJWT(val, usedFor)
		if err != nil {
			return err
		}
		ctx.Locals(ctxKey, claim)
		return ctx.Next()
	}
}
