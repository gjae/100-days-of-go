package memory

import (
	"context"
	"sync"

	"moviesexample.com/metadata/internal/repository"
	model "moviesexample.com/metadata/pkg"
)

// Repository define a memory movie metadata
type Repository struct {
	sync.RWMutex
	data map[string]*model.Metadata
}

// New creates a new memory repository
func New() *Repository {
	return &Repository{data: map[string]*model.Metadata{}}
}

// Get retrieves movie metadata for by movie id.
func (r *Repository) Get(_ context.Context, id string) (*model.Metadata, error) {
	r.RLock()

	defer r.RUnlock()

	m, ok := r.data[id]

	if !ok {
		return nil, repository.ErrNotFound
	}

	return m, nil
}

// Put adds movie metadata for a given movie id
func (r *Repository) Put(_ context.Context, id string, metadata *model.Metadata) error {
	r.Lock()
	defer r.Unlock()

	r.data[id] = metadata
	return nil
}
