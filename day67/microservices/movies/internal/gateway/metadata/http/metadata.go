package http

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	model "moviesexample.com/metadata/pkg"
	"moviesexample.com/movies/internal/gateway"
)

// Gateway defines a movie metadata HTTP gateway
type Gateway struct {
	addr string
}

// New creates a new HTTP gateway for a movie metadata
// service
func New(addr string) *Gateway {
	return &Gateway{addr}
}

// Get gets movie metadata by a movie id
func (g *Gateway) Get(ctx context.Context, id string) (*model.Metadata, error) {
	req, err := http.NewRequest(http.MethodGet, g.addr+"/metadata", nil)

	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	values := req.URL.Query()
	values.Add("id", id)
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, gateway.ErrNotFound
	} else if resp.StatusCode/100 != 2 {
		return nil, fmt.Errorf("Non-2xx response: %v", resp)
	}

	var v *model.Metadata

	if err := json.NewDecoder(resp.Body).Decode(&v); err != nil {
		return nil, err
	}

	return v, nil
}
