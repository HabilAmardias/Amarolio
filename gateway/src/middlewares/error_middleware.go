package middlewares

import (
	"amarolio-gateway/src/customerrors"
	"amarolio-gateway/src/dto"
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v3"
)

type Logger interface {
	Errorln(args ...interface{})
	Infoln(args ...interface{})
}

func NewErrorMiddleware(logger Logger) fiber.ErrorHandler {
	return func(ctx fiber.Ctx, err error) error {
		code := http.StatusInternalServerError
		var errDetail string = "Internal Server Error"

		var ce *customerrors.CustomError
		if errors.As(err, &ce) {
			code = ce.GetErrStatusCode()
			errDetail = ce.UserErr

			logger.Errorln(ctx.Method(), ctx.Path(), code, ce.Error())
		}

		return ctx.Status(code).JSON(dto.ServerResponse[dto.ErrorResponse]{
			Data: dto.ErrorResponse{
				Detail: errDetail,
			},
		})
	}
}
