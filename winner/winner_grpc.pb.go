// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.0
// source: winner.proto

package winner

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

const (
	Winner_Update_FullMethodName = "/Winner/Update"
)

// WinnerClient is the client API for Winner service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WinnerClient interface {
	Update(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
}

type winnerClient struct {
	cc grpc.ClientConnInterface
}

func NewWinnerClient(cc grpc.ClientConnInterface) WinnerClient {
	return &winnerClient{cc}
}

func (c *winnerClient) Update(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, Winner_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WinnerServer is the server API for Winner service.
// All implementations must embed UnimplementedWinnerServer
// for forward compatibility
type WinnerServer interface {
	Update(context.Context, *CreateRequest) (*CreateResponse, error)
	mustEmbedUnimplementedWinnerServer()
}

// UnimplementedWinnerServer must be embedded to have forward compatible implementations.
type UnimplementedWinnerServer struct {
}

func (UnimplementedWinnerServer) Update(context.Context, *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedWinnerServer) mustEmbedUnimplementedWinnerServer() {}

// UnsafeWinnerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WinnerServer will
// result in compilation errors.
type UnsafeWinnerServer interface {
	mustEmbedUnimplementedWinnerServer()
}

func RegisterWinnerServer(s grpc.ServiceRegistrar, srv WinnerServer) {
	s.RegisterService(&Winner_ServiceDesc, srv)
}

func _Winner_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WinnerServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Winner_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WinnerServer).Update(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Winner_ServiceDesc is the grpc.ServiceDesc for Winner service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Winner_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Winner",
	HandlerType: (*WinnerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Update",
			Handler:    _Winner_Update_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "winner.proto",
}
