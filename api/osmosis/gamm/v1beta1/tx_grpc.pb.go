// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: osmosis/gamm/v1beta1/tx.proto

package gammv1beta1

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

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	JoinPool(ctx context.Context, in *MsgJoinPool, opts ...grpc.CallOption) (*MsgJoinPoolResponse, error)
	ExitPool(ctx context.Context, in *MsgExitPool, opts ...grpc.CallOption) (*MsgExitPoolResponse, error)
	SwapExactAmountIn(ctx context.Context, in *MsgSwapExactAmountIn, opts ...grpc.CallOption) (*MsgSwapExactAmountInResponse, error)
	SwapExactAmountOut(ctx context.Context, in *MsgSwapExactAmountOut, opts ...grpc.CallOption) (*MsgSwapExactAmountOutResponse, error)
	JoinSwapExternAmountIn(ctx context.Context, in *MsgJoinSwapExternAmountIn, opts ...grpc.CallOption) (*MsgJoinSwapExternAmountInResponse, error)
	JoinSwapShareAmountOut(ctx context.Context, in *MsgJoinSwapShareAmountOut, opts ...grpc.CallOption) (*MsgJoinSwapShareAmountOutResponse, error)
	ExitSwapExternAmountOut(ctx context.Context, in *MsgExitSwapExternAmountOut, opts ...grpc.CallOption) (*MsgExitSwapExternAmountOutResponse, error)
	ExitSwapShareAmountIn(ctx context.Context, in *MsgExitSwapShareAmountIn, opts ...grpc.CallOption) (*MsgExitSwapShareAmountInResponse, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) JoinPool(ctx context.Context, in *MsgJoinPool, opts ...grpc.CallOption) (*MsgJoinPoolResponse, error) {
	out := new(MsgJoinPoolResponse)
	err := c.cc.Invoke(ctx, "/osmosis.gamm.v1beta1.Msg/JoinPool", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) ExitPool(ctx context.Context, in *MsgExitPool, opts ...grpc.CallOption) (*MsgExitPoolResponse, error) {
	out := new(MsgExitPoolResponse)
	err := c.cc.Invoke(ctx, "/osmosis.gamm.v1beta1.Msg/ExitPool", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) SwapExactAmountIn(ctx context.Context, in *MsgSwapExactAmountIn, opts ...grpc.CallOption) (*MsgSwapExactAmountInResponse, error) {
	out := new(MsgSwapExactAmountInResponse)
	err := c.cc.Invoke(ctx, "/osmosis.gamm.v1beta1.Msg/SwapExactAmountIn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) SwapExactAmountOut(ctx context.Context, in *MsgSwapExactAmountOut, opts ...grpc.CallOption) (*MsgSwapExactAmountOutResponse, error) {
	out := new(MsgSwapExactAmountOutResponse)
	err := c.cc.Invoke(ctx, "/osmosis.gamm.v1beta1.Msg/SwapExactAmountOut", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) JoinSwapExternAmountIn(ctx context.Context, in *MsgJoinSwapExternAmountIn, opts ...grpc.CallOption) (*MsgJoinSwapExternAmountInResponse, error) {
	out := new(MsgJoinSwapExternAmountInResponse)
	err := c.cc.Invoke(ctx, "/osmosis.gamm.v1beta1.Msg/JoinSwapExternAmountIn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) JoinSwapShareAmountOut(ctx context.Context, in *MsgJoinSwapShareAmountOut, opts ...grpc.CallOption) (*MsgJoinSwapShareAmountOutResponse, error) {
	out := new(MsgJoinSwapShareAmountOutResponse)
	err := c.cc.Invoke(ctx, "/osmosis.gamm.v1beta1.Msg/JoinSwapShareAmountOut", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) ExitSwapExternAmountOut(ctx context.Context, in *MsgExitSwapExternAmountOut, opts ...grpc.CallOption) (*MsgExitSwapExternAmountOutResponse, error) {
	out := new(MsgExitSwapExternAmountOutResponse)
	err := c.cc.Invoke(ctx, "/osmosis.gamm.v1beta1.Msg/ExitSwapExternAmountOut", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) ExitSwapShareAmountIn(ctx context.Context, in *MsgExitSwapShareAmountIn, opts ...grpc.CallOption) (*MsgExitSwapShareAmountInResponse, error) {
	out := new(MsgExitSwapShareAmountInResponse)
	err := c.cc.Invoke(ctx, "/osmosis.gamm.v1beta1.Msg/ExitSwapShareAmountIn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility
type MsgServer interface {
	JoinPool(context.Context, *MsgJoinPool) (*MsgJoinPoolResponse, error)
	ExitPool(context.Context, *MsgExitPool) (*MsgExitPoolResponse, error)
	SwapExactAmountIn(context.Context, *MsgSwapExactAmountIn) (*MsgSwapExactAmountInResponse, error)
	SwapExactAmountOut(context.Context, *MsgSwapExactAmountOut) (*MsgSwapExactAmountOutResponse, error)
	JoinSwapExternAmountIn(context.Context, *MsgJoinSwapExternAmountIn) (*MsgJoinSwapExternAmountInResponse, error)
	JoinSwapShareAmountOut(context.Context, *MsgJoinSwapShareAmountOut) (*MsgJoinSwapShareAmountOutResponse, error)
	ExitSwapExternAmountOut(context.Context, *MsgExitSwapExternAmountOut) (*MsgExitSwapExternAmountOutResponse, error)
	ExitSwapShareAmountIn(context.Context, *MsgExitSwapShareAmountIn) (*MsgExitSwapShareAmountInResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) JoinPool(context.Context, *MsgJoinPool) (*MsgJoinPoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinPool not implemented")
}
func (UnimplementedMsgServer) ExitPool(context.Context, *MsgExitPool) (*MsgExitPoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExitPool not implemented")
}
func (UnimplementedMsgServer) SwapExactAmountIn(context.Context, *MsgSwapExactAmountIn) (*MsgSwapExactAmountInResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SwapExactAmountIn not implemented")
}
func (UnimplementedMsgServer) SwapExactAmountOut(context.Context, *MsgSwapExactAmountOut) (*MsgSwapExactAmountOutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SwapExactAmountOut not implemented")
}
func (UnimplementedMsgServer) JoinSwapExternAmountIn(context.Context, *MsgJoinSwapExternAmountIn) (*MsgJoinSwapExternAmountInResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinSwapExternAmountIn not implemented")
}
func (UnimplementedMsgServer) JoinSwapShareAmountOut(context.Context, *MsgJoinSwapShareAmountOut) (*MsgJoinSwapShareAmountOutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinSwapShareAmountOut not implemented")
}
func (UnimplementedMsgServer) ExitSwapExternAmountOut(context.Context, *MsgExitSwapExternAmountOut) (*MsgExitSwapExternAmountOutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExitSwapExternAmountOut not implemented")
}
func (UnimplementedMsgServer) ExitSwapShareAmountIn(context.Context, *MsgExitSwapShareAmountIn) (*MsgExitSwapShareAmountInResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExitSwapShareAmountIn not implemented")
}
func (UnimplementedMsgServer) mustEmbedUnimplementedMsgServer() {}

// UnsafeMsgServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MsgServer will
// result in compilation errors.
type UnsafeMsgServer interface {
	mustEmbedUnimplementedMsgServer()
}

func RegisterMsgServer(s grpc.ServiceRegistrar, srv MsgServer) {
	s.RegisterService(&Msg_ServiceDesc, srv)
}

func _Msg_JoinPool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgJoinPool)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).JoinPool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/osmosis.gamm.v1beta1.Msg/JoinPool",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).JoinPool(ctx, req.(*MsgJoinPool))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_ExitPool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgExitPool)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ExitPool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/osmosis.gamm.v1beta1.Msg/ExitPool",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ExitPool(ctx, req.(*MsgExitPool))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_SwapExactAmountIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSwapExactAmountIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SwapExactAmountIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/osmosis.gamm.v1beta1.Msg/SwapExactAmountIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SwapExactAmountIn(ctx, req.(*MsgSwapExactAmountIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_SwapExactAmountOut_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSwapExactAmountOut)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SwapExactAmountOut(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/osmosis.gamm.v1beta1.Msg/SwapExactAmountOut",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SwapExactAmountOut(ctx, req.(*MsgSwapExactAmountOut))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_JoinSwapExternAmountIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgJoinSwapExternAmountIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).JoinSwapExternAmountIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/osmosis.gamm.v1beta1.Msg/JoinSwapExternAmountIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).JoinSwapExternAmountIn(ctx, req.(*MsgJoinSwapExternAmountIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_JoinSwapShareAmountOut_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgJoinSwapShareAmountOut)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).JoinSwapShareAmountOut(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/osmosis.gamm.v1beta1.Msg/JoinSwapShareAmountOut",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).JoinSwapShareAmountOut(ctx, req.(*MsgJoinSwapShareAmountOut))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_ExitSwapExternAmountOut_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgExitSwapExternAmountOut)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ExitSwapExternAmountOut(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/osmosis.gamm.v1beta1.Msg/ExitSwapExternAmountOut",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ExitSwapExternAmountOut(ctx, req.(*MsgExitSwapExternAmountOut))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_ExitSwapShareAmountIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgExitSwapShareAmountIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ExitSwapShareAmountIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/osmosis.gamm.v1beta1.Msg/ExitSwapShareAmountIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ExitSwapShareAmountIn(ctx, req.(*MsgExitSwapShareAmountIn))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "osmosis.gamm.v1beta1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "JoinPool",
			Handler:    _Msg_JoinPool_Handler,
		},
		{
			MethodName: "ExitPool",
			Handler:    _Msg_ExitPool_Handler,
		},
		{
			MethodName: "SwapExactAmountIn",
			Handler:    _Msg_SwapExactAmountIn_Handler,
		},
		{
			MethodName: "SwapExactAmountOut",
			Handler:    _Msg_SwapExactAmountOut_Handler,
		},
		{
			MethodName: "JoinSwapExternAmountIn",
			Handler:    _Msg_JoinSwapExternAmountIn_Handler,
		},
		{
			MethodName: "JoinSwapShareAmountOut",
			Handler:    _Msg_JoinSwapShareAmountOut_Handler,
		},
		{
			MethodName: "ExitSwapExternAmountOut",
			Handler:    _Msg_ExitSwapExternAmountOut_Handler,
		},
		{
			MethodName: "ExitSwapShareAmountIn",
			Handler:    _Msg_ExitSwapShareAmountIn_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "osmosis/gamm/v1beta1/tx.proto",
}
