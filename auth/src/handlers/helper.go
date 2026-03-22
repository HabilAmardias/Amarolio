package handlers

import (
	"amarolio-auth/src/constants"
	"amarolio-auth/src/customerrors"
	"errors"

	"github.com/gofiber/fiber/v3"
)

func GetAuth(ctx fiber.Ctx) (string, error) {
	auth := ctx.Get(constants.AUTH_KEY)
	if len(auth) == 0 {
		return "", customerrors.NewError(
			"no credential found",
			errors.New("no credential found"),
			customerrors.Unauthenticate,
		)
	}
	return auth, nil
}
