package config

import (
	"amarfin/src/db"
	"amarfin/src/modules/transactions"
	"amarfin/src/routers"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func Bootstrap(m *mux.Router, db *db.DBHandle, lg *zap.SugaredLogger) {
	tus := transactions.NewTransactionUseCase(db)
	th := transactions.NewTransactionHandler(tus)

	r := &routers.AppRouter{
		TransactionHandler: th,
		Logger:             lg,
	}
	r.Setup()
}
