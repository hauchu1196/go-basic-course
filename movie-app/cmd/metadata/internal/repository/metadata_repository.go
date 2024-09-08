package repository

import "movie-app/cmd/metadata/pkg"

type MetadataRepository interface {
	GetMetadata(movieId string) (models.Metadata, error)
	CreateMetadata(metadata models.Metadata) (models.Metadata, error)
	UpdateMetadata(metadata models.Metadata) (models.Metadata, error)
	DeleteMetadata(movieId string) error
}
