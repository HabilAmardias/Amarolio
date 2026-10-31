package middlewares

import (
	"amarfin/src/constant"
	"amarfin/src/customerror"
	"amarfin/src/dto"
	"amarfin/src/handler"
	"errors"
	"html"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

type ErrorLogger interface {
	Errorln(args ...interface{})
}

type ErrorMiddleware struct {
	lg ErrorLogger
}

func NewErrorMiddleware(lg ErrorLogger) *ErrorMiddleware {
	return &ErrorMiddleware{lg}
}

func (em *ErrorMiddleware) Run(f http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		f.ServeHTTP(w, r)
		val := r.Context().Value(constant.ERROR_CONTEXT_KEY)
		err, ok := val.(error)
		if !ok {
			return
		}

		httpCode := http.StatusInternalServerError
		errCode := customerror.CommonErr

		path := html.EscapeString(
			strings.ReplaceAll(
				strings.ReplaceAll(r.URL.Path, "\r", ""), "\n", "",
			),
		)

		errDetail := "Internal Server Error"

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
			errCode = customerror.ValidationErr

			em.lg.Errorln(r.Method, path, httpCode, errCode, ve.Error())
			handler.WriteJSON(dto.ServerResponse[dto.ErrorResponse]{
				Success: false,
				Data: dto.ErrorResponse{
					Detail:    errDetail,
					ErrorCode: errCode,
				},
			}, httpCode, w)
			return
		}

		var ce *customerror.CustomError
		if errors.As(err, &ce) {
			httpCode = ce.GetErrStatusCode()
			errDetail = ce.UserErr
			errCode = ce.ErrCode

			em.lg.Errorln(r.Method, path, httpCode, errCode, ce.Error())
			handler.WriteJSON(dto.ServerResponse[dto.ErrorResponse]{
				Success: false,
				Data: dto.ErrorResponse{
					Detail:    errDetail,
					ErrorCode: errCode,
				},
			}, httpCode, w)
			return
		}

		em.lg.Errorln(r.Method, path, httpCode, errCode, err.Error())
		handler.WriteJSON(dto.ServerResponse[dto.ErrorResponse]{
			Success: false,
			Data: dto.ErrorResponse{
				Detail:    errDetail,
				ErrorCode: errCode,
			},
		}, httpCode, w)
	})
}
