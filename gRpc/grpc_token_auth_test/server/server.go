package main

import (
	"context"
	"fmt"
	"goStudy/gRpc/grpc_token_auth_test/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
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

		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return resp, status.Error(codes.Unauthenticated, "无token认证信息")
		}

		var (
			appid  string
			appkey string
		)

		if v1, ok := md["appid"]; ok {
			appid = v1[0]
		}

		if v2, ok := md["appkey"]; ok {
			appkey = v2[0]
		}

		if appid != "111222" || appkey != "ado" {
			return resp, status.Error(codes.Unauthenticated, "无token认证信息")
		}

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
