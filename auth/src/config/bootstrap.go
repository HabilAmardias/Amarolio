package config

import (
	"amarolio-auth/src/db"
	"amarolio-auth/src/routers"
	"amarolio-auth/src/users"

	"github.com/gofiber/fiber/v3"
)

func Bootstrap(db *db.DBHandle, app *fiber.App) {
	ju := users.CreateJWTUtil()
	oau := users.CreateGoogleOauthUtil()

	us := users.NewUserService(oau, ju, db)

	uh := users.NewUserHandler(us)

	r := &routers.AppRouter{
		App: app,
		Uh:  uh,
	}

	r.Setup()
}
