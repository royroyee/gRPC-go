package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	v1 "github.com/royroyee/gRPC-go/ex5/proto/v1/info"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

var (
	port     = flag.Int("port", 10000, "The server port")
	jsonFile = flag.String("json_file", "", "Json file containing list of content")
)

// flag 를 통해 CLI 옵션 정의
// flag 정의에 대한 내용 보기 : go run server.go --help

type RouteServer struct {
	v1.UnimplementedRouteServer
	savedContents []*v1.Content
	// 전송할 내용에 대해 담을 배열 공간이다.
	// loadContents 함수를 통해 json_file 의 내용이 저장된다
}

func (s *RouteServer) loadContents(filePath string) {
	if filePath == "" {
		log.Fatalf("Must set jsonFile option")
	}

	data, err := os.ReadFile(filePath)

	if err != nil {
		log.Fatalf("Failed to load Contents: %v", err)
	}

	if err := json.Unmarshal(data, &s.savedContents); err != nil {
		log.Fatalf("Failed to load : %v", err)
	}
}

// RPC Hanlder

// Unary RPC
func (s *RouteServer) GetInfo(ctx context.Context, req *v1.Content) (*v1.Content, error) {
	log.Printf("GetInfo - %v", req)
	return &v1.Content{Message: "Hi!"}, nil
}

// Server-side Stream RPC
func (s *RouteServer) ListInfo(req *v1.Content, stream v1.Route_ListInfoServer) error {
	log.Printf("ListInfo - %v", req)

	for _, content := range s.savedContents {
		// loadContents를 통해 저장된 데이터를 사용
		if err := stream.Send(content); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))

	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}

	s := &RouteServer{}
	s.loadContents(*jsonFile)

	grpcServer := grpc.NewServer()
	v1.RegisterRouteServer(grpcServer, s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("%v", err)
	}
}
