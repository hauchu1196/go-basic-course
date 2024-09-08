package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"movie-app/cmd/rating/internal/handler"
	"movie-app/cmd/rating/internal/repository"
	"movie-app/cmd/rating/internal/service"
	"movie-app/pkg/discovery"
	"movie-app/pkg/discovery/consul"
	"time"

	"github.com/gin-gonic/gin"
)

const serviceName = "rating"

func main() {
	var port int
	flag.IntVar(&port, "port", 8082, "port to listen on")
	flag.Parse()
	log.Printf("Starting the rating service on port %d", port)
	registry, err := consul.NewRegistry("localhost:8500")
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	instanceID := discovery.GenerateInstanceID(serviceName)
	if err := registry.Register(ctx, instanceID, serviceName, fmt.Sprintf("localhost:%d", port)); err != nil {
		log.Fatal(err)
	}

	go func() {
		for {
			if err := registry.ReportHealthyState(instanceID, serviceName); err != nil {
				log.Println("Failed to report healthy state: " + err.Error())
			}
			time.Sleep(1 * time.Second)
		}
	}()
	defer registry.Deregister(ctx, instanceID, serviceName)

	repo := repository.NewRatingMemoryRepository()
	service := service.NewRatingService(&repo)
	handler := handler.NewRatingHandler(service)

	router := gin.Default()
	router.GET("/ratings/:id", handler.GetRating)
	router.POST("/ratings", handler.CreateRating)
	router.PUT("/ratings/:id", handler.UpdateRating)
	router.DELETE("/ratings/:id", handler.DeleteRating)

	router.Run(fmt.Sprintf(":%d", port))
}
