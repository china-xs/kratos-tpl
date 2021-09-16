// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package account

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

// AccountClient is the client API for Account service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccountClient interface {
	CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*CreateAccountReply, error)
	UpdateAccount(ctx context.Context, in *UpdateAccountRequest, opts ...grpc.CallOption) (*UpdateAccountReply, error)
	DeleteAccount(ctx context.Context, in *DeleteAccountRequest, opts ...grpc.CallOption) (*DeleteAccountReply, error)
	GetAccount(ctx context.Context, in *GetAccountRequest, opts ...grpc.CallOption) (*GetAccountReply, error)
	ListAccount(ctx context.Context, in *ListAccountRequest, opts ...grpc.CallOption) (*ListAccountReply, error)
}

type accountClient struct {
	cc grpc.ClientConnInterface
}

func NewAccountClient(cc grpc.ClientConnInterface) AccountClient {
	return &accountClient{cc}
}

func (c *accountClient) CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*CreateAccountReply, error) {
	out := new(CreateAccountReply)
	err := c.cc.Invoke(ctx, "/api.auth.v1.account.Account/CreateAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) UpdateAccount(ctx context.Context, in *UpdateAccountRequest, opts ...grpc.CallOption) (*UpdateAccountReply, error) {
	out := new(UpdateAccountReply)
	err := c.cc.Invoke(ctx, "/api.auth.v1.account.Account/UpdateAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) DeleteAccount(ctx context.Context, in *DeleteAccountRequest, opts ...grpc.CallOption) (*DeleteAccountReply, error) {
	out := new(DeleteAccountReply)
	err := c.cc.Invoke(ctx, "/api.auth.v1.account.Account/DeleteAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) GetAccount(ctx context.Context, in *GetAccountRequest, opts ...grpc.CallOption) (*GetAccountReply, error) {
	out := new(GetAccountReply)
	err := c.cc.Invoke(ctx, "/api.auth.v1.account.Account/GetAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) ListAccount(ctx context.Context, in *ListAccountRequest, opts ...grpc.CallOption) (*ListAccountReply, error) {
	out := new(ListAccountReply)
	err := c.cc.Invoke(ctx, "/api.auth.v1.account.Account/ListAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountServer is the server API for Account service.
// All implementations must embed UnimplementedAccountServer
// for forward compatibility
type AccountServer interface {
	CreateAccount(context.Context, *CreateAccountRequest) (*CreateAccountReply, error)
	UpdateAccount(context.Context, *UpdateAccountRequest) (*UpdateAccountReply, error)
	DeleteAccount(context.Context, *DeleteAccountRequest) (*DeleteAccountReply, error)
	GetAccount(context.Context, *GetAccountRequest) (*GetAccountReply, error)
	ListAccount(context.Context, *ListAccountRequest) (*ListAccountReply, error)
	mustEmbedUnimplementedAccountServer()
}

// UnimplementedAccountServer must be embedded to have forward compatible implementations.
type UnimplementedAccountServer struct {
}

func (UnimplementedAccountServer) CreateAccount(context.Context, *CreateAccountRequest) (*CreateAccountReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAccount not implemented")
}
func (UnimplementedAccountServer) UpdateAccount(context.Context, *UpdateAccountRequest) (*UpdateAccountReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAccount not implemented")
}
func (UnimplementedAccountServer) DeleteAccount(context.Context, *DeleteAccountRequest) (*DeleteAccountReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAccount not implemented")
}
func (UnimplementedAccountServer) GetAccount(context.Context, *GetAccountRequest) (*GetAccountReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccount not implemented")
}
func (UnimplementedAccountServer) ListAccount(context.Context, *ListAccountRequest) (*ListAccountReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAccount not implemented")
}
func (UnimplementedAccountServer) mustEmbedUnimplementedAccountServer() {}

// UnsafeAccountServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccountServer will
// result in compilation errors.
type UnsafeAccountServer interface {
	mustEmbedUnimplementedAccountServer()
}

func RegisterAccountServer(s grpc.ServiceRegistrar, srv AccountServer) {
	s.RegisterService(&Account_ServiceDesc, srv)
}

func _Account_CreateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).CreateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.auth.v1.account.Account/CreateAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).CreateAccount(ctx, req.(*CreateAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_UpdateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).UpdateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.auth.v1.account.Account/UpdateAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).UpdateAccount(ctx, req.(*UpdateAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_DeleteAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).DeleteAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.auth.v1.account.Account/DeleteAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).DeleteAccount(ctx, req.(*DeleteAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_GetAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).GetAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.auth.v1.account.Account/GetAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).GetAccount(ctx, req.(*GetAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_ListAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).ListAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.auth.v1.account.Account/ListAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).ListAccount(ctx, req.(*ListAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Account_ServiceDesc is the grpc.ServiceDesc for Account service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Account_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.auth.v1.account.Account",
	HandlerType: (*AccountServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAccount",
			Handler:    _Account_CreateAccount_Handler,
		},
		{
			MethodName: "UpdateAccount",
			Handler:    _Account_UpdateAccount_Handler,
		},
		{
			MethodName: "DeleteAccount",
			Handler:    _Account_DeleteAccount_Handler,
		},
		{
			MethodName: "GetAccount",
			Handler:    _Account_GetAccount_Handler,
		},
		{
			MethodName: "ListAccount",
			Handler:    _Account_ListAccount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/account/service.proto",
}
