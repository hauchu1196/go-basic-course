package service

import (
	"movie-app/cmd/rating/pkg"
	"movie-app/cmd/rating/internal/repository"
)

type RatingService interface {
	GetRating(id string) (*models.Rating, error)
	CreateRating(rating *models.Rating) (*models.Rating, error)
	UpdateRating(rating *models.Rating) (*models.Rating, error)
	DeleteRating(id string) error
}

type ratingService struct {
	repo repository.RatingRepository
}

func NewRatingService(repo repository.RatingRepository) RatingService {
	return &ratingService{repo: repo}
}

func (s *ratingService) GetRating(id string) (*models.Rating, error) {
	return s.repo.GetRating(id)
}

func (s *ratingService) CreateRating(rating *models.Rating) (*models.Rating, error) {
	return s.repo.CreateRating(rating)
}

func (s *ratingService) UpdateRating(rating *models.Rating) (*models.Rating, error) {
	return s.repo.UpdateRating(rating)
}

func (s *ratingService) DeleteRating(id string) error {
	return s.repo.DeleteRating(id)
}
