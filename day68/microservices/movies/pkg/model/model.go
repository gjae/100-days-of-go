package model

import (
	model "moviesexample.com/metadata/pkg"
)

// MovieDetails includes movie metadata its aggregated
// rating
type MovieMetadata struct {
	Rating   *float64 `json:"rating,omitEmpty"`
	Metadata model.Metadata
}
