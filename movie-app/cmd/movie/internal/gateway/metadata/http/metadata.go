package http

import (
	"context"
	"encoding/json"
	"fmt"
	models "movie-app/cmd/metadata/pkg"
	"net/http"
)

type Gateway struct {
	addr string
}

func New(addr string) *Gateway {
	return &Gateway{addr: addr}
}

func (m *Gateway) Get(ctx context.Context, id string) (*models.Metadata, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("%s/%s", m.addr, id), nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get metadata: %s", resp.Status)
	}

	var metadata models.Metadata
	err = json.NewDecoder(resp.Body).Decode(&metadata)
	if err != nil {
		return nil, err
	}

	return &metadata, nil
}
