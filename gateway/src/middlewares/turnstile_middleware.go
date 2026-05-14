package middlewares

import (
	"amarolio-gateway/src/dto"

	"github.com/gofiber/fiber/v3"
)

type TurnstileUtil interface {
	Validate(token string) error
}

func NewTurnstileMiddleware(tu TurnstileUtil) fiber.Handler {
	return func(ctx fiber.Ctx) error {
		req := new(dto.TurnstileReq)
		if err := ctx.Bind().Query(req); err != nil {
			return err
		}
		if err := tu.Validate(req.Token); err != nil {
			return err
		}
		return ctx.Next()
	}
}
