package main

import (
	"context"
	"google.golang.org/grpc"
	"learngo/src/protobuf_grpc/proto"
	"net"
)

type Server struct {
	proto.UnimplementedGreeterServer
}

func (s *Server) SayHello(ctx context.Context, request *proto.HelloRequest) (*proto.HelloReply, error) {
	return &proto.HelloReply{Message: "Hello " + request.Name}, nil
}

func main() {
	g := grpc.NewServer()
	proto.RegisterGreeterServer(g, &Server{})
	listen, err := net.Listen("tcp", ":8089")
	if err != nil {
		panic("listen failed:" + err.Error())
	}
	err = g.Serve(listen)
	if err != nil {
		panic("start server failed:" + err.Error())
	}
}
