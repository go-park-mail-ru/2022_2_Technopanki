// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: session.proto

package handler

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

// AuthCheckerClient is the client API for AuthChecker service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthCheckerClient interface {
	NewSession(ctx context.Context, in *Email, opts ...grpc.CallOption) (*Token, error)
	GetSession(ctx context.Context, in *Token, opts ...grpc.CallOption) (*Email, error)
	DeleteSession(ctx context.Context, in *Token, opts ...grpc.CallOption) (*Nothing, error)
	CreateConfirmationCode(ctx context.Context, in *Email, opts ...grpc.CallOption) (*Token, error)
	GetEmailFromCode(ctx context.Context, in *Token, opts ...grpc.CallOption) (*Email, error)
	GetCodeFromEmail(ctx context.Context, in *Email, opts ...grpc.CallOption) (*Token, error)
}

type authCheckerClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthCheckerClient(cc grpc.ClientConnInterface) AuthCheckerClient {
	return &authCheckerClient{cc}
}

func (c *authCheckerClient) NewSession(ctx context.Context, in *Email, opts ...grpc.CallOption) (*Token, error) {
	out := new(Token)
	err := c.cc.Invoke(ctx, "/session.AuthChecker/NewSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authCheckerClient) GetSession(ctx context.Context, in *Token, opts ...grpc.CallOption) (*Email, error) {
	out := new(Email)
	err := c.cc.Invoke(ctx, "/session.AuthChecker/GetSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authCheckerClient) DeleteSession(ctx context.Context, in *Token, opts ...grpc.CallOption) (*Nothing, error) {
	out := new(Nothing)
	err := c.cc.Invoke(ctx, "/session.AuthChecker/DeleteSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authCheckerClient) CreateConfirmationCode(ctx context.Context, in *Email, opts ...grpc.CallOption) (*Token, error) {
	out := new(Token)
	err := c.cc.Invoke(ctx, "/session.AuthChecker/CreateConfirmationCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authCheckerClient) GetEmailFromCode(ctx context.Context, in *Token, opts ...grpc.CallOption) (*Email, error) {
	out := new(Email)
	err := c.cc.Invoke(ctx, "/session.AuthChecker/GetEmailFromCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authCheckerClient) GetCodeFromEmail(ctx context.Context, in *Email, opts ...grpc.CallOption) (*Token, error) {
	out := new(Token)
	err := c.cc.Invoke(ctx, "/session.AuthChecker/GetCodeFromEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthCheckerServer is the server API for AuthChecker service.
// All implementations must embed UnimplementedAuthCheckerServer
// for forward compatibility
type AuthCheckerServer interface {
	NewSession(context.Context, *Email) (*Token, error)
	GetSession(context.Context, *Token) (*Email, error)
	DeleteSession(context.Context, *Token) (*Nothing, error)
	CreateConfirmationCode(context.Context, *Email) (*Token, error)
	GetEmailFromCode(context.Context, *Token) (*Email, error)
	GetCodeFromEmail(context.Context, *Email) (*Token, error)
	mustEmbedUnimplementedAuthCheckerServer()
}

// UnimplementedAuthCheckerServer must be embedded to have forward compatible implementations.
type UnimplementedAuthCheckerServer struct {
}

func (UnimplementedAuthCheckerServer) NewSession(context.Context, *Email) (*Token, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewSession not implemented")
}
func (UnimplementedAuthCheckerServer) GetSession(context.Context, *Token) (*Email, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSession not implemented")
}
func (UnimplementedAuthCheckerServer) DeleteSession(context.Context, *Token) (*Nothing, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSession not implemented")
}
func (UnimplementedAuthCheckerServer) CreateConfirmationCode(context.Context, *Email) (*Token, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateConfirmationCode not implemented")
}
func (UnimplementedAuthCheckerServer) GetEmailFromCode(context.Context, *Token) (*Email, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEmailFromCode not implemented")
}
func (UnimplementedAuthCheckerServer) GetCodeFromEmail(context.Context, *Email) (*Token, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCodeFromEmail not implemented")
}
func (UnimplementedAuthCheckerServer) mustEmbedUnimplementedAuthCheckerServer() {}

// UnsafeAuthCheckerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthCheckerServer will
// result in compilation errors.
type UnsafeAuthCheckerServer interface {
	mustEmbedUnimplementedAuthCheckerServer()
}

func RegisterAuthCheckerServer(s grpc.ServiceRegistrar, srv AuthCheckerServer) {
	s.RegisterService(&AuthChecker_ServiceDesc, srv)
}

func _AuthChecker_NewSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Email)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthCheckerServer).NewSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/session.AuthChecker/NewSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthCheckerServer).NewSession(ctx, req.(*Email))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthChecker_GetSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Token)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthCheckerServer).GetSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/session.AuthChecker/GetSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthCheckerServer).GetSession(ctx, req.(*Token))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthChecker_DeleteSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Token)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthCheckerServer).DeleteSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/session.AuthChecker/DeleteSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthCheckerServer).DeleteSession(ctx, req.(*Token))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthChecker_CreateConfirmationCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Email)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthCheckerServer).CreateConfirmationCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/session.AuthChecker/CreateConfirmationCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthCheckerServer).CreateConfirmationCode(ctx, req.(*Email))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthChecker_GetEmailFromCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Token)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthCheckerServer).GetEmailFromCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/session.AuthChecker/GetEmailFromCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthCheckerServer).GetEmailFromCode(ctx, req.(*Token))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthChecker_GetCodeFromEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Email)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthCheckerServer).GetCodeFromEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/session.AuthChecker/GetCodeFromEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthCheckerServer).GetCodeFromEmail(ctx, req.(*Email))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthChecker_ServiceDesc is the grpc.ServiceDesc for AuthChecker service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthChecker_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "session.AuthChecker",
	HandlerType: (*AuthCheckerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewSession",
			Handler:    _AuthChecker_NewSession_Handler,
		},
		{
			MethodName: "GetSession",
			Handler:    _AuthChecker_GetSession_Handler,
		},
		{
			MethodName: "DeleteSession",
			Handler:    _AuthChecker_DeleteSession_Handler,
		},
		{
			MethodName: "CreateConfirmationCode",
			Handler:    _AuthChecker_CreateConfirmationCode_Handler,
		},
		{
			MethodName: "GetEmailFromCode",
			Handler:    _AuthChecker_GetEmailFromCode_Handler,
		},
		{
			MethodName: "GetCodeFromEmail",
			Handler:    _AuthChecker_GetCodeFromEmail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "session.proto",
}
