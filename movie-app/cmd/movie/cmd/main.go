package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"movie-app/cmd/movie/internal/handler"
	"movie-app/cmd/movie/internal/repository"
	"movie-app/cmd/movie/internal/service"
	"movie-app/pkg/discovery"
	"movie-app/pkg/discovery/consul"
	"time"

	metadata_gateway "movie-app/cmd/movie/internal/gateway/metadata/http"
	rating_gateway "movie-app/cmd/movie/internal/gateway/rating/http"

	"github.com/gin-gonic/gin"
)

const serviceName = "movie"

func main() {
	var port int
	flag.IntVar(&port, "port", 8081, "port to listen on")
	flag.Parse()
	log.Printf("Starting the movie service on port %d", port)
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

	metadataGateway := metadata_gateway.New(registry)
	ratingGateway := rating_gateway.New(registry)

	service := service.NewMovieService(repository.NewMovieMemoryRepository(), ratingGateway, metadataGateway)
	handler := handler.NewMovieHandler(service)

	router := gin.Default()
	router.GET("/movie/:id", handler.Get)
	router.POST("/movie", handler.Create)

	router.Run(":8083")
}
