package repository

import "movie-app/cmd/rating/pkg"

type RatingRepository interface {
	GetRating(id string) (*models.Rating, error)
	CreateRating(rating *models.Rating) (*models.Rating, error)
	UpdateRating(rating *models.Rating) (*models.Rating, error)
	DeleteRating(id string) error
}
