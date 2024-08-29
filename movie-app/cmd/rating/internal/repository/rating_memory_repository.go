package repository

import (
	"errors"
	"movie-app/cmd/rating/pkg"
)

type RatingMemoryRepository struct {
	ratings []*models.Rating
}

func NewRatingMemoryRepository() RatingMemoryRepository {
	return RatingMemoryRepository{ratings: []*models.Rating{}}
}

func (r *RatingMemoryRepository) GetRating(id string) (*models.Rating, error) {
	for _, rating := range r.ratings {
		if rating.ID == id {
			return rating, nil
		}
	}
	return &models.Rating{}, errors.New("rating not found")
}

func (r *RatingMemoryRepository) CreateRating(rating *models.Rating) (*models.Rating, error) {
	r.ratings = append(r.ratings, rating)
	return rating, nil
}

func (r *RatingMemoryRepository) UpdateRating(rating *models.Rating) (*models.Rating, error) {
	for idx, item := range r.ratings {
		if item.ID == rating.ID {
			r.ratings[idx].Rating = rating.Rating
			r.ratings[idx].Comment = rating.Comment
			return r.ratings[idx], nil
		}
	}
	return &models.Rating{}, errors.New("rating not found")
}

func (r *RatingMemoryRepository) DeleteRating(id string) error {
	for idx, item := range r.ratings {
		if item.ID == id {
			r.ratings = append(r.ratings[:idx], r.ratings[idx+1:]...)
			return nil
		}
	}
	return errors.New("rating not found")
}
