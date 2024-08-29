package handler

import (
	"movie-app/cmd/movie/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MovieHandler struct {
	service *service.MovieService
}

func NewMovieHandler(service *service.MovieService) *MovieHandler {
	return &MovieHandler{service: service}
}

func (h *MovieHandler) Get(ctx *gin.Context) {
	id := ctx.Param("id")

	movie, err := h.service.Get(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, movie)
}
