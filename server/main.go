package main

import (
	"context"
	"fmt"
	"grpc/pb"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedFileServiceServer
}

func myLogging() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		fmt.Println("Request: ", req)
		res, err := handler(ctx, req)
		if err != nil {
			return nil, err
		}
		fmt.Println("Response: ", res)
		return res, nil
	}
}

func (*server) ListFiles(ctx context.Context, req *pb.ListFilesRequest) (*pb.ListFilesResponse, error) {
	fmt.Println("ListFiles function was invoked with request: ", req)
	dir := "../storage"

	paths, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var filenames []string
	for _, path := range paths {
		if !path.IsDir() {
			filenames = append(filenames, path.Name())
		}
	}

	res := &pb.ListFilesResponse{
		Filenames: filenames,
	}

	return res, nil
}

func main() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer(
		grpc.UnaryInterceptor(myLogging()),
	)
	pb.RegisterFileServiceServer(s, &server{})

	fmt.Println("Server is running on port: 50051")

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

if _, err osState(path); os.IsNotExist(err) {
		return nil, status.Errorf(codes.NotFound, "file not found: %v", err)
	}

}
