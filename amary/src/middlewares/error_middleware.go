package middlewares

import (
	"amary/src/customerror"
	"amary/src/dto"
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Logger interface {
	Errorln(args ...interface{})
	Infoln(args ...interface{})
}

func NewErrorMiddleware(lg Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
		if len(ctx.Errors) == 0 {
			return
		}

		httpCode := http.StatusInternalServerError
		errorCode := customerror.CommonErr
		var errDetail string = "Internal Server Error"
		err := ctx.Errors[0]

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
			errorCode = customerror.ValidationErr

			lg.Errorln(ctx.Request.Method, ctx.Request.URL.Path, httpCode, errorCode, ve.Error())
			ctx.JSON(httpCode, dto.ServerResponse[dto.ErrorResponse]{
				Success: false,
				Data: dto.ErrorResponse{
					Detail:    errDetail,
					ErrorCode: errorCode,
				},
			})
			return
		}

		var ce *customerror.CustomError
		if errors.As(err, &ce) {
			httpCode = ce.GetErrStatusCode()
			errDetail = ce.UserErr
			errorCode = ce.ErrCode

			lg.Errorln(ctx.Request.Method, ctx.Request.URL.Path, httpCode, errorCode, ce.Error())
			ctx.JSON(httpCode, dto.ServerResponse[dto.ErrorResponse]{
				Success: false,
				Data: dto.ErrorResponse{
					Detail:    errDetail,
					ErrorCode: errorCode,
				},
			})
			return
		}

		lg.Errorln(ctx.Request.Method, ctx.Request.URL.Path, httpCode, errorCode, err.Error())
		ctx.AbortWithStatusJSON(httpCode, dto.ServerResponse[dto.ErrorResponse]{
			Success: false,
			Data: dto.ErrorResponse{
				Detail:    errDetail,
				ErrorCode: errorCode,
			},
		})
	}
}
