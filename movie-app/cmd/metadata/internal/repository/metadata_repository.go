package repository

import "movie-app/cmd/metadata/internal/models"

type MetadataRepository interface {
	GetMetadata(id string) (models.Metadata, error)
	CreateMetadata(metadata models.Metadata) (models.Metadata, error)
	UpdateMetadata(metadata models.Metadata) (models.Metadata, error)
	DeleteMetadata(id string) error
}
