package main

import (
	"movie-app/cmd/metadata/internal/handler"
	"movie-app/cmd/metadata/internal/repository"
	"movie-app/cmd/metadata/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	repo := repository.NewMetadataMemoryRepository()
	service := service.NewMetadataService(repo)
	handler := handler.NewMetadataHandler(service)

	router := gin.Default()
	router.GET("/metadata/:id", handler.GetMetadata)
	router.POST("/metadata", handler.CreateMetadata)
	router.PUT("/metadata/:id", handler.UpdateMetadata)
	router.DELETE("/metadata/:id", handler.DeleteMetadata)

	router.Run(":8081")
}
