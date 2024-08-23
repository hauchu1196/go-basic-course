package data

import "time"

type Movie struct {
	ID        int64     `json:"id,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	Title     string    `json:"title,omitempty"`
	Year      int64     `json:"year,omitempty"`
	Runtime   Runtime   `json:"runtime,omitempty"`
	Genres    []string  `json:"genres,omitempty"`
	Version   int64     `json:"version,omitempty"`
}
