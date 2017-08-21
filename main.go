package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

const (
	port = ":50051"
)

server := controllers.Setup()
func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// Creates a new gRPC server
	server := grpc.NewServer()
	server.Serve(lis)
	grpclog.Println("Gateway profiling address: ", lis.Addr().String())
}
