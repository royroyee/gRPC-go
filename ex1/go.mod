module github.com/royroyee/gRPC-go/ex1

go 1.19

require (
	golang.org/x/net v0.4.0
	google.golang.org/grpc v1.52.1
	google.golang.org/protobuf v1.28.1
)

replace github.com/royroyee/gRPC-go/ex1/gRPC-server/protos/helloworld => ./gRPC-server/protos/helloworld

require (
	github.com/golang/protobuf v1.5.2 // indirect
	golang.org/x/sys v0.3.0 // indirect
	golang.org/x/text v0.5.0 // indirect
	google.golang.org/genproto v0.0.0-20221118155620-16455021b5e6 // indirect
)
