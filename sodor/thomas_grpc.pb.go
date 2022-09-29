// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.6
// source: sodor/thomas.proto

package sodor

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

// ThomasClient is the client API for Thomas service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ThomasClient interface {
	HandShake(ctx context.Context, in *ThomasInstance, opts ...grpc.CallOption) (*ThomasReply, error)
	RunTask(ctx context.Context, in *RunTaskRequest, opts ...grpc.CallOption) (*EmptyResponse, error)
}

type thomasClient struct {
	cc grpc.ClientConnInterface
}

func NewThomasClient(cc grpc.ClientConnInterface) ThomasClient {
	return &thomasClient{cc}
}

func (c *thomasClient) HandShake(ctx context.Context, in *ThomasInstance, opts ...grpc.CallOption) (*ThomasReply, error) {
	out := new(ThomasReply)
	err := c.cc.Invoke(ctx, "/Thomas/HandShake", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *thomasClient) RunTask(ctx context.Context, in *RunTaskRequest, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, "/Thomas/RunTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ThomasServer is the server API for Thomas service.
// All implementations must embed UnimplementedThomasServer
// for forward compatibility
type ThomasServer interface {
	HandShake(context.Context, *ThomasInstance) (*ThomasReply, error)
	RunTask(context.Context, *RunTaskRequest) (*EmptyResponse, error)
	mustEmbedUnimplementedThomasServer()
}

// UnimplementedThomasServer must be embedded to have forward compatible implementations.
type UnimplementedThomasServer struct {
}

func (UnimplementedThomasServer) HandShake(context.Context, *ThomasInstance) (*ThomasReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandShake not implemented")
}
func (UnimplementedThomasServer) RunTask(context.Context, *RunTaskRequest) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RunTask not implemented")
}
func (UnimplementedThomasServer) mustEmbedUnimplementedThomasServer() {}

// UnsafeThomasServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ThomasServer will
// result in compilation errors.
type UnsafeThomasServer interface {
	mustEmbedUnimplementedThomasServer()
}

func RegisterThomasServer(s grpc.ServiceRegistrar, srv ThomasServer) {
	s.RegisterService(&Thomas_ServiceDesc, srv)
}

func _Thomas_HandShake_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ThomasInstance)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThomasServer).HandShake(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Thomas/HandShake",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThomasServer).HandShake(ctx, req.(*ThomasInstance))
	}
	return interceptor(ctx, in, info, handler)
}

func _Thomas_RunTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RunTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThomasServer).RunTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Thomas/RunTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThomasServer).RunTask(ctx, req.(*RunTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Thomas_ServiceDesc is the grpc.ServiceDesc for Thomas service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Thomas_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Thomas",
	HandlerType: (*ThomasServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HandShake",
			Handler:    _Thomas_HandShake_Handler,
		},
		{
			MethodName: "RunTask",
			Handler:    _Thomas_RunTask_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sodor/thomas.proto",
}
