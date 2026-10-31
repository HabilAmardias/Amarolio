package middlewares

import (
	"amarfin/src/constant"
	"amarfin/src/handler"
	"context"
	"net/http"

	"github.com/google/uuid"
)

type AuthMiddleware struct{}

func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{}
}

func (am *AuthMiddleware) Run(f http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			uid := r.Header.Get(constant.X_USER_ID)
			_, err := uuid.Parse(uid)
			if err != nil {
				r = handler.Error(err, r)
				return
			}
			r = r.WithContext(context.WithValue(r.Context(), constant.AUTH_KEY, uid))
			f.ServeHTTP(w, r)
		},
	)
}
