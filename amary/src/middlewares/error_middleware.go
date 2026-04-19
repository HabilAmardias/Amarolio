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

		code := http.StatusInternalServerError
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
			code = http.StatusBadRequest
			errDetail = strings.Join(fes, "; ")

			lg.Errorln(ctx.Request.Method, ctx.Request.URL.Path, code, ve.Error())
			ctx.JSON(code, dto.ServerResponse[dto.ErrorResponse]{
				Success: false,
				Data: dto.ErrorResponse{
					Detail: errDetail,
				},
			})
			return
		}

		var ce *customerror.CustomError
		if errors.As(err, &ce) {
			code = ce.GetErrStatusCode()
			errDetail = ce.UserErr

			lg.Errorln(ctx.Request.Method, ctx.Request.URL.Path, code, ce.Error())
			ctx.JSON(code, dto.ServerResponse[dto.ErrorResponse]{
				Success: false,
				Data: dto.ErrorResponse{
					Detail: errDetail,
				},
			})
			return
		}

		lg.Errorln(ctx.Request.Method, ctx.Request.URL.Path, code, err.Error())
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, dto.ServerResponse[dto.ErrorResponse]{
			Success: false,
			Data: dto.ErrorResponse{
				Detail: errDetail,
			},
		})
	}
}
