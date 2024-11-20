package main

import (
	"google.golang.org/grpc"
	"learngo/src/stream_grpc/proto"
)

const Port = ":50052"

type Server struct {
	proto.UnimplementedGreeterServer
}

func (s *Server) GetStream(*proto.StreamReqData, grpc.ServerStreamingServer[proto.StreamResData]) error {
	return nil
}

func (s *Server) PostStream(grpc.ClientStreamingServer[proto.StreamReqData, proto.StreamReqData]) error {
	return nil
}

func (s *Server) AllStream(grpc.BidiStreamingServer[proto.StreamReqData, proto.StreamResData]) error {
	return nil
}

func main() {

}
