// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package whatsapp

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// WhatsAppClient is the client API for WhatsApp service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WhatsAppClient interface {
	// SendWhatsApp is gRPC function to send a WhatSapp mail based on the rest API request
	SendWhatsApp(ctx context.Context, in *WARequest, opts ...grpc.CallOption) (*WAResponse, error)
}

type whatsAppClient struct {
	cc grpc.ClientConnInterface
}

func NewWhatsAppClient(cc grpc.ClientConnInterface) WhatsAppClient {
	return &whatsAppClient{cc}
}

func (c *whatsAppClient) SendWhatsApp(ctx context.Context, in *WARequest, opts ...grpc.CallOption) (*WAResponse, error) {
	out := new(WAResponse)
	err := c.cc.Invoke(ctx, "/WhatsApp/SendWhatsApp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WhatsAppServer is the server API for WhatsApp service.
// All implementations must embed UnimplementedWhatsAppServer
// for forward compatibility
type WhatsAppServer interface {
	// SendWhatsApp is gRPC function to send a WhatSapp mail based on the rest API request
	SendWhatsApp(context.Context, *WARequest) (*WAResponse, error)
	mustEmbedUnimplementedWhatsAppServer()
}

// UnimplementedWhatsAppServer must be embedded to have forward compatible implementations.
type UnimplementedWhatsAppServer struct {
}

func (UnimplementedWhatsAppServer) SendWhatsApp(context.Context, *WARequest) (*WAResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendWhatsApp not implemented")
}
func (UnimplementedWhatsAppServer) mustEmbedUnimplementedWhatsAppServer() {}

// UnsafeWhatsAppServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WhatsAppServer will
// result in compilation errors.
type UnsafeWhatsAppServer interface {
	mustEmbedUnimplementedWhatsAppServer()
}

func RegisterWhatsAppServer(s grpc.ServiceRegistrar, srv WhatsAppServer) {
	s.RegisterService(&_WhatsApp_serviceDesc, srv)
}

func _WhatsApp_SendWhatsApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WARequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WhatsAppServer).SendWhatsApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/WhatsApp/SendWhatsApp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WhatsAppServer).SendWhatsApp(ctx, req.(*WARequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _WhatsApp_serviceDesc = grpc.ServiceDesc{
	ServiceName: "WhatsApp",
	HandlerType: (*WhatsAppServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendWhatsApp",
			Handler:    _WhatsApp_SendWhatsApp_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "whatsapp.proto",
}
