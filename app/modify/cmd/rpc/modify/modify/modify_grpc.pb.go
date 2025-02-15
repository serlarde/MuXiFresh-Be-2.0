// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: modify.proto

package modify

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
	Modify_ModifytheUsertype_FullMethodName   = "/modify.Modify/ModifytheUsertype"
	Modify_ModifytheUseravatar_FullMethodName = "/modify.Modify/ModifytheUseravatar"
	Modify_ModifytheUsername_FullMethodName   = "/modify.Modify/ModifytheUsername"
)

// ModifyClient is the client API for Modify service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ModifyClient interface {
	ModifytheUsertype(ctx context.Context, in *ModifytheUsertypeReq, opts ...grpc.CallOption) (*ModifytheUsertypeResp, error)
	ModifytheUseravatar(ctx context.Context, in *ModifytheUseravatarReq, opts ...grpc.CallOption) (*ModifytheUseravatarResp, error)
	ModifytheUsername(ctx context.Context, in *ModifytheUsernameReq, opts ...grpc.CallOption) (*ModifytheUsernameResp, error)
}

type modifyClient struct {
	cc grpc.ClientConnInterface
}

func NewModifyClient(cc grpc.ClientConnInterface) ModifyClient {
	return &modifyClient{cc}
}

func (c *modifyClient) ModifytheUsertype(ctx context.Context, in *ModifytheUsertypeReq, opts ...grpc.CallOption) (*ModifytheUsertypeResp, error) {
	out := new(ModifytheUsertypeResp)
	err := c.cc.Invoke(ctx, Modify_ModifytheUsertype_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modifyClient) ModifytheUseravatar(ctx context.Context, in *ModifytheUseravatarReq, opts ...grpc.CallOption) (*ModifytheUseravatarResp, error) {
	out := new(ModifytheUseravatarResp)
	err := c.cc.Invoke(ctx, Modify_ModifytheUseravatar_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modifyClient) ModifytheUsername(ctx context.Context, in *ModifytheUsernameReq, opts ...grpc.CallOption) (*ModifytheUsernameResp, error) {
	out := new(ModifytheUsernameResp)
	err := c.cc.Invoke(ctx, Modify_ModifytheUsername_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ModifyServer is the server API for Modify service.
// All implementations must embed UnimplementedModifyServer
// for forward compatibility
type ModifyServer interface {
	ModifytheUsertype(context.Context, *ModifytheUsertypeReq) (*ModifytheUsertypeResp, error)
	ModifytheUseravatar(context.Context, *ModifytheUseravatarReq) (*ModifytheUseravatarResp, error)
	ModifytheUsername(context.Context, *ModifytheUsernameReq) (*ModifytheUsernameResp, error)
	mustEmbedUnimplementedModifyServer()
}

// UnimplementedModifyServer must be embedded to have forward compatible implementations.
type UnimplementedModifyServer struct {
}

func (UnimplementedModifyServer) ModifytheUsertype(context.Context, *ModifytheUsertypeReq) (*ModifytheUsertypeResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModifytheUsertype not implemented")
}
func (UnimplementedModifyServer) ModifytheUseravatar(context.Context, *ModifytheUseravatarReq) (*ModifytheUseravatarResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModifytheUseravatar not implemented")
}
func (UnimplementedModifyServer) ModifytheUsername(context.Context, *ModifytheUsernameReq) (*ModifytheUsernameResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModifytheUsername not implemented")
}
func (UnimplementedModifyServer) mustEmbedUnimplementedModifyServer() {}

// UnsafeModifyServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ModifyServer will
// result in compilation errors.
type UnsafeModifyServer interface {
	mustEmbedUnimplementedModifyServer()
}

func RegisterModifyServer(s grpc.ServiceRegistrar, srv ModifyServer) {
	s.RegisterService(&Modify_ServiceDesc, srv)
}

func _Modify_ModifytheUsertype_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModifytheUsertypeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModifyServer).ModifytheUsertype(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Modify_ModifytheUsertype_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModifyServer).ModifytheUsertype(ctx, req.(*ModifytheUsertypeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Modify_ModifytheUseravatar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModifytheUseravatarReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModifyServer).ModifytheUseravatar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Modify_ModifytheUseravatar_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModifyServer).ModifytheUseravatar(ctx, req.(*ModifytheUseravatarReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Modify_ModifytheUsername_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModifytheUsernameReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModifyServer).ModifytheUsername(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Modify_ModifytheUsername_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModifyServer).ModifytheUsername(ctx, req.(*ModifytheUsernameReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Modify_ServiceDesc is the grpc.ServiceDesc for Modify service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Modify_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "modify.Modify",
	HandlerType: (*ModifyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ModifytheUsertype",
			Handler:    _Modify_ModifytheUsertype_Handler,
		},
		{
			MethodName: "ModifytheUseravatar",
			Handler:    _Modify_ModifytheUseravatar_Handler,
		},
		{
			MethodName: "ModifytheUsername",
			Handler:    _Modify_ModifytheUsername_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "modify.proto",
}
