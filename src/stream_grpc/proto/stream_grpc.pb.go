// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.0--rc2
// source: stream.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Greeter_GetStream_FullMethodName  = "/Greeter/GetStream"
	Greeter_PostStream_FullMethodName = "/Greeter/PostStream"
	Greeter_AllStream_FullMethodName  = "/Greeter/AllStream"
)

// GreeterClient is the client API for Greeter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GreeterClient interface {
	GetStream(ctx context.Context, in *StreamReqData, opts ...grpc.CallOption) (grpc.ServerStreamingClient[StreamResData], error)
	PostStream(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[StreamReqData, StreamReqData], error)
	AllStream(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[StreamReqData, StreamResData], error)
}

type greeterClient struct {
	cc grpc.ClientConnInterface
}

func NewGreeterClient(cc grpc.ClientConnInterface) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) GetStream(ctx context.Context, in *StreamReqData, opts ...grpc.CallOption) (grpc.ServerStreamingClient[StreamResData], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Greeter_ServiceDesc.Streams[0], Greeter_GetStream_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[StreamReqData, StreamResData]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Greeter_GetStreamClient = grpc.ServerStreamingClient[StreamResData]

func (c *greeterClient) PostStream(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[StreamReqData, StreamReqData], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Greeter_ServiceDesc.Streams[1], Greeter_PostStream_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[StreamReqData, StreamReqData]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Greeter_PostStreamClient = grpc.ClientStreamingClient[StreamReqData, StreamReqData]

func (c *greeterClient) AllStream(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[StreamReqData, StreamResData], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Greeter_ServiceDesc.Streams[2], Greeter_AllStream_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[StreamReqData, StreamResData]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Greeter_AllStreamClient = grpc.BidiStreamingClient[StreamReqData, StreamResData]

// GreeterServer is the server API for Greeter service.
// All implementations must embed UnimplementedGreeterServer
// for forward compatibility.
type GreeterServer interface {
	GetStream(*StreamReqData, grpc.ServerStreamingServer[StreamResData]) error
	PostStream(grpc.ClientStreamingServer[StreamReqData, StreamReqData]) error
	AllStream(grpc.BidiStreamingServer[StreamReqData, StreamResData]) error
	mustEmbedUnimplementedGreeterServer()
}

// UnimplementedGreeterServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedGreeterServer struct{}

func (UnimplementedGreeterServer) GetStream(*StreamReqData, grpc.ServerStreamingServer[StreamResData]) error {
	return status.Errorf(codes.Unimplemented, "method GetStream not implemented")
}
func (UnimplementedGreeterServer) PostStream(grpc.ClientStreamingServer[StreamReqData, StreamReqData]) error {
	return status.Errorf(codes.Unimplemented, "method PostStream not implemented")
}
func (UnimplementedGreeterServer) AllStream(grpc.BidiStreamingServer[StreamReqData, StreamResData]) error {
	return status.Errorf(codes.Unimplemented, "method AllStream not implemented")
}
func (UnimplementedGreeterServer) mustEmbedUnimplementedGreeterServer() {}
func (UnimplementedGreeterServer) testEmbeddedByValue()                 {}

// UnsafeGreeterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GreeterServer will
// result in compilation errors.
type UnsafeGreeterServer interface {
	mustEmbedUnimplementedGreeterServer()
}

func RegisterGreeterServer(s grpc.ServiceRegistrar, srv GreeterServer) {
	// If the following call pancis, it indicates UnimplementedGreeterServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Greeter_ServiceDesc, srv)
}

func _Greeter_GetStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamReqData)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GreeterServer).GetStream(m, &grpc.GenericServerStream[StreamReqData, StreamResData]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Greeter_GetStreamServer = grpc.ServerStreamingServer[StreamResData]

func _Greeter_PostStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreeterServer).PostStream(&grpc.GenericServerStream[StreamReqData, StreamReqData]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Greeter_PostStreamServer = grpc.ClientStreamingServer[StreamReqData, StreamReqData]

func _Greeter_AllStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreeterServer).AllStream(&grpc.GenericServerStream[StreamReqData, StreamResData]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Greeter_AllStreamServer = grpc.BidiStreamingServer[StreamReqData, StreamResData]

// Greeter_ServiceDesc is the grpc.ServiceDesc for Greeter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Greeter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetStream",
			Handler:       _Greeter_GetStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "PostStream",
			Handler:       _Greeter_PostStream_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "AllStream",
			Handler:       _Greeter_AllStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "stream.proto",
}
