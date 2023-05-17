package main

import (
	"context"
	"fmt"
	"goStudy/gRpc/grpc_token_auth_test/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type customCredential struct{}

func (c customCredential) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"appid":  "111222",
		"appkey": "ado",
	}, nil
}

func (c customCredential) RequireTransportSecurity() bool {
	// 不使用安全传输
	return false
}

func main() {
	// 传统写法
	//interceptor := func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	//	startTime := time.Now()
	//
	//	md := metadata.New(map[string]string{
	//		"appid":  "111222",
	//		"appkey": "aado",
	//	})
	//	ctx = metadata.NewOutgoingContext(context.Background(), md)
	//
	//	err := invoker(ctx, method, req, reply, cc, opts...)
	//	fmt.Printf("耗时 %s.\n", time.Since(startTime))
	//	return err
	//}

	//grpc.WithPerRPCCredentials(customCredential{})

	// 二选一
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	opts = append(opts, grpc.WithPerRPCCredentials(customCredential{}))
	conn, err := grpc.Dial("127.0.0.1:50054", opts...)

	//conn, err := grpc.Dial("127.0.0.1:50054", grpc.WithTransportCredentials(insecure.NewCredentials()), opt)
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
