// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package stream

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// GreetereClient is the client API for Greetere service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GreetereClient interface {
	GetStream(ctx context.Context, in *StreamReq, opts ...grpc.CallOption) (Greetere_GetStreamClient, error)
	PutStream(ctx context.Context, opts ...grpc.CallOption) (Greetere_PutStreamClient, error)
	AllStream(ctx context.Context, opts ...grpc.CallOption) (Greetere_AllStreamClient, error)
}

type greetereClient struct {
	cc grpc.ClientConnInterface
}

func NewGreetereClient(cc grpc.ClientConnInterface) GreetereClient {
	return &greetereClient{cc}
}

func (c *greetereClient) GetStream(ctx context.Context, in *StreamReq, opts ...grpc.CallOption) (Greetere_GetStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &Greetere_ServiceDesc.Streams[0], "/Greetere/GetStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &greetereGetStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Greetere_GetStreamClient interface {
	Recv() (*StreamRes, error)
	grpc.ClientStream
}

type greetereGetStreamClient struct {
	grpc.ClientStream
}

func (x *greetereGetStreamClient) Recv() (*StreamRes, error) {
	m := new(StreamRes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *greetereClient) PutStream(ctx context.Context, opts ...grpc.CallOption) (Greetere_PutStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &Greetere_ServiceDesc.Streams[1], "/Greetere/PutStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &greeterePutStreamClient{stream}
	return x, nil
}

type Greetere_PutStreamClient interface {
	Send(*StreamReq) error
	CloseAndRecv() (*StreamRes, error)
	grpc.ClientStream
}

type greeterePutStreamClient struct {
	grpc.ClientStream
}

func (x *greeterePutStreamClient) Send(m *StreamReq) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greeterePutStreamClient) CloseAndRecv() (*StreamRes, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(StreamRes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *greetereClient) AllStream(ctx context.Context, opts ...grpc.CallOption) (Greetere_AllStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &Greetere_ServiceDesc.Streams[2], "/Greetere/AllStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &greetereAllStreamClient{stream}
	return x, nil
}

type Greetere_AllStreamClient interface {
	Send(*StreamReq) error
	Recv() (*StreamRes, error)
	grpc.ClientStream
}

type greetereAllStreamClient struct {
	grpc.ClientStream
}

func (x *greetereAllStreamClient) Send(m *StreamReq) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greetereAllStreamClient) Recv() (*StreamRes, error) {
	m := new(StreamRes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GreetereServer is the server API for Greetere service.
// All implementations must embed UnimplementedGreetereServer
// for forward compatibility
type GreetereServer interface {
	GetStream(*StreamReq, Greetere_GetStreamServer) error
	PutStream(Greetere_PutStreamServer) error
	AllStream(Greetere_AllStreamServer) error
	mustEmbedUnimplementedGreetereServer()
}

// UnimplementedGreetereServer must be embedded to have forward compatible implementations.
type UnimplementedGreetereServer struct {
}

func (UnimplementedGreetereServer) GetStream(*StreamReq, Greetere_GetStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method GetStream not implemented")
}
func (UnimplementedGreetereServer) PutStream(Greetere_PutStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method PutStream not implemented")
}
func (UnimplementedGreetereServer) AllStream(Greetere_AllStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method AllStream not implemented")
}
func (UnimplementedGreetereServer) mustEmbedUnimplementedGreetereServer() {}

// UnsafeGreetereServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GreetereServer will
// result in compilation errors.
type UnsafeGreetereServer interface {
	mustEmbedUnimplementedGreetereServer()
}

func RegisterGreetereServer(s grpc.ServiceRegistrar, srv GreetereServer) {
	s.RegisterService(&Greetere_ServiceDesc, srv)
}

func _Greetere_GetStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GreetereServer).GetStream(m, &greetereGetStreamServer{stream})
}

type Greetere_GetStreamServer interface {
	Send(*StreamRes) error
	grpc.ServerStream
}

type greetereGetStreamServer struct {
	grpc.ServerStream
}

func (x *greetereGetStreamServer) Send(m *StreamRes) error {
	return x.ServerStream.SendMsg(m)
}

func _Greetere_PutStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreetereServer).PutStream(&greeterePutStreamServer{stream})
}

type Greetere_PutStreamServer interface {
	SendAndClose(*StreamRes) error
	Recv() (*StreamReq, error)
	grpc.ServerStream
}

type greeterePutStreamServer struct {
	grpc.ServerStream
}

func (x *greeterePutStreamServer) SendAndClose(m *StreamRes) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greeterePutStreamServer) Recv() (*StreamReq, error) {
	m := new(StreamReq)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Greetere_AllStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreetereServer).AllStream(&greetereAllStreamServer{stream})
}

type Greetere_AllStreamServer interface {
	Send(*StreamRes) error
	Recv() (*StreamReq, error)
	grpc.ServerStream
}

type greetereAllStreamServer struct {
	grpc.ServerStream
}

func (x *greetereAllStreamServer) Send(m *StreamRes) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greetereAllStreamServer) Recv() (*StreamReq, error) {
	m := new(StreamReq)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Greetere_ServiceDesc is the grpc.ServiceDesc for Greetere service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Greetere_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Greetere",
	HandlerType: (*GreetereServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetStream",
			Handler:       _Greetere_GetStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "PutStream",
			Handler:       _Greetere_PutStream_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "AllStream",
			Handler:       _Greetere_AllStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "stream/stream.proto",
}
