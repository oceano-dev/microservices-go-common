// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: grpc/email/proto/email.proto

package proto

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

// EmailServiceClient is the client API for EmailService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EmailServiceClient interface {
	SendPasswordCode(ctx context.Context, in *PasswordCodeReq, opts ...grpc.CallOption) (*PasswordCodeRes, error)
	SendSupportMessage(ctx context.Context, in *SupportMessageReq, opts ...grpc.CallOption) (*SupportMessageRes, error)
}

type emailServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEmailServiceClient(cc grpc.ClientConnInterface) EmailServiceClient {
	return &emailServiceClient{cc}
}

func (c *emailServiceClient) SendPasswordCode(ctx context.Context, in *PasswordCodeReq, opts ...grpc.CallOption) (*PasswordCodeRes, error) {
	out := new(PasswordCodeRes)
	err := c.cc.Invoke(ctx, "/src.EmailService/SendPasswordCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *emailServiceClient) SendSupportMessage(ctx context.Context, in *SupportMessageReq, opts ...grpc.CallOption) (*SupportMessageRes, error) {
	out := new(SupportMessageRes)
	err := c.cc.Invoke(ctx, "/src.EmailService/SendSupportMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EmailServiceServer is the server API for EmailService service.
// All implementations must embed UnimplementedEmailServiceServer
// for forward compatibility
type EmailServiceServer interface {
	SendPasswordCode(context.Context, *PasswordCodeReq) (*PasswordCodeRes, error)
	SendSupportMessage(context.Context, *SupportMessageReq) (*SupportMessageRes, error)
	mustEmbedUnimplementedEmailServiceServer()
}

// UnimplementedEmailServiceServer must be embedded to have forward compatible implementations.
type UnimplementedEmailServiceServer struct {
}

func (UnimplementedEmailServiceServer) SendPasswordCode(context.Context, *PasswordCodeReq) (*PasswordCodeRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendPasswordCode not implemented")
}
func (UnimplementedEmailServiceServer) SendSupportMessage(context.Context, *SupportMessageReq) (*SupportMessageRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendSupportMessage not implemented")
}
func (UnimplementedEmailServiceServer) mustEmbedUnimplementedEmailServiceServer() {}

// UnsafeEmailServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EmailServiceServer will
// result in compilation errors.
type UnsafeEmailServiceServer interface {
	mustEmbedUnimplementedEmailServiceServer()
}

func RegisterEmailServiceServer(s grpc.ServiceRegistrar, srv EmailServiceServer) {
	s.RegisterService(&EmailService_ServiceDesc, srv)
}

func _EmailService_SendPasswordCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PasswordCodeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmailServiceServer).SendPasswordCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/src.EmailService/SendPasswordCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmailServiceServer).SendPasswordCode(ctx, req.(*PasswordCodeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmailService_SendSupportMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SupportMessageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmailServiceServer).SendSupportMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/src.EmailService/SendSupportMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmailServiceServer).SendSupportMessage(ctx, req.(*SupportMessageReq))
	}
	return interceptor(ctx, in, info, handler)
}

// EmailService_ServiceDesc is the grpc.ServiceDesc for EmailService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EmailService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "src.EmailService",
	HandlerType: (*EmailServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendPasswordCode",
			Handler:    _EmailService_SendPasswordCode_Handler,
		},
		{
			MethodName: "SendSupportMessage",
			Handler:    _EmailService_SendSupportMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/email/proto/email.proto",
}