package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"learngo/src/stream_grpc/proto"
	"sync"
	"time"
)

func main() {
	client, err := grpc.NewClient("127.0.0.1:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer client.Close()
	greeterClient := proto.NewGreeterClient(client)
	//// 服务端流模式
	//res, _ := greeterClient.GetStream(context.Background(), &proto.StreamReqData{Data: "陈陈"})
	//for {
	//	recv, err := res.Recv()
	//	if err != nil {
	//		fmt.Println(err)
	//		break
	//	}
	//	fmt.Println(recv.Data)
	//}
	//// 客户端流模式
	//stream, _ := greeterClient.PostStream(context.Background())
	//i := 0
	//for {
	//	i++
	//	_ = stream.Send(&proto.StreamReqData{Data: strconv.Itoa(i)})
	//	time.Sleep(time.Second)
	//	if i > 10 {
	//		break
	//	}
	//}
	alls, _ := greeterClient.AllStream(context.Background())
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			recv, _ := alls.Recv()
			fmt.Println("收到服务端消息：" + recv.Data)
		}
	}()

	go func() {
		defer wg.Done()
		for {
			_ = alls.Send(&proto.StreamReqData{Data: "我是客户端，" + time.Now().Format("2006-01-02 15:04:05"+"\n")})
			time.Sleep(time.Second)
		}
	}()
	wg.Wait()
}
