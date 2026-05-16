package handlers

import (
	"amarolio-auth/src/constants"
	"amarolio-auth/src/customerrors"
	"errors"

	"github.com/gofiber/fiber/v3"
)

func GetAuth(ctx fiber.Ctx) (string, error) {
	auth := ctx.Locals(constants.AUTH_KEY)
	if auth == nil {
		return "", customerrors.NewError(
			"no credential found",
			errors.New("no credential found"),
			customerrors.Unauthenticate,
		)
	}
	authStr, ok := auth.(string)
	if !ok {
		return "", customerrors.NewError(
			"invalid credential",
			errors.New("invalid credential type"),
			customerrors.Unauthenticate,
		)
	}
	return authStr, nil
}
