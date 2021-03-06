// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

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

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserClient interface {
	// 用户注册
	RegisterUser(ctx context.Context, in *RegisterUserRequest, opts ...grpc.CallOption) (*RegisterUserReply, error)
	// 用户登陆-账号密码
	LoginByUsername(ctx context.Context, in *LoginByUsernameRequest, opts ...grpc.CallOption) (*LoginUserReply, error)
	// 用户登陆-短信验证码
	LoginByVerifyCode(ctx context.Context, in *LoginByVerifyCodeRequest, opts ...grpc.CallOption) (*LoginUserReply, error)
	// 获取用户信息
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserReply, error)
	// 检测用户是否正常
	CheckUserOK(ctx context.Context, in *CheckUserOKRequest, opts ...grpc.CallOption) (*CheckUserOKReply, error)
	// 地址列表
	ListAddress(ctx context.Context, in *ListAddressRequest, opts ...grpc.CallOption) (*ListAddressReply, error)
	// 地址创建
	CreateAddress(ctx context.Context, in *CreateAddressRequest, opts ...grpc.CallOption) (*CreateAddressReply, error)
	// 地址详情
	GetAddress(ctx context.Context, in *GetAddressRequest, opts ...grpc.CallOption) (*GetAddressReply, error)
	// 地址更新
	UpdateAddress(ctx context.Context, in *UpdateAddressRequest, opts ...grpc.CallOption) (*UpdateAddressReply, error)
	// 地址删除
	DeleteAddress(ctx context.Context, in *DeleteAddressRequest, opts ...grpc.CallOption) (*DeleteAddressReply, error)
}

type userClient struct {
	cc grpc.ClientConnInterface
}

func NewUserClient(cc grpc.ClientConnInterface) UserClient {
	return &userClient{cc}
}

