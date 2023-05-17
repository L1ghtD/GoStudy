package main

import (
	"context"
	"fmt"
	"goStudy/gRpc/grpc_interceptor/proto"
	"google.golang.org/grpc"
	"net"
)

type Server struct{}

func (s *Server) SayHello(ctx context.Context, req *proto.HelloRequest) (*proto.HelloReply, error) {
	return &proto.HelloReply{
		Message: "Hello " + req.Name,
	}, nil
}

func main() {
	interceptor := func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		fmt.Println("拦截器触发！")
		// return handler(ctx, req)
		res, err := handler(ctx, req)
		fmt.Println("拦截器结束！")
		return res, err
	}

	opt := grpc.UnaryInterceptor(interceptor)

	g := grpc.NewServer(opt)
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
