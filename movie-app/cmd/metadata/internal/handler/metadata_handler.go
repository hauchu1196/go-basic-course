package handler

import (
	"movie-app/cmd/metadata/internal/models"
	"movie-app/cmd/metadata/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MetadataHandler struct {
	service service.MetadataService
}

func NewMetadataHandler(service service.MetadataService) MetadataHandler {
	return MetadataHandler{service: service}
}

func (h MetadataHandler) GetMetadata(c *gin.Context) {
	id := c.Param("id")
	metadata, err := h.service.GetMetadata(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, metadata)
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
	id := c.Param("id")
	if err := h.service.DeleteMetadata(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