func (c *userClient) RegisterUser(ctx context.Context, in *RegisterUserRequest, opts ...grpc.CallOption) (*RegisterUserReply, error) {
	out := new(RegisterUserReply)
	err := c.cc.Invoke(ctx, "/api.user.v1.User/RegisterUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) LoginByUsername(ctx context.Context, in *LoginByUsernameRequest, opts ...grpc.CallOption) (*LoginUserReply, error) {
	out := new(LoginUserReply)
	err := c.cc.Invoke(ctx, "/api.user.v1.User/LoginByUsername", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) LoginByVerifyCode(ctx context.Context, in *LoginByVerifyCodeRequest, opts ...grpc.CallOption) (*LoginUserReply, error) {
	out := new(LoginUserReply)
	err := c.cc.Invoke(ctx, "/api.user.v1.User/LoginByVerifyCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserReply, error) {
	out := new(GetUserReply)
	err := c.cc.Invoke(ctx, "/api.user.v1.User/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) CheckUserOK(ctx context.Context, in *CheckUserOKRequest, opts ...grpc.CallOption) (*CheckUserOKReply, error) {
	out := new(CheckUserOKReply)
	err := c.cc.Invoke(ctx, "/api.user.v1.User/CheckUserOK", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) ListAddress(ctx context.Context, in *ListAddressRequest, opts ...grpc.CallOption) (*ListAddressReply, error) {
	out := new(ListAddressReply)
	err := c.cc.Invoke(ctx, "/api.user.v1.User/ListAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) CreateAddress(ctx context.Context, in *CreateAddressRequest, opts ...grpc.CallOption) (*CreateAddressReply, error) {
	out := new(CreateAddressReply)
	err := c.cc.Invoke(ctx, "/api.user.v1.User/CreateAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetAddress(ctx context.Context, in *GetAddressRequest, opts ...grpc.CallOption) (*GetAddressReply, error) {
	out := new(GetAddressReply)
	err := c.cc.Invoke(ctx, "/api.user.v1.User/GetAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UpdateAddress(ctx context.Context, in *UpdateAddressRequest, opts ...grpc.CallOption) (*UpdateAddressReply, error) {
	out := new(UpdateAddressReply)
	err := c.cc.Invoke(ctx, "/api.user.v1.User/UpdateAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) DeleteAddress(ctx context.Context, in *DeleteAddressRequest, opts ...grpc.CallOption) (*DeleteAddressReply, error) {
	out := new(DeleteAddressReply)
	err := c.cc.Invoke(ctx, "/api.user.v1.User/DeleteAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
// All implementations must embed UnimplementedUserServer
// for forward compatibility
type UserServer interface {
	// 用户注册
	RegisterUser(context.Context, *RegisterUserRequest) (*RegisterUserReply, error)
	// 用户登陆-账号密码
	LoginByUsername(context.Context, *LoginByUsernameRequest) (*LoginUserReply, error)
	// 用户登陆-短信验证码
	LoginByVerifyCode(context.Context, *LoginByVerifyCodeRequest) (*LoginUserReply, error)
	// 获取用户信息
	GetUser(context.Context, *GetUserRequest) (*GetUserReply, error)
	// 检测用户是否正常
	CheckUserOK(context.Context, *CheckUserOKRequest) (*CheckUserOKReply, error)
	// 地址列表
	ListAddress(context.Context, *ListAddressRequest) (*ListAddressReply, error)
	// 地址创建
	CreateAddress(context.Context, *CreateAddressRequest) (*CreateAddressReply, error)
	// 地址详情
	GetAddress(context.Context, *GetAddressRequest) (*GetAddressReply, error)
	// 地址更新
	UpdateAddress(context.Context, *UpdateAddressRequest) (*UpdateAddressReply, error)
	// 地址删除
	DeleteAddress(context.Context, *DeleteAddressRequest) (*DeleteAddressReply, error)
	mustEmbedUnimplementedUserServer()
}

// UnimplementedUserServer must be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (UnimplementedUserServer) RegisterUser(context.Context, *RegisterUserRequest) (*RegisterUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterUser not implemented")
}
func (UnimplementedUserServer) LoginByUsername(context.Context, *LoginByUsernameRequest) (*LoginUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginByUsername not implemented")
}
func (UnimplementedUserServer) LoginByVerifyCode(context.Context, *LoginByVerifyCodeRequest) (*LoginUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginByVerifyCode not implemented")
}
func (UnimplementedUserServer) GetUser(context.Context, *GetUserRequest) (*GetUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedUserServer) CheckUserOK(context.Context, *CheckUserOKRequest) (*CheckUserOKReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckUserOK not implemented")
}
func (UnimplementedUserServer) ListAddress(context.Context, *ListAddressRequest) (*ListAddressReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAddress not implemented")
}
func (UnimplementedUserServer) CreateAddress(context.Context, *CreateAddressRequest) (*CreateAddressReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAddress not implemented")
}
func (UnimplementedUserServer) GetAddress(context.Context, *GetAddressRequest) (*GetAddressReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAddress not implemented")
}
func (UnimplementedUserServer) UpdateAddress(context.Context, *UpdateAddressRequest) (*UpdateAddressReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAddress not implemented")
}
func (UnimplementedUserServer) DeleteAddress(context.Context, *DeleteAddressRequest) (*DeleteAddressReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAddress not implemented")
}
func (UnimplementedUserServer) mustEmbedUnimplementedUserServer() {}

// UnsafeUserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServer will
// result in compilation errors.
type UnsafeUserServer interface {
	mustEmbedUnimplementedUserServer()
}

func RegisterUserServer(s grpc.ServiceRegistrar, srv UserServer) {
	s.RegisterService(&User_ServiceDesc, srv)
}

func _User_RegisterUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).RegisterUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.user.v1.User/RegisterUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).RegisterUser(ctx, req.(*RegisterUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_LoginByUsername_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginByUsernameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).LoginByUsername(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.user.v1.User/LoginByUsername",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).LoginByUsername(ctx, req.(*LoginByUsernameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_LoginByVerifyCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginByVerifyCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).LoginByVerifyCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.user.v1.User/LoginByVerifyCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).LoginByVerifyCode(ctx, req.(*LoginByVerifyCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.user.v1.User/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_CheckUserOK_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckUserOKRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).CheckUserOK(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.user.v1.User/CheckUserOK",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).CheckUserOK(ctx, req.(*CheckUserOKRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_ListAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).ListAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.user.v1.User/ListAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).ListAddress(ctx, req.(*ListAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_CreateAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).CreateAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.user.v1.User/CreateAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).CreateAddress(ctx, req.(*CreateAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.user.v1.User/GetAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetAddress(ctx, req.(*GetAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UpdateAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UpdateAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.user.v1.User/UpdateAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UpdateAddress(ctx, req.(*UpdateAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_DeleteAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).DeleteAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.user.v1.User/DeleteAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).DeleteAddress(ctx, req.(*DeleteAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// User_ServiceDesc is the grpc.ServiceDesc for User service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var User_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.user.v1.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterUser",
			Handler:    _User_RegisterUser_Handler,
		},
		{
			MethodName: "LoginByUsername",
			Handler:    _User_LoginByUsername_Handler,
		},
		{
			MethodName: "LoginByVerifyCode",
			Handler:    _User_LoginByVerifyCode_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _User_GetUser_Handler,
		},
		{
			MethodName: "CheckUserOK",
			Handler:    _User_CheckUserOK_Handler,
		},
		{
			MethodName: "ListAddress",
			Handler:    _User_ListAddress_Handler,
		},
		{
			MethodName: "CreateAddress",
			Handler:    _User_CreateAddress_Handler,
		},
		{
			MethodName: "GetAddress",
			Handler:    _User_GetAddress_Handler,
		},
		{
			MethodName: "UpdateAddress",
			Handler:    _User_UpdateAddress_Handler,
		},
		{
			MethodName: "DeleteAddress",
			Handler:    _User_DeleteAddress_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/user/v1/user.proto",
}
