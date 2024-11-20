package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"learngo/src/protobuf_grpc/proto"
)

func main() {
	client, err := grpc.NewClient("127.0.0.1:8089", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer client.Close()
	greeterClient := proto.NewGreeterClient(client)
	r, err := greeterClient.SayHello(context.Background(), &proto.HelloRequest{Name: "cpq1!~"})
	if err != nil {
		panic(err)
	}
	fmt.Println(r.Message)
}
