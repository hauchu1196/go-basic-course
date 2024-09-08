package repository

import "movie-app/cmd/movie/internal/models"

type MovieRepository interface {
	GetMovie(id string) (models.Movie, error)
	CreateMovie(movie models.Movie) (models.Movie, error)
	UpdateMovie(movie models.Movie) (models.Movie, error)
	DeleteMovie(id string) error
}
