package middlewares

import (
	"html"
	"strings"
	"time"

	"github.com/gofiber/fiber/v3"
)

func NewLoggerMiddleware(lg Logger) fiber.Handler {
	return func(ctx fiber.Ctx) error {
		start := time.Now()

		if err := ctx.Next(); err != nil {
			return err
		}

		elapsed := time.Since(start)
		path := html.EscapeString(string(ctx.Request().URI().Path()))
		path = strings.ReplaceAll(path, "\n", "")
		path = strings.ReplaceAll(path, "\r", "")

		lg.Infoln(ctx.Method(), path, ctx.Response().StatusCode(), elapsed.Abs().String())

		return nil
	}
}
