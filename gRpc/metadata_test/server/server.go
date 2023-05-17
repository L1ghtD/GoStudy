package main

import (
	"context"
	"fmt"
	"goStudy/gRpc/metadata_test/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"net"
)

type Server struct{}

func (s *Server) SayHello(ctx context.Context, req *proto.HelloRequest) (*proto.HelloReply, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		panic("FromIncomingContext error")
	}
	for k, v := range md {
		fmt.Println(k, v)
	}
	/*  输出
	user-agent [grpc-go/1.54.0]
	:authority [127.0.0.1:50054]
	content-type [application/grpc]
	name [ado]
	password [123456]
	*/

	//获取header中的name*********
	//if nameSlice, ok := md["name"]; ok {
	//	fmt.Println(nameSlice) // [ado]
	//	for i, e := range nameSlice {
	//		fmt.Println(i, e) // 0 ado
	//	}
	//}

	return &proto.HelloReply{
		Message: "Hello " + req.Name,
	}, nil
}

func main() {
	g := grpc.NewServer()
	proto.RegisterGreeterServer(g, &Server{})
	lis, err := net.Listen("tcp", "127.0.0.1:50054")
	if err != nil {
		panic("fail to listen: " + err.Error())
	}
	err = g.Serve(lis)
	if err != nil {
		panic("fail to start grpc: " + err.Error())
	}
}
