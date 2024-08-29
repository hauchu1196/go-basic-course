package models

import (
	metadata_models "movie-app/cmd/metadata/pkg"
)

type Movie struct {
	ID       string                    `json:"id"`
	Rating   *float64                  `json:"rating"`
	Metadata *metadata_models.Metadata `json:"metadata"`
}
