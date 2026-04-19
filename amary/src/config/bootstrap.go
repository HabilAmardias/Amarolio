package config

import (
	"amary/src/db"
	shortenurl "amary/src/domain/shorten_url"
	visitrecords "amary/src/domain/visit_records"
	"amary/src/routers"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

func Bootstrap(db *db.DBHandle, app *gin.Engine, rc *redis.Client, lg *zap.SugaredLogger) {
	suc := shortenurl.NewShortenURLCache(rc)
	vrs := visitrecords.NewVisitRecordStream(rc)
	ue := shortenurl.NewURLEncryptor()
	ide := shortenurl.NewIDEncoder()

	sus := shortenurl.NewShortenURLServ(db, ue, ide, vrs, suc)

	suh := shortenurl.NewShortenURLHandler(sus)

	router := &routers.AppRouter{
		App:               app,
		ShortenURLHandler: suh,
		Logger:            lg,
	}

	router.Setup()
}
