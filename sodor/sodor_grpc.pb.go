// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.6
// source: sodor/sodor.proto

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

// FatControllerClient is the client API for FatController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FatControllerClient interface {
	CreateJob(ctx context.Context, in *Job, opts ...grpc.CallOption) (*Reply, error)
}

type fatControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewFatControllerClient(cc grpc.ClientConnInterface) FatControllerClient {
	return &fatControllerClient{cc}
}

func (c *fatControllerClient) CreateJob(ctx context.Context, in *Job, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/FatController/CreateJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FatControllerServer is the server API for FatController service.
// All implementations must embed UnimplementedFatControllerServer
// for forward compatibility
type FatControllerServer interface {
	CreateJob(context.Context, *Job) (*Reply, error)
	mustEmbedUnimplementedFatControllerServer()
}

// UnimplementedFatControllerServer must be embedded to have forward compatible implementations.
type UnimplementedFatControllerServer struct {
}

func (UnimplementedFatControllerServer) CreateJob(context.Context, *Job) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateJob not implemented")
}
func (UnimplementedFatControllerServer) mustEmbedUnimplementedFatControllerServer() {}

// UnsafeFatControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FatControllerServer will
// result in compilation errors.
type UnsafeFatControllerServer interface {
	mustEmbedUnimplementedFatControllerServer()
}

func RegisterFatControllerServer(s grpc.ServiceRegistrar, srv FatControllerServer) {
	s.RegisterService(&FatController_ServiceDesc, srv)
}

func _FatController_CreateJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Job)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FatControllerServer).CreateJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FatController/CreateJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FatControllerServer).CreateJob(ctx, req.(*Job))
	}
	return interceptor(ctx, in, info, handler)
}

// FatController_ServiceDesc is the grpc.ServiceDesc for FatController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FatController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "FatController",
	HandlerType: (*FatControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateJob",
			Handler:    _FatController_CreateJob_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sodor/sodor.proto",
}