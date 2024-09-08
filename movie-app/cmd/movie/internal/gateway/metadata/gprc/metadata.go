package gprc

import (
	"context"
	"log"
	"movie-app/gen"
	"movie-app/internal/grpcutil"
)

type Gateway struct {
	client gen.MetadataServiceClient
}

func New(client gen.MetadataServiceClient) *Gateway {
	return &Gateway{client: client}
}

func (m *Gateway) Get(ctx context.Context, id string) (*models.Metadata, error) {
	conn, err := grpcutil.ServiceConnection(ctx, "metadata", m.registry)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	client := gen.NewMetadataServiceClient(conn)
	resp, err := client.GetMetadata(ctx, &gen.GetMetadataRequest{MovieID: id})
	if err != nil {
		log.Println("Error getting metadata:", err)
		return nil, err
	}
	return resp.Metadata, nil
}

func (m *Gateway) Create(ctx context.Context, metadata *models.Metadata) error {
	req := &gen.CreateMetadataRequest{Metadata: metadata}
	req := &gen.GetMetadataRequest{MovieID: id}
	resp, err := m.client.GetMetadata(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.Metadata, nil
}

func (m *Gateway) Create(ctx context.Context, metadata *models.Metadata) error {
	req := &gen.CreateMetadataRequest{Metadata: metadata}
	_, err := m.client.CreateMetadata(ctx, req)
	return err
}
