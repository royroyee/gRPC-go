# Example 3 : Server to server communication (gRPC Server)


> Example 3 deals with the communication between the User service in Example 2 and the Post service (microservice architecture)


### 1. Define new Protobuff service (post.proto)


- ListPostByUserId : rpc to return all posts equivalent to UserId
- ListPosts : rpc to return all posts of exists in post service

Suppose the Post service knows which user id created which posts, but does not know the name of the user id. 
However, by the protobuf definition, the rpcs of the Post service must be sent with the user's name in the field Author.
--> Data must be received through communication with the User service.

```go
syntax = "proto3";

package v1.post;

option go_package = "github.com/royroyee/gRPC-go/ex3/protos/v1/post";

service Post {
  rpc ListPostsByUserId(ListPostsByUserIdRequest) returns (ListPostsByUserIdResponse);
  rpc ListPosts(ListPostsRequest) returns (ListPostsResponse);
}

message PostMessage {
  string post_id = 1;
  string author = 2;
  string title = 3;
  string body = 4;
  repeated string tags = 5;
}

message ListPostsByUserIdRequest {
  string user_id = 1;
}


message ListPostsByUserIdResponse {
  repeated PostMessage post_messages = 1;
}

message ListPostsRequest{}

message ListPostsResponse {
  repeated PostMessage post_messages = 1;
}
```

- flow : bloomrpc <---> Post <---> User

### 2. Define Post Server
```go
package main

import (
	"context"
	client "github.com/royroyee/gRPC-go/ex3"
	postData "github.com/royroyee/gRPC-go/ex3/data/post"
	postpb "github.com/royroyee/gRPC-go/ex3/protos/v1/post"
	userpb "github.com/royroyee/gRPC-go/ex3/protos/v1/user"
	"google.golang.org/grpc"
	"log"
	"net"
)

const portNumber = "9001"

type postServer struct {
	postpb.PostServer

	userCli userpb.UserClient // to use User Service
}

// ListPostsByUserId returns post messages by user_id
func (s *postServer) ListPostsByUserId(ctx context.Context, req *postpb.ListPostsByUserIdRequest) (*postpb.ListPostsByUserIdResponse, error) {
	userID := req.UserId

	resp, err := s.userCli.GetUser(ctx, &userpb.GetUserRequest{UserId: userID})
	if err != nil {
		return nil, err
	}

	var postMessages []*postpb.PostMessage
	for _, up := range postData.UserPosts {
		if up.UserID != userID {
			continue
		}

		for _, p := range up.Posts {
			p.Author = resp.UserMessage.Name
		}

		postMessages = up.Posts
		break
	}

	return &postpb.ListPostsByUserIdResponse{
		PostMessages: postMessages,
	}, nil
}

// ListPosts returns all post messages
func (s *postServer) ListPosts(ctx context.Context, req *postpb.ListPostsRequest) (*postpb.ListPostsResponse, error) {
	var postMessages []*postpb.PostMessage
	for _, up := range postData.UserPosts {
		resp, err := s.userCli.GetUser(ctx, &userpb.GetUserRequest{UserId: up.UserID})
		if err != nil {
			return nil, err
		}

		for _, p := range up.Posts {
			p.Author = resp.UserMessage.Name
		}

		postMessages = append(postMessages, up.Posts...)
	}

	return &postpb.ListPostsResponse{
		PostMessages: postMessages,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":"+portNumber)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	userCli := client.GetUserClient("localhost:9000") // to connect User gRPC server
	grpcServer := grpc.NewServer()
	postpb.RegisterPostServer(grpcServer, &postServer{
		userCli: userCli, // to connect User gRPC server
	})

	log.Printf("start gRPC server on %s port", portNumber)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

```
#### client(GetUserClient)
```go
package ex3

import (
	userpb "github.com/royroyee/gRPC-go/ex3/protos/v1/user"
	"google.golang.org/grpc"
	"sync"
)

var (
	once sync.Once
	cli  userpb.UserClient
)

func GetUserClient(serviceHost string) userpb.UserClient {
	once.Do(func() {
		conn, _ := grpc.Dial(serviceHost,
			grpc.WithInsecure(),
			grpc.WithBlock())

		cli = userpb.NewUserClient(conn)
	})

	return cli
}

```

- sync.Once : to create client only once in the beginning (singleton)
- grpc.Dial() : grpc-go's function, Dial creates a client connection to the given target.
  - WithInsecure() : WithInsecure is an option to disable transport security 
  - WithBlock() : WithBlock is an option to block until a connection is established. If the connection does not need to be established immediately, you can remove this option. Without this option, the default will result in connection in the background.
