package config

import (
	"amarolio-auth/src/db"
	"amarolio-auth/src/domain/users"
	"amarolio-auth/src/routers"

	"github.com/gofiber/fiber/v3"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

func Bootstrap(db *db.DBHandle, rc *redis.Client, lg *zap.SugaredLogger, app *fiber.App) {
	ju := users.CreateJWTUtil()
	oau := users.CreateGoogleOauthUtil()

	uc := users.NewUserCache(rc)

	us := users.NewUserService(oau, ju, db, uc, lg)

	uh := users.NewUserHandler(us)

	r := &routers.AppRouter{
		App: app,
		Uh:  uh,
	}

	r.Setup()
}
