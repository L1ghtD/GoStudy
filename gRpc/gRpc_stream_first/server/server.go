package main

import (
	"fmt"
	"goStudy/gRpc/gRpc_stream_first/proto"
	"google.golang.org/grpc"
	"net"
	"sync"
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
	for {
		if r, err := clientStr.Recv(); err != nil {
			fmt.Println(err)
			break
		} else {
			fmt.Println(r)
		}
	}
	return nil
}

// 双向流模式
func (*server) AllStream(allStr proto.Greeter_AllStreamServer) error {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			res, _ := allStr.Recv()
			fmt.Println(res)
		}
	}()
	go func() {
		defer wg.Done()
		for {
			_ = allStr.Send(&proto.StreamRspData{
				Data: "我是服务器端",
			})
			time.Sleep(time.Second)
		}
	}()
	wg.Wait()
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
