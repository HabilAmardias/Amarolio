package middlewares

import (
	"amarolio-auth/src/customerrors"
	"amarolio-auth/src/dto"
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

type Logger interface {
	Errorln(args ...interface{})
	Infoln(args ...interface{})
}

func NewErrorMiddleware(logger Logger) fiber.ErrorHandler {
	return func(ctx fiber.Ctx, err error) error {
		code := http.StatusInternalServerError
		var errDetail any = "Internal Server Error"

		var ce *customerrors.CustomError
		if errors.As(err, &ce) {
			code = ce.GetErrStatusCode()
			errDetail = ce.UserErr

			logger.Errorln(ctx.Method(), ctx.Path(), code, ce.Error())
		}

		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			fes := []dto.DetailsError{}
			for _, fe := range ve {
				fes = append(fes, dto.DetailsError{
					Field:   fe.Field(),
					Message: fe.Error(),
				})
			}
			code = http.StatusBadRequest
			errDetail = fes

			logger.Errorln(ctx.Method(), ctx.Path(), code, ve.Error())
		}

		return ctx.Status(code).JSON(dto.ServerResponse{
			Data: dto.ErrorResponse{
				Detail: errDetail,
			},
		})
	}
}
