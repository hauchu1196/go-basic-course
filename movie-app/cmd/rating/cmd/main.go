package main

import (
	"movie-app/cmd/rating/internal/handler"
	"movie-app/cmd/rating/internal/repository"
	"movie-app/cmd/rating/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	repo := repository.NewRatingMemoryRepository()
	service := service.NewRatingService(&repo)
	handler := handler.NewRatingHandler(service)

	router := gin.Default()
	router.GET("/ratings/:id", handler.GetRating)
	router.POST("/ratings", handler.CreateRating)
	router.PUT("/ratings/:id", handler.UpdateRating)
	router.DELETE("/ratings/:id", handler.DeleteRating)

	router.Run(":8082")
}
