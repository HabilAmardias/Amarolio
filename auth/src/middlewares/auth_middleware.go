package middlewares

import (
	"amarolio-auth/src/constants"
	"amarolio-auth/src/customerrors"
	"errors"

	"github.com/gofiber/fiber/v3"
)

func NewAuthMiddleware() fiber.Handler {
	return func(ctx fiber.Ctx) error {
		userID := ctx.Get("x-user-id")
		if userID == "" {
			return customerrors.NewError(
				"unauthorized",
				errors.New("user id not found"),
				customerrors.Unauthenticate,
			)
		}
		ctx.Locals(constants.AUTH_KEY, userID)
		return ctx.Next()
	}
}
