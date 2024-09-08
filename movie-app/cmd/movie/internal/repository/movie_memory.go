package repository

import (
	"errors"
	"movie-app/cmd/movie/internal/models"
	"sync"

	"github.com/google/uuid"
)

type MovieMemoryRepository struct {
	sync.RWMutex
	movies map[string]models.Movie
}

func NewMovieMemoryRepository() *MovieMemoryRepository {
	return &MovieMemoryRepository{
		movies: make(map[string]models.Movie),
	}
}

func (r *MovieMemoryRepository) GetMovie(id string) (models.Movie, error) {
	movie, ok := r.movies[id]
	if !ok {
		return models.Movie{}, errors.New("movie not found")
	}
	return movie, nil
}

func (r *MovieMemoryRepository) CreateMovie(movie models.Movie) (models.Movie, error) {
	r.Lock()
	defer r.Unlock()

	movie.ID = uuid.New().String()
	r.movies[movie.ID] = movie
	return movie, nil
}

func (r *MovieMemoryRepository) UpdateMovie(movie models.Movie) (models.Movie, error) {
	r.Lock()
	defer r.Unlock()

	r.movies[movie.ID] = movie
	return movie, nil
}

func (r *MovieMemoryRepository) DeleteMovie(id string) error {
	r.Lock()
	defer r.Unlock()

	delete(r.movies, id)
	return nil
}
