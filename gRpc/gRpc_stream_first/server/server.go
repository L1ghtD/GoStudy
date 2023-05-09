package main

import (
	"fmt"
	"goStudy/gRpc/gRpc_stream_first/proto"
	"google.golang.org/grpc"
	"net"
	"time"
)

const PORT = ":50052"

type server struct{}

// 以下3个方法实现了GreeterServer接口
// 服务端流模式
func (*server) GetStream(req *proto.StreamReqData, rsp proto.Greeter_GetStreamServer) error {
	i := 0
	for {
		_ = rsp.Send(&proto.StreamRspData{
			Data: fmt.Sprintf("%v", time.Now().Unix()),
		})
		time.Sleep(time.Second)
		i++
		if i > 10 {
			break
		}
	}
	return nil
}

// 客户端流模式
func (*server) PutStream(clientStr proto.Greeter_PutStreamServer) error {
	return nil
}

// 双向流模式
func (*server) AllStream(allStr proto.Greeter_AllStreamServer) error {
	return nil
}

func main() {
	listen, err := net.Listen("tcp", PORT)
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	proto.RegisterGreeterServer(s, &server{})
	err = s.Serve(listen)
	if err != nil {
		panic(err)
	}
}
