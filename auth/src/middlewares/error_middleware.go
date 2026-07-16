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
		httpCode := http.StatusInternalServerError
		errorCode := customerrors.CommonErr
		var errDetail string = "Internal Server Error"

		var ce *customerrors.CustomError
		if errors.As(err, &ce) {
			httpCode = ce.GetErrStatusCode()
			errDetail = ce.UserErr
			errorCode = ce.ErrCode

			logger.Errorln(ctx.Method(), ctx.Path(), httpCode, errorCode, ce.Error())
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
			httpCode = http.StatusBadRequest
			errDetail = strings.Join(fes, "; ")
			errorCode = customerrors.ValidationErr

			logger.Errorln(ctx.Method(), ctx.Path(), httpCode, errorCode, ve.Error())
		}

		return ctx.Status(httpCode).JSON(dto.ServerResponse{
			Data: dto.ErrorResponse{
				Detail:    errDetail,
				ErrorCode: errorCode,
			},
		})
	}
}
