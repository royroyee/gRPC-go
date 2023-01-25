# gRPC-go
gRPC example with Golang 

## Reference
- https://tutorialedge.net/golang/go-grpc-beginners-tutorial/  (ex1)
---

## The most basic (refer to ex1 directory)


### Install Protocol Buffer (for mac (local))
```
brew install protobuf
```
to be updated linux (ubuntu)


### Install Plugins
``` go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```
### Define .proto ( ex) chat.proto )
```
syntax = "proto3";
package chat;

message Message {
  string body = 1;
}

service ChatService {
  rpc SayHello(Message) returns (Message) {}
}

```


### Generate the Go specific gRPC code using the protoc tool
```
protoc -I . --go_out=. chat.proto
```
chat.proto must have already been created and this command must be run in a directory that can be found the chat.proto
