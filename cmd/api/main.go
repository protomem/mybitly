package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/protomem/mybitly/internal/controller"
	"github.com/protomem/mybitly/internal/repository"
	"github.com/protomem/mybitly/internal/service"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	clientOpts := options.Client().ApplyURI("mongodb://db:27017/")
	client, err := mongo.Connect(context.Background(), clientOpts)
	if err != nil {
		logrus.Fatal(err)
	}

	if err := client.Ping(context.Background(), nil); err != nil {
		logrus.Fatal(err)
	}

	defer func() {
		client.Disconnect(context.Background())
	}()

	repositories := repository.New(client)
	services := service.New(repositories)
	controllers := controller.New(services)

	//TODO Refactor: Split into modules
	router := gin.Default()

	v1 := router.Group("/api/v1")
	controllers.LinkPair.Route("/linkPairs", v1)

	server := &http.Server{
		Addr:         ":3000",
		Handler:      router,
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			logrus.Error(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	if err := server.Shutdown(context.Background()); err != nil {
		logrus.Error(err)
	}

}
