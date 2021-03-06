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

// CouponClient is the client API for Coupon service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CouponClient interface {
	// 优惠券列表
	ListCoupon(ctx context.Context, in *ListCouponRequest, opts ...grpc.CallOption) (*ListCouponReply, error)
	// 抢券
	RushCollectCoupon(ctx context.Context, in *RushCollectCouponRequest, opts ...grpc.CallOption) (*RushCollectCouponReply, error)
	// 用户拥有的优惠券列表
	ListUserCoupon(ctx context.Context, in *ListUserCouponRequest, opts ...grpc.CallOption) (*ListUserCouponReply, error)
}

type couponClient struct {
	cc grpc.ClientConnInterface
}

func NewCouponClient(cc grpc.ClientConnInterface) CouponClient {
	return &couponClient{cc}
}

func (c *couponClient) ListCoupon(ctx context.Context, in *ListCouponRequest, opts ...grpc.CallOption) (*ListCouponReply, error) {
	out := new(ListCouponReply)
	err := c.cc.Invoke(ctx, "/api.coupon.v1.Coupon/ListCoupon", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *couponClient) RushCollectCoupon(ctx context.Context, in *RushCollectCouponRequest, opts ...grpc.CallOption) (*RushCollectCouponReply, error) {
	out := new(RushCollectCouponReply)
	err := c.cc.Invoke(ctx, "/api.coupon.v1.Coupon/RushCollectCoupon", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *couponClient) ListUserCoupon(ctx context.Context, in *ListUserCouponRequest, opts ...grpc.CallOption) (*ListUserCouponReply, error) {
	out := new(ListUserCouponReply)
	err := c.cc.Invoke(ctx, "/api.coupon.v1.Coupon/ListUserCoupon", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CouponServer is the server API for Coupon service.
// All implementations must embed UnimplementedCouponServer
// for forward compatibility
type CouponServer interface {
	// 优惠券列表
	ListCoupon(context.Context, *ListCouponRequest) (*ListCouponReply, error)
	// 抢券
	RushCollectCoupon(context.Context, *RushCollectCouponRequest) (*RushCollectCouponReply, error)
	// 用户拥有的优惠券列表
	ListUserCoupon(context.Context, *ListUserCouponRequest) (*ListUserCouponReply, error)
	mustEmbedUnimplementedCouponServer()
}

// UnimplementedCouponServer must be embedded to have forward compatible implementations.
type UnimplementedCouponServer struct {
}

func (UnimplementedCouponServer) ListCoupon(context.Context, *ListCouponRequest) (*ListCouponReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCoupon not implemented")
}
func (UnimplementedCouponServer) RushCollectCoupon(context.Context, *RushCollectCouponRequest) (*RushCollectCouponReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RushCollectCoupon not implemented")
}
func (UnimplementedCouponServer) ListUserCoupon(context.Context, *ListUserCouponRequest) (*ListUserCouponReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUserCoupon not implemented")
}
func (UnimplementedCouponServer) mustEmbedUnimplementedCouponServer() {}

// UnsafeCouponServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CouponServer will
// result in compilation errors.
type UnsafeCouponServer interface {
	mustEmbedUnimplementedCouponServer()
}

func RegisterCouponServer(s grpc.ServiceRegistrar, srv CouponServer) {
	s.RegisterService(&Coupon_ServiceDesc, srv)
}

func _Coupon_ListCoupon_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCouponRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CouponServer).ListCoupon(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.coupon.v1.Coupon/ListCoupon",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CouponServer).ListCoupon(ctx, req.(*ListCouponRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Coupon_RushCollectCoupon_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RushCollectCouponRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CouponServer).RushCollectCoupon(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.coupon.v1.Coupon/RushCollectCoupon",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CouponServer).RushCollectCoupon(ctx, req.(*RushCollectCouponRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Coupon_ListUserCoupon_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUserCouponRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CouponServer).ListUserCoupon(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.coupon.v1.Coupon/ListUserCoupon",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CouponServer).ListUserCoupon(ctx, req.(*ListUserCouponRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Coupon_ServiceDesc is the grpc.ServiceDesc for Coupon service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Coupon_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.coupon.v1.Coupon",
	HandlerType: (*CouponServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListCoupon",
			Handler:    _Coupon_ListCoupon_Handler,
		},
		{
			MethodName: "RushCollectCoupon",
			Handler:    _Coupon_RushCollectCoupon_Handler,
		},
		{
			MethodName: "ListUserCoupon",
			Handler:    _Coupon_ListUserCoupon_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/coupon/v1/coupon.proto",
}
