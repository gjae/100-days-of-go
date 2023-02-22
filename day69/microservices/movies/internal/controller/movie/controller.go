package movie

import (
	"context"
	"errors"

	metadatamodel "moviesexample.com/metadata/pkg"
	"moviesexample.com/movies/internal/gateway"
	"moviesexample.com/movies/pkg/model"
	ratingmodel "moviesexample.com/rating/pkg"
)

// ErrNotFound is returned when the movie metadata is not
// found
var ErrNotFound = errors.New("Not found")

type ratingGateway interface {
	GetAggregatedRating(ctx context.Context, recordID ratingmodel.RecordID, recordType ratingmodel.RecordType) (float64, error)
	PutRating(ctx context.Context, recordID ratingmodel.RecordID, recordType ratingmodel.RecordType, rating *ratingmodel.Rating) error
}

type metadataGateway interface {
	Get(ctx context.Context, id string) (*metadatamodel.Metadata, error)
}

// Controller defines a movie service controller.
type Controller struct {
	ratingGateway   ratingGateway
	metadataGateway metadataGateway
}

// New creates a new movie service controller

func New(ratingGateway ratingGateway, metadatametadataGateway metadataGateway) *Controller {
	return &Controller{ratingGateway: ratingGateway, metadataGateway: metadatametadataGateway}
}

// Get returns the movie detail including the aggregated
// rating and movie metadata
// Get returns the movie details including the aggregated rating and movie metadata
func (c *Controller) Get(ctx context.Context, id string) (*model.MovieMetadata, error) {
	metadata, err := c.metadataGateway.Get(ctx, id)

	if err != nil && errors.Is(err, gateway.ErrNotFound) {
		return nil, ErrNotFound
	} else if err != nil {
		return nil, err
	}

	details := &model.MovieMetadata{Metadata: *metadata}

	rating, err := c.ratingGateway.GetAggregatedRating(ctx, ratingmodel.RecordID(id), ratingmodel.RecordTypeMovie)
	if err != nil && !errors.Is(err, gateway.ErrNotFound) {
		// Jos proceed in this case, it's ok not to have ratings yet.
	} else if err != nil {
		return nil, err
	} else {
		details.Rating = &rating
	}

	return details, nil
}
