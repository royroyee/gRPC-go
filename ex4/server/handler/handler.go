package handler

import (
	"context"
	pb "github.com/royroyee/gRPC-go/ex4/api/proto"
	"log"
)

type APIServer struct {
	pb.ApiServer
}

func (s *APIServer) GetHello(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	log.Fatalf("Received: %v", in.GetName())

	return &pb.Response{Message: "Hello " + in.GetName()}, nil
}
