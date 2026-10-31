package config

import (
	dbcommand "amarfin/db/command"
	"amarfin/src/constant"
	"amarfin/src/db"
	"amarfin/src/logger"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/gorilla/mux"
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
	defer dbh.Close()

	if err := dbcommand.Migration(dbh, lg); err != nil {
		lg.Fatalln(err.Error())
	}

	r := mux.NewRouter()

	Bootstrap(r, dbh, lg)

	port := ":" + os.Getenv("SERVER_PORT")
	srv := http.Server{
		Handler: r,
		Addr:    port,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			lg.Fatalln("listen: %s\n", err)
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
