# gRPC-go
gRPC example with Golang 

## Reference
- https://grpc.io/docs/languages/go/
- https://github.com/grpc/grpc-go
- https://github.com/dojinkimm/go-grpc-example

---
#### [ex1(simple example)](https://github.com/royroyee/gRPC-go/tree/master/ex1)
#### [ex2(simple tutorial)](https://github.com/royroyee/gRPC-go/tree/master/ex2)
#### [ex3(server to server communiction)](https://github.com/royroyee/gRPC-go/tree/master/ex3)

---

## Tutorial(ex2)


### Install Protocol Buffer (for mac (local))
```
brew install protobuf
```
to be updated linux (ubuntu)


### Install Plugins
``` 
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

> GOPATH, Other settings are not described here

### 1. Define Protobuf Service

- GetUser : GetUser returns user message by user_id
- ListUser : ListUser returns all user Messages

```protobuf
// ex2/protos/v1/user/user.proto

syntax = "proto3";

package v1.user;

option go_package = "github.com/royroyee/gRPC-go/ex2/protos/v1/user";

service User {
  rpc GetUser(GetUserRequest) returns (GetUserResponse);
  rpc ListUser(ListUserRequest) returns (ListUserResponse);
}

message UserMessage {
  string user_id = 1;
  string name = 2;
  string phone_number = 3;
  int32 age = 4;
}

message GetUserRequest {
  string user_id = 1;
}

message GetUserResponse {
  UserMessage user_message = 1;
}

message ListUserRequest{}

message ListUserResponse {
  repeated UserMessage user_messages = 1;
}

```

### 2. Compile protofile
```
protoc -I=. \
	    --go_out . --go_opt paths=source_relative \
	    --go-grpc_out . --go-grpc_opt paths=source_relative \
	    protos/v1/user/user.proto
```
After compilation, the **user.pb.go**, **user_grpc.pb.go** file is created in the directory where the user.proto file is located.

### 3. Implement gRPC server with defined proto(user.pb,proto)

#### 3.1 Database
> If it's real microservice,  get the data from the Database.
   but it's a simple example, so I'll use static variable

```go

// ex2/data/user.go

import userpb "github.com/royroyee/gRPC-go/ex2/protos/v1/user"

var UserData = []*userpb.UserMessage{
    {   
        UserId:      "1",
        Name:        "kiny",
        PhoneNumber: "01012345678",
        Age:         21,
    },

    {
        UserId:      "2",
        Name:        "roy",
        PhoneNumber: "01012345678",
        Age:         24,
    },

    {
        UserId:      "3",
        Name:        "mini",
        PhoneNumber: "01012484428",
        Age:         13,
    },

    {
        UserId:      "4",
        Name:        "jenny",
        PhoneNumber: "01012731533",
        Age:         27,
    },

    {
        UserId:      "5",
        Name:        "jamin",
        PhoneNumber: "01024674568",
        Age:         31,
    },
}

```
total 5 users

#### 3.2 gRPC Server
```go
package main

import (
	"context"
	"github.com/royroyee/gRPC-go/ex2/data"
	userpb "github.com/royroyee/gRPC-go/ex2/protos/v1/user"
	"google.golang.org/grpc"
	"log"
	"net"
)

const portNumber = "9000"

type userServer struct {
	userpb.UserServer
}

// GetUser returns user message by user_id
func (s *userServer) GetUser(ctx context.Context, req *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	userId := req.UserId

	var userMessage *userpb.UserMessage
	for _, u := range data.UserData {
		if u.UserId != userId {
			continue
		}
		userMessage = u
		break
	}
	return &userpb.GetUserResponse{
		UserMessage: userMessage,
	}, nil
}

// ListUsers returns all user Messages
func (s *userServer) ListUser(ctx context.Context, req *userpb.ListUserRequest) (*userpb.ListUserResponse, error) {
	userMessages := make([]*userpb.UserMessage, len(data.UserData))
	for i, u := range data.UserData {
		userMessages[i] = u
	}

	return &userpb.ListUserResponse{
		UserMessages: userMessages,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":"+portNumber)
	if err != nil {
		log.Fatalf("failed to listen : %v", err)
	}

	grpcServer := grpc.NewServer()
	userpb.RegisterUserServer(grpcServer, &userServer{})

	log.Printf("start gRPC server on %s port", portNumber)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve %s", err)
	}
}

```

### 4. Check it out
> I'll use bloomrpc (gRPC gui tool)

> mac(homebrew) : brew install --cask bloomrpc

#### 4.1 Import proto file
-  how to use bloomrpc : https://github.com/bloomrpc/bloomrpc

#### 4.2 gRPC server
```go run main.go```
