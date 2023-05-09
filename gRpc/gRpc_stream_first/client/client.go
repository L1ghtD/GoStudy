package main

import (
	"context"
	"fmt"
	"goStudy/gRpc/gRpc_stream_first/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	c := proto.NewGreeterClient(conn)
	res, _ := c.GetStream(context.Background(), &proto.StreamReqData{Data: "Server stream"})
	for {
		r, err := res.Recv()
		if err != nil {
			panic(err)
		}
		fmt.Println(r)
	}
}
