package handlers

import (
	"amarolio-gateway/src/customerrors"
	"amarolio-gateway/src/utils"
	"errors"

	"github.com/gofiber/fiber/v3"
)

func GetAuthPayload(ctx fiber.Ctx, key string) (*utils.CustomClaim, error) {
	claim, ok := ctx.Locals(key).(*utils.CustomClaim)
	if !ok {
		return nil, customerrors.NewError(
			"credentials does not found",
			errors.New("jwt claim is missing"),
			customerrors.Unauthenticate,
		)
	}
	return claim, nil
}
