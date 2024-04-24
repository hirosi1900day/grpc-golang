package main

import (
	"context"
	"fmt"
	"log"

	"grpc/pb" // Ensure your protobuf package is correctly imported

	"google.golang.org/grpc"
)

// createClient initializes a gRPC connection and returns a FileServiceClient.
func createClient(serverAddress string) (pb.FileServiceClient, func(), error) {
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	if err != nil {
		return nil, nil, fmt.Errorf("failed to dial server: %w", err)
	}

	// Returns the client, a closure for closing the connection, and an error if any
	return pb.NewFileServiceClient(conn), func() { conn.Close() }, nil
}

// callListFiles fetches and prints the list of filenames from the gRPC service.
func callListFiles(client pb.FileServiceClient) error {
	res, err := client.ListFiles(context.Background(), &pb.ListFilesRequest{})
	if err != nil {
		return fmt.Errorf("error calling ListFiles: %w", err)
	}

	fmt.Println("Filenames:", res.GetFilenames())
	return nil
}

func main() {
	// Initialize client
	client, closeConn, err := createClient("localhost:50051")
	if err != nil {
		log.Fatalf("Error creating client: %v", err)
	}
	defer closeConn() // Ensure the connection is closed at the end of the program

	// Fetch and display file list
	if err := callListFiles(client); err != nil {
		log.Fatalf("Failed to list files: %v", err)
	}
}
