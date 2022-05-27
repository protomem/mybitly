package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/protomem/mybitly/internal/controller"
	"github.com/protomem/mybitly/internal/repository"
	"github.com/protomem/mybitly/internal/router"
	"github.com/protomem/mybitly/internal/service"
	"github.com/protomem/mybitly/pkg/httpserver"
	"github.com/protomem/mybitly/pkg/mdb"
	"github.com/sirupsen/logrus"
)

func main() {

	client, err := mdb.NewClient("mongodb://db:27017/")
	if err != nil {
		logrus.Fatal(err)
	}

	defer func() {
		client.Disconnect(context.Background())
	}()

	repositories := repository.New(client)
	services := service.New(repositories)
	controllers := controller.New(services)

	router := router.New(controllers)

	server := httpserver.New(router, 3000)

	go func() {
		if err := server.Run(); err != nil {
			logrus.Error(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	if err := server.ShutDown(context.Background()); err != nil {
		logrus.Error(err)
	}

}
