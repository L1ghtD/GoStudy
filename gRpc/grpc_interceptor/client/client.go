package main

import (
	"context"
	"fmt"
	"goStudy/gRpc/grpc_interceptor/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

func main() {
	interceptor := func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		startTime := time.Now()
		err := invoker(ctx, method, req, reply, cc, opts...)
		fmt.Printf("耗时 %s.\n", time.Since(startTime))
		return err
	}

	// 二选一
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	opts = append(opts, grpc.WithUnaryInterceptor(interceptor))
	conn, err := grpc.Dial("127.0.0.1:50054", opts...)

	//conn, err := grpc.Dial("127.0.0.1:50054", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithUnaryInterceptor(interceptor))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := proto.NewGreeterClient(conn)

	r, err := c.SayHello(context.Background(), &proto.HelloRequest{
		Name: "world",
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(r.Message)
}
