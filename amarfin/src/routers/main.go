package routers

import (
	"amarfin/src/middlewares"
	"amarfin/src/modules/transactions"
	"net/http"

	"github.com/gorilla/mux"
)

type Logger interface {
	Infoln(args ...interface{})
	Errorln(args ...interface{})
}

type AppRouter struct {
	App                *mux.Router
	TransactionHandler *transactions.TransactionHandler
	Logger             Logger
}

func (ar *AppRouter) Setup() {
	ar.App.Use(middlewares.NewLoggerMiddleware(ar.Logger).Run)
	ar.App.Use(middlewares.NewErrorMiddleware(ar.Logger).Run)
}

func (ar *AppRouter) SetupPublicRouter() {}

func (ar *AppRouter) SetupPrivateRouter() {
	v1 := ar.App.PathPrefix("/api/v1").Subrouter()
	v1.Use(middlewares.NewAuthMiddleware().Run)

	v1.HandleFunc("/transactions", ar.TransactionHandler.Create).Methods(http.MethodPost)
	v1.HandleFunc("/transactions/{id}", ar.TransactionHandler.Delete).Methods(http.MethodDelete)
}
