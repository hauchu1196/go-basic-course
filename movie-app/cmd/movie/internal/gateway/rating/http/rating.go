package http

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	models "movie-app/cmd/rating/pkg"
	"movie-app/pkg/discovery"

	"golang.org/x/exp/rand"
)

type Gateway struct {
	registry discovery.Registry
}

func New(registry discovery.Registry) *Gateway {
	return &Gateway{registry: registry}
}

func (g *Gateway) Get(ctx context.Context, movieId string) (float64, error) {
	addresses, err := g.registry.ServiceAddresses(ctx, "rating")
	if err != nil {
		return 0, err
	}
	url := "http://" + addresses[rand.Intn(len(addresses))] + "/ratings"
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("%s/%s", url, movieId), nil)
	if err != nil {
		return 0, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("failed to get rating: %s", resp.Status)
	}

	var rating float64
	err = json.NewDecoder(resp.Body).Decode(&rating)
	if err != nil {
		return 0, err
	}

	return rating, nil
}

func (g *Gateway) Create(ctx context.Context, rating *models.Rating) error {
	addresses, err := g.registry.ServiceAddresses(ctx, "rating")
	if err != nil {
		return err
	}

	url := "http://" + addresses[rand.Intn(len(addresses))] + "/ratings"
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, fmt.Sprintf("%s", url), nil)
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to create rating: %s", resp.Status)
	}

	return nil
}
