package main

import (
	"movie-app/cmd/movie/internal/handler"
	"movie-app/cmd/movie/internal/service"

	metadata_gateway "movie-app/cmd/movie/internal/gateway/metadata/http"
	rating_gateway "movie-app/cmd/movie/internal/gateway/rating/http"

	"github.com/gin-gonic/gin"
)

func main() {
	metadataGateway := metadata_gateway.New("http://localhost:8081")
	ratingGateway := rating_gateway.New("http://localhost:8082")

	service := service.NewMovieService(ratingGateway, metadataGateway)
	handler := handler.NewMovieHandler(service)	

	router := gin.Default()
	router.GET("/movie/:id", handler.Get)

	router.Run(":8083")
}
