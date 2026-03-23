package middlewares

import (
	"amarolio-auth/src/customerrors"
	"amarolio-auth/src/dto"
	"errors"
	"net/http"
	"strings"

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
		var errDetail string = "Internal Server Error"

		var ce *customerrors.CustomError
		if errors.As(err, &ce) {
			code = ce.GetErrStatusCode()
			errDetail = ce.UserErr

			logger.Errorln(ctx.Method(), ctx.Path(), code, ce.Error())
		}

		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			fes := []string{}
			for _, fe := range ve {
				de := dto.DetailsError{
					Field:   fe.Field(),
					Message: fe.Error(),
				}
				fes = append(fes, de.ToString())
			}
			code = http.StatusBadRequest
			errDetail = strings.Join(fes, "; ")

			logger.Errorln(ctx.Method(), ctx.Path(), code, ve.Error())
		}

		return ctx.Status(code).JSON(dto.ServerResponse{
			Data: dto.ErrorResponse{
				Detail: errDetail,
			},
		})
	}
}
