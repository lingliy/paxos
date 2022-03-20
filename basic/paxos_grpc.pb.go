// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package basic

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

// BasicClient is the client API for Basic service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BasicClient interface {
	Prepare(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	Accept(ctx context.Context, in *Response2, opts ...grpc.CallOption) (*Request, error)
}

type basicClient struct {
	cc grpc.ClientConnInterface
}

func NewBasicClient(cc grpc.ClientConnInterface) BasicClient {
	return &basicClient{cc}
}

func (c *basicClient) Prepare(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/Basic/Prepare", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basicClient) Accept(ctx context.Context, in *Response2, opts ...grpc.CallOption) (*Request, error) {
	out := new(Request)
	err := c.cc.Invoke(ctx, "/Basic/Accept", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BasicServer is the server API for Basic service.
// All implementations must embed UnimplementedBasicServer
// for forward compatibility
type BasicServer interface {
	Prepare(context.Context, *Request) (*Response, error)
	Accept(context.Context, *Response2) (*Request, error)
	mustEmbedUnimplementedBasicServer()
}

// UnimplementedBasicServer must be embedded to have forward compatible implementations.
type UnimplementedBasicServer struct {
}

func (UnimplementedBasicServer) Prepare(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Prepare not implemented")
}
func (UnimplementedBasicServer) Accept(context.Context, *Response2) (*Request, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Accept not implemented")
}
func (UnimplementedBasicServer) mustEmbedUnimplementedBasicServer() {}

// UnsafeBasicServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BasicServer will
// result in compilation errors.
type UnsafeBasicServer interface {
	mustEmbedUnimplementedBasicServer()
}

func RegisterBasicServer(s grpc.ServiceRegistrar, srv BasicServer) {
	s.RegisterService(&Basic_ServiceDesc, srv)
}

func _Basic_Prepare_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BasicServer).Prepare(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Basic/Prepare",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BasicServer).Prepare(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Basic_Accept_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Response2)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BasicServer).Accept(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Basic/Accept",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BasicServer).Accept(ctx, req.(*Response2))
	}
	return interceptor(ctx, in, info, handler)
}

// Basic_ServiceDesc is the grpc.ServiceDesc for Basic service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Basic_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Basic",
	HandlerType: (*BasicServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Prepare",
			Handler:    _Basic_Prepare_Handler,
		},
		{
			MethodName: "Accept",
			Handler:    _Basic_Accept_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "basic/paxos.proto",
}