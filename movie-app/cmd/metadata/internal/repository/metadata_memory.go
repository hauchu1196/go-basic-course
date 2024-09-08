package repository

import (
	"errors"
	"movie-app/cmd/metadata/pkg"
	"sync"

	"github.com/google/uuid"
)

type MetadataMemoryRepository struct {
	sync.RWMutex
	data map[string]*models.Metadata
}

func NewMetadataMemoryRepository() MetadataRepository {
	return &MetadataMemoryRepository{
		data: make(map[string]*models.Metadata),
	}
}

func (m *MetadataMemoryRepository) GetMetadata(movieId string) (models.Metadata, error) {
	m.RLock()
	defer m.RUnlock()

	metadata, ok := m.data[movieId]
	if !ok {
		return models.Metadata{}, errors.New("metadata not found")
	}

	return *metadata, nil
}

func (m *MetadataMemoryRepository) CreateMetadata(metadata models.Metadata) (models.Metadata, error) {
	m.Lock()
	defer m.Unlock()

	metadata.ID = uuid.New().String()
	m.data[metadata.MovieID] = &metadata

	return metadata, nil
}

func (m *MetadataMemoryRepository) UpdateMetadata(metadata models.Metadata) (models.Metadata, error) {
	m.Lock()
	metadata.ID = uuid.New().String()

	m.data[metadata.MovieID] = &metadata

	return metadata, nil
}

func (m *MetadataMemoryRepository) DeleteMetadata(movieId string) error {
	m.Lock()
	defer m.Unlock()

	delete(m.data, movieId)

	return nil
}
