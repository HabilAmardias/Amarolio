package middlewares

import (
	"amarolio-gateway/src/utils"

	"github.com/gofiber/fiber/v3"
)

type JWTUtilItf interface {
	VerifyJWT(tokenStr string, usedFor int) (*utils.CustomClaim, error)
}

// NewIdentityMiddleware resolves the caller's identity for cross-cutting
// concerns (e.g. rate limiting) without enforcing authentication. A missing,
// invalid or expired token is silently treated as anonymous and never blocks
// the request.
func NewIdentityMiddleware(tokenUtil JWTUtilItf, cookieKey string, usedFor int, ctxKey string) fiber.Handler {
	return func(ctx fiber.Ctx) error {
		if val := ctx.Cookies(cookieKey); len(val) > 0 {
			if claim, err := tokenUtil.VerifyJWT(val, usedFor); err == nil && claim != nil {
				ctx.Locals(ctxKey, claim)
			}
		}
		return ctx.Next()
	}
}

func NewAuthMiddleware(tokenUtil JWTUtilItf, cookieKey string, usedFor int, ctxKey string, optional bool) fiber.Handler {
	return func(ctx fiber.Ctx) error {
		val := ctx.Cookies(cookieKey)
		if len(val) == 0 && optional {
			return ctx.Next()
		}
		claim, err := tokenUtil.VerifyJWT(val, usedFor)
		if err != nil {
			return err
		}
		ctx.Locals(ctxKey, claim)
		return ctx.Next()
	}
}
