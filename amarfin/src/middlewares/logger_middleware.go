package middlewares

import (
	"amarfin/src/constant"
	"html"
	"net/http"
	"strings"
	"time"
)

type InfoLogger interface {
	Infoln(args ...interface{})
}

type LoggerMiddleware struct {
	lg InfoLogger
}

func NewLoggerMiddleware(lg InfoLogger) *LoggerMiddleware {
	return &LoggerMiddleware{lg}
}

func (lm *LoggerMiddleware) Run(f http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		f.ServeHTTP(w, r)

		elapsed := time.Since(start)

		path := html.EscapeString(
			strings.ReplaceAll(
				strings.ReplaceAll(r.URL.Path, "\r", ""), "\n", "",
			),
		)

		statusCode := w.Header().Get(constant.STATUS_CODE)

		lm.lg.Infoln(r.Method, path, statusCode, elapsed.Abs().String())
	})
}
