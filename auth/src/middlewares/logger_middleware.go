package middlewares

import (
	"time"

	"github.com/gofiber/fiber/v3"
)

func NewLoggerMiddleware(logger Logger) fiber.Handler {
	return func(ctx fiber.Ctx) error {
		start := time.Now()

		if err := ctx.Next(); err != nil {
			return err
		}

		elapsed := time.Since(start)
		logger.Infoln(ctx.Method(), ctx.Path(), ctx.Response().StatusCode(), elapsed.Abs().String())

		return nil
	}
}
