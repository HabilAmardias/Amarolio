package config

import (
	dbcommand "amary/db/command"
	"amary/src/db"
	"amary/src/domain/url"
	visitrecords "amary/src/domain/visit_records"
	"amary/src/repository"
	"amary/src/routers"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

func Bootstrap(db *db.DBHandle, app *gin.Engine, rc *redis.Client, lg *zap.SugaredLogger) {
	ue := url.NewURLEncryptor()
	ide := url.NewIDEncoder()

	if err := dbcommand.Backfill(db, ide); err != nil {
		log.Fatalf(err.Error())
	}

	suc := url.NewShortenURLCache(rc)
	sur := url.NewURLRepo(db)
	vrr := visitrecords.NewVisitRecordRepo(db)
	trm := repository.NewTransactionManager(db)

	sus := url.NewURLService(ue, ide, suc, sur, vrr, trm)

	suh := url.NewURLHandler(sus)

	router := &routers.AppRouter{
		App:               app,
		ShortenURLHandler: suh,
		Logger:            lg,
	}

	router.Setup()
}
