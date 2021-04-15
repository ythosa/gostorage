package main

import (
	"context"
	"flag"
	"github.com/Ythosa/gostorage/internal/delivery"
	"github.com/Ythosa/gostorage/internal/server"
	"github.com/Ythosa/gostorage/internal/service"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// @title Gostorage
// @version 1.0
// @description Simple storage service for files

// @host localhost:8080

// @contact.name API Support
// @contact.url http://ythosa.github.io

// @license.name GNU GENERAL PUBLIC LICENSE v3
// @license.url https://www.gnu.org/licenses/gpl-3.0.en.html

const defaultGracefulTimeout = time.Second * 15

func main() {
	var wait time.Duration
	flag.DurationVar(
		&wait,
		"graceful-timeout",
		defaultGracefulTimeout,
		"the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m",
	)
	flag.Parse()

	logger := logrus.New()
	services := service.NewService(logger)
	handler := delivery.NewHandler(logger, mux.NewRouter(), services)
	handler.ConfigureRouter()
	srv := server.NewServer(handler, logger)

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), wait)

	if err := srv.Shutdown(ctx); err != nil {
		log.Panic(err)
	}

	cancel()
	os.Exit(0)
}
