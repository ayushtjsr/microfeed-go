package main

import (
	"log"
	"net/http"
	"os"

	"postaggregator/internal/grpcclient"
	"postaggregator/internal/schema"

	"github.com/graphql-go/handler"
)

func main() {
	grpcServerAddr := os.Getenv("GRPC_SERVER_ADDRESS")
	if grpcServerAddr == "" {
		grpcServerAddr = "grpc-server:50051" // Use Docker service name
	}

	grpcClient, err := grpcclient.NewGRPCClient(grpcServerAddr) // use Singleton pattern
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer grpcClient.Close()

	schema, err := schema.SetupSchema(grpcClient)
	if err != nil {
		log.Fatalf("Failed to create GraphQL schema: %v", err)
	}

	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	http.Handle("/graphql", h)
	log.Println("GraphQL server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
