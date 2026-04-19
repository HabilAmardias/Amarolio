package config

import (
	dbcommand "amary/db/command"
	"amary/src/constant"
	"amary/src/db"
	"amary/src/logger"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func Run() {
	isProd := os.Getenv("ENVIRONMENT") == constant.PRODUCTION
	lg, err := logger.NewLogger(isProd)
	if err != nil {
		log.Fatalln(err.Error())
	}
	dbh, err := db.ConnectDB(lg, isProd)
	if err != nil {
		lg.Fatalln(err.Error())
	}
	if err := dbcommand.Migration(dbh, lg); err != nil {
		lg.Fatalln(err.Error())
	}

	app := gin.New()
	rc := db.NewRedisClient(
		os.Getenv("AMARY_REDIS_HOST"),
		os.Getenv("REDIS_PORT"),
		os.Getenv("AMARY_REDIS_PASSWORD"),
	)

	Bootstrap(dbh, app, rc, lg)

	defer dbh.Close()

	port := ":" + os.Getenv("SERVER_PORT")
	srv := http.Server{
		Handler: app.Handler(),
		Addr:    port,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	lg.Infoln("Shutdown Server....")
	timeoutEnv := os.Getenv("GRACEFUL_TIMEOUT")
	timeout, err := strconv.Atoi(timeoutEnv)
	if err != nil {
		log.Fatalf("Error: %s\n", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		lg.Infoln("Server Shutdown:", err)
	}

	<-ctx.Done()
	lg.Infof("timeout of %d seconds.\n", timeout)
	lg.Infoln("Server exiting")
}
