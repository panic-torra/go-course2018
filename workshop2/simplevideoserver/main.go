package main

import (
	"context"
	"github.com/panic-torra/go-course2018/workshop2/simplevideoserver/handlers"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

const ServerUrl = ":8000"

func main() {
	initLogger()

	srv := startServer()

	killSignalChan := getKillSignalChan()
	waitForKillSignal(killSignalChan)

	_ = srv.Shutdown(context.Background())
}

func startServer() *http.Server {
	router := handlers.Router()
	srv := &http.Server{Addr: ServerUrl, Handler: router}

	go func() {
		log.Fatal(srv.ListenAndServe())
	}()

	return srv
}

func getKillSignalChan() chan os.Signal {
	osKillSignalChan := make(chan os.Signal, 1)
	signal.Notify(osKillSignalChan, os.Interrupt, syscall.SIGTERM)
	return osKillSignalChan
}

func waitForKillSignal(killSignalChan <-chan os.Signal) {
	killSignal := <-killSignalChan
	switch killSignal {
	case os.Interrupt:
		log.Info("got SIGINT...")
	case syscall.SIGTERM:
		log.Info("got SIGTERM...")
	}
}

func initLogger() {
	log.SetFormatter(&log.JSONFormatter{})
	file, err := os.OpenFile("my.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err == nil {
		log.SetOutput(file)
	}

	log.WithFields(log.Fields{"url": ServerUrl}).Info("starting the server")
}
