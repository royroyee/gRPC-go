package main

import (
	"context"
	pb "github.com/royroyee/gRPC-go/ex4/api/proto"
	"google.golang.org/grpc"
	"log"
	"time"
)

const (
	address = "localhost:50051"
	name    = "royroyee"
)

func main() {
	// set up server connection
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Fatalf("did not connect : %v", err)
	}
	defer conn.Close()

	c := pb.NewApiClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	reply, err := c.GetHello(ctx, &pb.Request{Name: name})

	// call GetHello
	if err != nil {
		log.Fatalf("GetHello error : %v", err)
	}
	log.Printf("Person: &v", reply)
}
