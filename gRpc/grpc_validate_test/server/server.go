package main

import (
	"context"
	"goStudy/gRpc/grpc_validate_test/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net"
)

type Server struct{}

func (s *Server) SayHello(ctx context.Context, req *proto.Person) (*proto.Person, error) {
	return &proto.Person{
		Id: 23,
	}, nil
}

type Validator interface {
	Validate() error
}

func main() {
	var interceptor grpc.UnaryServerInterceptor
	interceptor = func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		// 继续处理请求
		// 所有 proto 的 message 都实现了 Validate 方法
		if r, ok := req.(Validator); ok {
			if err := r.Validate(); err != nil {
				return nil, status.Error(codes.InvalidArgument, err.Error())
			}
		}

		return handler(ctx, req)
	}

	var opts []grpc.ServerOption
	opts = append(opts, grpc.UnaryInterceptor(interceptor))

	g := grpc.NewServer(opts...)
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
