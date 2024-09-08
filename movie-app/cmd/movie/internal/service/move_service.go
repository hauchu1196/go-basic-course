package service

import (
	"context"
	"fmt"

	metadata_models "movie-app/cmd/metadata/pkg"
	"movie-app/cmd/movie/internal/models"
	"movie-app/cmd/movie/internal/repository"
	rating_models "movie-app/cmd/rating/pkg"
)

type ratingGateway interface {
	Create(ctx context.Context, rating *rating_models.Rating) error
	Get(ctx context.Context, movieId string) (float64, error)
}

type metadataGateway interface {
	Get(ctx context.Context, id string) (*metadata_models.Metadata, error)
	Create(ctx context.Context, metadata *metadata_models.Metadata) error
}

type MovieService struct {
	repository      repository.MovieRepository
	ratingGateway   ratingGateway
	metadataGateway metadataGateway
}

func NewMovieService(repository repository.MovieRepository, ratingGateway ratingGateway, metadataGateway metadataGateway) *MovieService {
	return &MovieService{repository: repository, ratingGateway: ratingGateway, metadataGateway: metadataGateway}
}

func (s *MovieService) Create(ctx context.Context, movie *models.Movie) (*models.Movie, error) {
	newMovie, err := s.repository.CreateMovie(*movie)
	if err != nil {
		return nil, err
	}
	movie.ID = newMovie.ID
	movie.Metadata.MovieID = newMovie.ID
	if err := s.metadataGateway.Create(ctx, movie.Metadata); err != nil {
		return nil, err
	}

	return movie, nil
}

func (s *MovieService) Get(ctx context.Context, id string) (*models.Movie, error) {
	rating, _ := s.ratingGateway.Get(ctx, id)

	metadata, err := s.metadataGateway.Get(ctx, id)
	if err != nil {
		fmt.Println("error getting metadata", err)
		return nil, err
	}

	return &models.Movie{
		ID:       id,
		Rating:   &rating,
		Metadata: metadata,
	}, nil
}
