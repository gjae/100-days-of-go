package metadata

import (
	"context"
	"errors"
	"moviesexample.com/metadata/internal/repository"
	"moviesexample.com/metadata/pkg"
)

// ErrNotFound is returned when a requested record  is not found
var ErrNotFound = errors.New("not found")

type metadataRepository interface {
	Get(ctx context.Context, id string) (*model.Metadata, error)
}

// Controller define a metadata service  controller
type Controller struct {
	repo metadataRepository
}

// New creates a metadata service controller
func New(repo metadataRepository) *Controller {
	return &Controller{repo}
}

// Get returns movie metadata by id
func (c *Controller) Get(ctx context.Context, id string) (*model.Metadata, error) {
	res, err := c.repo.Get(ctx, id)
	if err != nil && errors.Is(err, repository.ErrNotFound) {
		return nil, ErrNotFound
	}

	return res, err
}
