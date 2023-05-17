package main

import (
	"context"
	"fmt"
	"goStudy/gRpc/gRpc_stream_first/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"sync"
	"time"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	c := proto.NewGreeterClient(conn)

	// 服务端流模式
	getStr, _ := c.GetStream(context.Background(), &proto.StreamReqData{Data: "Server stream"})
	for {
		r, err := getStr.Recv()
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(r)
	}

	// 客户端流模式
	putStr, _ := c.PutStream(context.Background())
	i := 0
	for {
		_ = putStr.Send(&proto.StreamReqData{
			Data: fmt.Sprintf("客户流%d", i),
		})
		time.Sleep(time.Second)
		if i > 10 {
			break
		}
		i++
	}

	// 双向流模式
	allStr, _ := c.AllStream(context.Background())
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
			_ = allStr.Send(&proto.StreamReqData{
				Data: "我是客户端",
			})
			time.Sleep(time.Second)
		}
	}()
	wg.Wait()
}
