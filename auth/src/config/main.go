package config

import (
	dbcommand "amarolio-auth/db/command"
	"amarolio-auth/src/constants"
	"amarolio-auth/src/db"
	"amarolio-auth/src/logger"
	"amarolio-auth/src/middlewares"
	"context"
	"errors"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/valyala/fasthttp"
)

type structValidator struct {
	validate *validator.Validate
}

func (v *structValidator) Validate(out any) error {
	return v.validate.Struct(out)
}

func Run() {
	isProd := os.Getenv("ENVIRONMENT") == constants.PRODUCTION
	lg, err := logger.CreateNewLogger(isProd)
	if err != nil {
		panic(err)
	}

	dbh, err := db.ConnectDB(lg)
	if err != nil {
		lg.Panicln(err)
	}
	if err := dbcommand.Migration(dbh, lg); err != nil {
		lg.Fatalln(err)
	}

	rc := db.NewRedisClient()

	app := fiber.New(fiber.Config{
		ErrorHandler:    middlewares.NewErrorMiddleware(lg),
		StructValidator: &structValidator{validate: validator.New()},
	})
	Bootstrap(dbh, rc, lg, app)

	server := &fasthttp.Server{
		Handler: app.Handler(),
	}
	go func() {
		if err := server.ListenAndServe(":" + os.Getenv("AUTH_SERVER_PORT")); !errors.Is(err, fasthttp.ErrConnectionClosed) {
			lg.Fatalln(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	lg.Infoln("Shutdown Server")
	timeoutEnv := os.Getenv("GRACEFUL_TIMEOUT")
	timeout, err := strconv.Atoi(timeoutEnv)
	if err != nil {
		lg.Fatalf("Error: %s\n", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()
	if err := server.ShutdownWithContext(ctx); err != nil {
		lg.Infoln("Server Shutdown:", err)
	}

	<-ctx.Done()
	lg.Infof("timeout of %d seconds.\n", timeout)
	lg.Infoln("Server exiting")
}
