package middlewares

import (
	"amarolio-gateway/src/dto"
	"amarolio-gateway/src/utils"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/limiter"
)

const (
	UNAUTHENTICATED_LIMIT = 60
	AUTHENTICATED_LIMIT   = 500
)

func NewRateLimiterMiddleware(ctxKey string) fiber.Handler {
	return limiter.New(
		limiter.Config{
			MaxFunc: func(c fiber.Ctx) int {
				_, ok := c.Locals(ctxKey).(*utils.CustomClaim)
				if !ok {
					return UNAUTHENTICATED_LIMIT
				}
				return AUTHENTICATED_LIMIT
			},
			Expiration: time.Minute,
			KeyGenerator: func(c fiber.Ctx) string {
				claim, ok := c.Locals(ctxKey).(*utils.CustomClaim)
				if !ok {
					return c.IP()
				}
				return claim.Subject
			},
			LimitReached: func(c fiber.Ctx) error {
				return c.Status(http.StatusTooManyRequests).JSON(dto.ServerResponse[dto.ErrorResponse]{
					Success: false,
					Data: dto.ErrorResponse{
						Detail: "Too Many Request, Please Try Again",
					},
				})
			},
		},
	)
}
