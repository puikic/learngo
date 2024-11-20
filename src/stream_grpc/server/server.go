package main

import (
	"fmt"
	"google.golang.org/grpc"
	"learngo/src/stream_grpc/proto"
	"net"
	"sync"
	"time"
)

const Port = ":50052"

type Server struct {
	proto.UnimplementedGreeterServer
}

func (s *Server) GetStream(req *proto.StreamReqData, res grpc.ServerStreamingServer[proto.StreamResData]) error {
	i := 0
	for {
		i++
		_ = res.Send(&proto.StreamResData{Data: req.Data})
		time.Sleep(time.Second)
		if i > 10 {
			break
		}
	}
	return nil
}

func (s *Server) PostStream(cs grpc.ClientStreamingServer[proto.StreamReqData, proto.StreamReqData]) error {
	for {
		recv, err := cs.Recv()
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(recv.Data)
	}
	return nil
}

func (s *Server) AllStream(alls grpc.BidiStreamingServer[proto.StreamReqData, proto.StreamResData]) error {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			recv, _ := alls.Recv()
			fmt.Println("收到客户端消息：" + recv.Data)
		}
	}()

	go func() {
		defer wg.Done()
		for {
			_ = alls.Send(&proto.StreamResData{Data: "我是服务器，" + time.Now().Format("2006-01-02 15:04:05") + "\n"})
			time.Sleep(time.Second)
		}
	}()
	wg.Wait()
	return nil
}

func main() {
	listen, err := net.Listen("tcp", Port)
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	proto.RegisterGreeterServer(s, &Server{})
	err = s.Serve(listen)
	if err != nil {
		panic(err)
	}
}
