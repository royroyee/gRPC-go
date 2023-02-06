package main

import (
	"context"
	"flag"
	v1 "github.com/royroyee/gRPC-go/ex5/proto/v1/info"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
)

var serverAddr = flag.String("server_adddr", "localhost:10000", "The server address with port")

func main() {
	flag.Parse()

	conn, err := grpc.Dial(*serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	// 인증서 없이 -> WithInsecure() -> 최근 버전에서 WithTransportCredentials 이용 권장으로 바뀜

	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}

	defer conn.Close()

	client := v1.NewRouteClient(conn)

	content, err := client.GetInfo(context.Background(), &v1.Content{Message: "Hi GetInfo Unary RPC"})
	// Unary RPC인 GetInfo 에 대한 요청
	if err != nil {
		log.Fatalf("%v", err)
	}
	log.Printf("%s", content)

	stream, err := client.ListInfo(context.Background(), &v1.Content{Message: "Hi ListInfo Server Stream RPC"})
	// Server stream RPC 인 ListInfo 에 대한 요청

	if err != nil {
		log.Fatalf("ListInfo - %v", err)
	}
	for {
		content, err := stream.Recv()
		// Stream 은 Unary 와 달리 Recv() 를 이용해 값을 받는다.
		// For문을 통해 하나씩 풀어본다.
		if err == io.EOF {
			break
		}
		// stream data를 모두 받았는 지 확인하기 위해 io.EOF 이용

		if err != nil {
			log.Fatalf("ListInfo stream - %v", err)
		}
		log.Printf("Content: Message: %s", content.GetMessage())
	}
}
