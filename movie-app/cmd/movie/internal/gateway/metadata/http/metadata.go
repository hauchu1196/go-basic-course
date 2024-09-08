package http

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	models "movie-app/cmd/metadata/pkg"
	"movie-app/pkg/discovery"
	"net/http"

	"golang.org/x/exp/rand"
)

type Gateway struct {
	registry discovery.Registry
}

func New(registry discovery.Registry) *Gateway {
	return &Gateway{registry: registry}
}

func (m *Gateway) Get(ctx context.Context, id string) (*models.Metadata, error) {
	addresses, err := m.registry.ServiceAddresses(ctx, "metadata")
	if err != nil {
		return nil, err
	}
	url := "http://" + addresses[rand.Intn(len(addresses))] + "/metadata"
	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("%s/%s", url, id), nil)
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

func (m *Gateway) Create(ctx context.Context, metadata *models.Metadata) error {
	body, err := json.Marshal(metadata)
	if err != nil {
		return err
	}

	addresses, err := m.registry.ServiceAddresses(ctx, "metadata")
	if err != nil {
		return err
	}
	url := "http://" + addresses[rand.Intn(len(addresses))] + "/metadata"
	req, err := http.NewRequestWithContext(ctx, "POST", fmt.Sprintf("%s", url), bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}
