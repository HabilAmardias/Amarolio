package middlewares

import (
	"amarolio-auth/src/constants"
	"amarolio-auth/src/customerrors"
	"errors"

	"github.com/gofiber/fiber/v3"
)

func NewAuthMiddleware() fiber.Handler {
	return func(ctx fiber.Ctx) error {
		headers := ctx.Req().GetHeaders()
		userID, ok := headers["x-user-id"]
		if !ok || len(userID) == 0 {
			return customerrors.NewError(
				"unauthorized",
				errors.New("user id not found"),
				customerrors.Unauthenticate,
			)
		}
		ctx.Set(constants.AUTH_KEY, userID[0])
		return ctx.Next()
	}
}
