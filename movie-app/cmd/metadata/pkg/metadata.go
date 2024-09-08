package models

import (
	"movie-app/gen"
)

type Metadata struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Director    string `json:"director"`
	MovieID     string `json:"movie_id"`
}

func MetadataFromProto(m *Metadata) *gen.Metadata {
	return &gen.Metadata{
		ID:          m.ID,
		Title:       m.Title,
		Description: m.Description,
		Director:    m.Director,
		MovieID:     m.MovieID,
	}
}

func MetadataToProto(m *gen.Metadata) *Metadata {
	return &Metadata{
		ID:          m.ID,
		Title:       m.Title,
		Description: m.Description,
		Director:    m.Director,
		MovieID:     m.MovieID,
	}
}
