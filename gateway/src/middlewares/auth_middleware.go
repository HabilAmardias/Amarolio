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

func NewOptionalAuthMiddleware(tokenUtil JWTUtilItf, cookieKey string, usedFor int, ctxKey string) fiber.Handler {
	return func(ctx fiber.Ctx) error {
		val := ctx.Cookies(cookieKey)
		// if the cookies does not exist, just go next
		if len(val) == 0 {
			return ctx.Next()
		}
		// if exist check is the jwt valid or not
		claim, err := tokenUtil.VerifyJWT(val, usedFor)
		if err != nil {
			return err
		}
		ctx.Locals(ctxKey, claim)
		return ctx.Next()
	}
}
