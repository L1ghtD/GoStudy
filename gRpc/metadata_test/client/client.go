package main

import (
	"context"
	"fmt"
	"goStudy/gRpc/grpc_first/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:50054", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := proto.NewGreeterClient(conn)

	// 写入 metadata
	// 也可以使用 metadata.Pairs
	md := metadata.New(map[string]string{
		"name":     "ado",
		"password": "123456",
	})
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	r, err := c.SayHello(ctx, &proto.HelloRequest{
		Name: "world",
	})
	fmt.Println(r.Message)
}
