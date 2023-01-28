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
