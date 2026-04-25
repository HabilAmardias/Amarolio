package config

import (
	"amary/src/db"
	"amary/src/domain/url"
	visitrecords "amary/src/domain/visit_records"
	"amary/src/routers"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

func Bootstrap(db *db.DBHandle, app *gin.Engine, rc *redis.Client, lg *zap.SugaredLogger) {
	suc := url.NewShortenURLCache(rc)
	sur := url.NewURLRepo(db)
	vrr := visitrecords.NewVisitRecordRepo(db)

	ue := url.NewURLEncryptor()
	ide := url.NewIDEncoder()

	sus := url.NewURLService(ue, ide, suc, sur, vrr)

	suh := url.NewURLHandler(sus)

	router := &routers.AppRouter{
		App:               app,
		ShortenURLHandler: suh,
		Logger:            lg,
	}

	router.Setup()
}
