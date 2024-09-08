package handler

import (
	"context"
	"movie-app/cmd/metadata/internal/service"
	models "movie-app/cmd/metadata/pkg"
	"movie-app/gen"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MetadataHandler struct {
	gen.UnimplementedMetadataServiceServer
	service service.MetadataService
}

func NewMetadataHandler(service service.MetadataService) MetadataHandler {
	return MetadataHandler{service: service}
}

func (h MetadataHandler) GetMetadata(ctx context.Context, req *gen.GetMetadataRequest) (*gen.GetMetadataResponse, error) {
	metadata, err := h.service.GetMetadata(req.MovieID)
	if err != nil {
		return nil, err
	}
	return &gen.GetMetadataResponse{Metadata: metadata}, nil
}

func (h MetadataHandler) CreateMetadata(c *gin.Context) {
	var metadata models.Metadata
	if err := c.ShouldBindJSON(&metadata); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	metadata, err := h.service.CreateMetadata(metadata)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, metadata)
}

func (h MetadataHandler) UpdateMetadata(c *gin.Context) {
	var metadata models.Metadata
	if err := c.ShouldBindJSON(&metadata); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	metadata, err := h.service.UpdateMetadata(metadata)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, metadata)
}

func (h MetadataHandler) DeleteMetadata(c *gin.Context) {
	movieId := c.Param("movieId")
	if err := h.service.DeleteMetadata(movieId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
