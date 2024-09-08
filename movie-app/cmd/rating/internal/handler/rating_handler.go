package handler

import (
	"movie-app/cmd/rating/pkg"
	"movie-app/cmd/rating/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RatingHandler struct {
	service service.RatingService
}

func NewRatingHandler(service service.RatingService) *RatingHandler {
	return &RatingHandler{service: service}
}

func (h *RatingHandler) GetRating(c *gin.Context) {
	id := c.Param("id")
	rating, err := h.service.GetRating(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, rating)
}

func (h *RatingHandler) CreateRating(c *gin.Context) {
	var rating models.Rating
	if err := c.ShouldBindJSON(&rating); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	createdRating, err := h.service.CreateRating(&rating)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, createdRating)
}

func (h *RatingHandler) UpdateRating(c *gin.Context) {
	var rating models.Rating
	if err := c.ShouldBindJSON(&rating); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedRating, err := h.service.UpdateRating(&rating)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedRating)
}

func (h *RatingHandler) DeleteRating(c *gin.Context) {
	id := c.Param("id")
	err := h.service.DeleteRating(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
