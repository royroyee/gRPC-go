package main

import (
	pb "github.com/royroyee/gRPC-go/ex4/api/proto"
	handler "github.com/royroyee/gRPC-go/ex4/server/handler"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterApiServer(grpcServer, &handler.APIServer{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve %v", err)
	}
}
