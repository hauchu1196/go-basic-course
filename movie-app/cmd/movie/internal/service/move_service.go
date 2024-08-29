package service

import (
	"context"

	metadata_models "movie-app/cmd/metadata/pkg"
	"movie-app/cmd/movie/internal/models"
	rating_models "movie-app/cmd/rating/pkg"
)

type ratingGateway interface {
	Create(ctx context.Context, rating *rating_models.Rating) error
	Get(ctx context.Context, movieId string) (float64, error)
}

type metadataGateway interface {
	Get(ctx context.Context, id string) (*metadata_models.Metadata, error)
}

type MovieService struct {
	ratingGateway   ratingGateway
	metadataGateway metadataGateway
}

func NewMovieService(ratingGateway ratingGateway, metadataGateway metadataGateway) *MovieService {
	return &MovieService{ratingGateway: ratingGateway, metadataGateway: metadataGateway}
}

func (s *MovieService) Get(ctx context.Context, id string) (*models.Movie, error) {
	rating, err := s.ratingGateway.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	metadata, err := s.metadataGateway.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return &models.Movie{
		ID:       id,
		Rating:   &rating,
		Metadata: metadata,
	}, nil
}
