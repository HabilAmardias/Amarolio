package config

import (
	"amarolio-gateway/src/constants"
	"amarolio-gateway/src/db"
	"amarolio-gateway/src/logger"
	"amarolio-gateway/src/middlewares"
	"context"
	"errors"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/valyala/fasthttp"
)

func Run() {
	isProd := os.Getenv("ENVIRONMENT") == constants.PRODUCTION
	lg, err := logger.NewLogger(isProd)
	if err != nil {
		panic(err)
	}
	rc := db.NewRedisClient(
		os.Getenv("AMAROLIO_REDIS_HOST"),
		os.Getenv("REDIS_PORT"),
		os.Getenv("AMAROLIO_REDIS_PASSWORD"),
	)

	app := fiber.New(fiber.Config{
		ErrorHandler: middlewares.NewErrorMiddleware(lg),
	})
	Bootstrap(rc, lg, app)

	server := &fasthttp.Server{
		Handler: app.Handler(),
	}
	go func() {
		if err := server.ListenAndServe(":" + os.Getenv("SERVER_PORT")); !errors.Is(err, fasthttp.ErrConnectionClosed) {
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
