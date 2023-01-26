# gRPC-go
gRPC example with Golang 

## Reference
- https://alnova2.tistory.com/1373

---
#### [ex1(simple example)](https://github.com/royroyee/gRPC-go/tree/master/ex1)
#### [ex2(simple tutorial)](https://github.com/royroyee/gRPC-go/tree/master/ex2)

---

## Basic(ex1)


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
### Define .proto ( [ex1-helloword.proto](https://github.com/royroyee/gRPC-go/blob/master/ex1/gRPC-server/protos/helloworld/helloworld.proto) )
```

syntax = "proto3";


option go_package="helloWorld/helloworld";
package helloworld;

// The greeting service definition.
service Greeter {
  // Sends a greeting
  rpc SayHello (HelloRequest) returns (HelloReply) {}
}

// The request message containing the user's name.
message HelloRequest {
  string name = 1;
}

// The response message containing the greetings
message HelloReply {
  string message = 1;
}

```


### Generate the Go specific gRPC code using the protoc tool (generate pb.go & _grpc.pb.go)
```
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative helloworld/helloworld.proto
```
helloworld.proto must have already been created and this command must be run in a directory that can be found the helloworld.proto
