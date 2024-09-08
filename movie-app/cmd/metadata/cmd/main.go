package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"movie-app/cmd/metadata/internal/handler"
	"movie-app/cmd/metadata/internal/repository"
	"movie-app/cmd/metadata/internal/service"
	"movie-app/gen"
	"movie-app/pkg/discovery"
	"movie-app/pkg/discovery/consul"
	"net"
	"time"

	"google.golang.org/grpc"
)

const serviceName = "metadata"

func main() {
	var port int
	flag.IntVar(&port, "port", 8084, "port to listen on")
	flag.Parse()
	log.Printf("Starting the metadata service on port %d", port)
	registry, err := consul.NewRegistry("localhost:8500")
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	instanceID := discovery.GenerateInstanceID(serviceName)
	if err := registry.Register(ctx, instanceID, serviceName, fmt.Sprintf("localhost:%d", port)); err != nil {
		log.Fatal(err)
	}

	go func() {
		for {
			if err := registry.ReportHealthyState(instanceID, serviceName); err != nil {
				log.Println("Failed to report healthy state: " + err.Error())
			}
			time.Sleep(1 * time.Second)
		}
	}()
	defer registry.Deregister(ctx, instanceID, serviceName)

	repo := repository.NewMetadataMemoryRepository()
	service := service.NewMetadataService(repo)
	handler := handler.NewMetadataHandler(service)

	srv := grpc.NewServer()
	gen.RegisterMetadataServiceServer(srv, handler)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal(err)
	}
	srv.Serve(lis)
}
