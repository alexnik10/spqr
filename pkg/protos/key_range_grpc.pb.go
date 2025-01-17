// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: protos/key_range.proto

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

const (
	KeyRangeService_ListKeyRange_FullMethodName     = "/spqr.KeyRangeService/ListKeyRange"
	KeyRangeService_LockKeyRange_FullMethodName     = "/spqr.KeyRangeService/LockKeyRange"
	KeyRangeService_AddKeyRange_FullMethodName      = "/spqr.KeyRangeService/AddKeyRange"
	KeyRangeService_DropKeyRange_FullMethodName     = "/spqr.KeyRangeService/DropKeyRange"
	KeyRangeService_DropAllKeyRanges_FullMethodName = "/spqr.KeyRangeService/DropAllKeyRanges"
	KeyRangeService_UnlockKeyRange_FullMethodName   = "/spqr.KeyRangeService/UnlockKeyRange"
	KeyRangeService_SplitKeyRange_FullMethodName    = "/spqr.KeyRangeService/SplitKeyRange"
	KeyRangeService_MergeKeyRange_FullMethodName    = "/spqr.KeyRangeService/MergeKeyRange"
	KeyRangeService_MoveKeyRange_FullMethodName     = "/spqr.KeyRangeService/MoveKeyRange"
	KeyRangeService_ResolveKeyRange_FullMethodName  = "/spqr.KeyRangeService/ResolveKeyRange"
)

// KeyRangeServiceClient is the client API for KeyRangeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KeyRangeServiceClient interface {
	ListKeyRange(ctx context.Context, in *ListKeyRangeRequest, opts ...grpc.CallOption) (*KeyRangeReply, error)
	LockKeyRange(ctx context.Context, in *LockKeyRangeRequest, opts ...grpc.CallOption) (*ModifyReply, error)
	AddKeyRange(ctx context.Context, in *AddKeyRangeRequest, opts ...grpc.CallOption) (*ModifyReply, error)
	DropKeyRange(ctx context.Context, in *DropKeyRangeRequest, opts ...grpc.CallOption) (*ModifyReply, error)
	DropAllKeyRanges(ctx context.Context, in *DropAllKeyRangesRequest, opts ...grpc.CallOption) (*DropAllKeyRangesResponse, error)
	UnlockKeyRange(ctx context.Context, in *UnlockKeyRangeRequest, opts ...grpc.CallOption) (*ModifyReply, error)
	SplitKeyRange(ctx context.Context, in *SplitKeyRangeRequest, opts ...grpc.CallOption) (*ModifyReply, error)
	MergeKeyRange(ctx context.Context, in *MergeKeyRangeRequest, opts ...grpc.CallOption) (*ModifyReply, error)
	MoveKeyRange(ctx context.Context, in *MoveKeyRangeRequest, opts ...grpc.CallOption) (*ModifyReply, error)
	ResolveKeyRange(ctx context.Context, in *ResolveKeyRangeRequest, opts ...grpc.CallOption) (*ResolveKeyRangeReply, error)
}

type keyRangeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewKeyRangeServiceClient(cc grpc.ClientConnInterface) KeyRangeServiceClient {
	return &keyRangeServiceClient{cc}
}

func (c *keyRangeServiceClient) ListKeyRange(ctx context.Context, in *ListKeyRangeRequest, opts ...grpc.CallOption) (*KeyRangeReply, error) {
	out := new(KeyRangeReply)
	err := c.cc.Invoke(ctx, KeyRangeService_ListKeyRange_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyRangeServiceClient) LockKeyRange(ctx context.Context, in *LockKeyRangeRequest, opts ...grpc.CallOption) (*ModifyReply, error) {
	out := new(ModifyReply)
	err := c.cc.Invoke(ctx, KeyRangeService_LockKeyRange_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyRangeServiceClient) AddKeyRange(ctx context.Context, in *AddKeyRangeRequest, opts ...grpc.CallOption) (*ModifyReply, error) {
	out := new(ModifyReply)
	err := c.cc.Invoke(ctx, KeyRangeService_AddKeyRange_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyRangeServiceClient) DropKeyRange(ctx context.Context, in *DropKeyRangeRequest, opts ...grpc.CallOption) (*ModifyReply, error) {
	out := new(ModifyReply)
	err := c.cc.Invoke(ctx, KeyRangeService_DropKeyRange_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyRangeServiceClient) DropAllKeyRanges(ctx context.Context, in *DropAllKeyRangesRequest, opts ...grpc.CallOption) (*DropAllKeyRangesResponse, error) {
	out := new(DropAllKeyRangesResponse)
	err := c.cc.Invoke(ctx, KeyRangeService_DropAllKeyRanges_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyRangeServiceClient) UnlockKeyRange(ctx context.Context, in *UnlockKeyRangeRequest, opts ...grpc.CallOption) (*ModifyReply, error) {
	out := new(ModifyReply)
	err := c.cc.Invoke(ctx, KeyRangeService_UnlockKeyRange_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyRangeServiceClient) SplitKeyRange(ctx context.Context, in *SplitKeyRangeRequest, opts ...grpc.CallOption) (*ModifyReply, error) {
	out := new(ModifyReply)
	err := c.cc.Invoke(ctx, KeyRangeService_SplitKeyRange_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyRangeServiceClient) MergeKeyRange(ctx context.Context, in *MergeKeyRangeRequest, opts ...grpc.CallOption) (*ModifyReply, error) {
	out := new(ModifyReply)
	err := c.cc.Invoke(ctx, KeyRangeService_MergeKeyRange_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyRangeServiceClient) MoveKeyRange(ctx context.Context, in *MoveKeyRangeRequest, opts ...grpc.CallOption) (*ModifyReply, error) {
	out := new(ModifyReply)
	err := c.cc.Invoke(ctx, KeyRangeService_MoveKeyRange_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyRangeServiceClient) ResolveKeyRange(ctx context.Context, in *ResolveKeyRangeRequest, opts ...grpc.CallOption) (*ResolveKeyRangeReply, error) {
	out := new(ResolveKeyRangeReply)
	err := c.cc.Invoke(ctx, KeyRangeService_ResolveKeyRange_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KeyRangeServiceServer is the server API for KeyRangeService service.
// All implementations must embed UnimplementedKeyRangeServiceServer
// for forward compatibility
type KeyRangeServiceServer interface {
	ListKeyRange(context.Context, *ListKeyRangeRequest) (*KeyRangeReply, error)
	LockKeyRange(context.Context, *LockKeyRangeRequest) (*ModifyReply, error)
	AddKeyRange(context.Context, *AddKeyRangeRequest) (*ModifyReply, error)
	DropKeyRange(context.Context, *DropKeyRangeRequest) (*ModifyReply, error)
	DropAllKeyRanges(context.Context, *DropAllKeyRangesRequest) (*DropAllKeyRangesResponse, error)
	UnlockKeyRange(context.Context, *UnlockKeyRangeRequest) (*ModifyReply, error)
	SplitKeyRange(context.Context, *SplitKeyRangeRequest) (*ModifyReply, error)
	MergeKeyRange(context.Context, *MergeKeyRangeRequest) (*ModifyReply, error)
	MoveKeyRange(context.Context, *MoveKeyRangeRequest) (*ModifyReply, error)
	ResolveKeyRange(context.Context, *ResolveKeyRangeRequest) (*ResolveKeyRangeReply, error)
	mustEmbedUnimplementedKeyRangeServiceServer()
}

// UnimplementedKeyRangeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedKeyRangeServiceServer struct {
}

func (UnimplementedKeyRangeServiceServer) ListKeyRange(context.Context, *ListKeyRangeRequest) (*KeyRangeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListKeyRange not implemented")
}
func (UnimplementedKeyRangeServiceServer) LockKeyRange(context.Context, *LockKeyRangeRequest) (*ModifyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LockKeyRange not implemented")
}
func (UnimplementedKeyRangeServiceServer) AddKeyRange(context.Context, *AddKeyRangeRequest) (*ModifyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddKeyRange not implemented")
}
func (UnimplementedKeyRangeServiceServer) DropKeyRange(context.Context, *DropKeyRangeRequest) (*ModifyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DropKeyRange not implemented")
}
func (UnimplementedKeyRangeServiceServer) DropAllKeyRanges(context.Context, *DropAllKeyRangesRequest) (*DropAllKeyRangesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DropAllKeyRanges not implemented")
}
func (UnimplementedKeyRangeServiceServer) UnlockKeyRange(context.Context, *UnlockKeyRangeRequest) (*ModifyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnlockKeyRange not implemented")
}
func (UnimplementedKeyRangeServiceServer) SplitKeyRange(context.Context, *SplitKeyRangeRequest) (*ModifyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SplitKeyRange not implemented")
}
func (UnimplementedKeyRangeServiceServer) MergeKeyRange(context.Context, *MergeKeyRangeRequest) (*ModifyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MergeKeyRange not implemented")
}
func (UnimplementedKeyRangeServiceServer) MoveKeyRange(context.Context, *MoveKeyRangeRequest) (*ModifyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MoveKeyRange not implemented")
}
func (UnimplementedKeyRangeServiceServer) ResolveKeyRange(context.Context, *ResolveKeyRangeRequest) (*ResolveKeyRangeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResolveKeyRange not implemented")
}
func (UnimplementedKeyRangeServiceServer) mustEmbedUnimplementedKeyRangeServiceServer() {}

// UnsafeKeyRangeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KeyRangeServiceServer will
// result in compilation errors.
type UnsafeKeyRangeServiceServer interface {
	mustEmbedUnimplementedKeyRangeServiceServer()
}

func RegisterKeyRangeServiceServer(s grpc.ServiceRegistrar, srv KeyRangeServiceServer) {
	s.RegisterService(&KeyRangeService_ServiceDesc, srv)
}

func _KeyRangeService_ListKeyRange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListKeyRangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyRangeServiceServer).ListKeyRange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KeyRangeService_ListKeyRange_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyRangeServiceServer).ListKeyRange(ctx, req.(*ListKeyRangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyRangeService_LockKeyRange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LockKeyRangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyRangeServiceServer).LockKeyRange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KeyRangeService_LockKeyRange_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyRangeServiceServer).LockKeyRange(ctx, req.(*LockKeyRangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyRangeService_AddKeyRange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddKeyRangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyRangeServiceServer).AddKeyRange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KeyRangeService_AddKeyRange_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyRangeServiceServer).AddKeyRange(ctx, req.(*AddKeyRangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyRangeService_DropKeyRange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DropKeyRangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyRangeServiceServer).DropKeyRange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KeyRangeService_DropKeyRange_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyRangeServiceServer).DropKeyRange(ctx, req.(*DropKeyRangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyRangeService_DropAllKeyRanges_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DropAllKeyRangesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyRangeServiceServer).DropAllKeyRanges(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KeyRangeService_DropAllKeyRanges_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyRangeServiceServer).DropAllKeyRanges(ctx, req.(*DropAllKeyRangesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyRangeService_UnlockKeyRange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnlockKeyRangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyRangeServiceServer).UnlockKeyRange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KeyRangeService_UnlockKeyRange_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyRangeServiceServer).UnlockKeyRange(ctx, req.(*UnlockKeyRangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyRangeService_SplitKeyRange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SplitKeyRangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyRangeServiceServer).SplitKeyRange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KeyRangeService_SplitKeyRange_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyRangeServiceServer).SplitKeyRange(ctx, req.(*SplitKeyRangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyRangeService_MergeKeyRange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MergeKeyRangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyRangeServiceServer).MergeKeyRange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KeyRangeService_MergeKeyRange_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyRangeServiceServer).MergeKeyRange(ctx, req.(*MergeKeyRangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyRangeService_MoveKeyRange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MoveKeyRangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyRangeServiceServer).MoveKeyRange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KeyRangeService_MoveKeyRange_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyRangeServiceServer).MoveKeyRange(ctx, req.(*MoveKeyRangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyRangeService_ResolveKeyRange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResolveKeyRangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyRangeServiceServer).ResolveKeyRange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KeyRangeService_ResolveKeyRange_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyRangeServiceServer).ResolveKeyRange(ctx, req.(*ResolveKeyRangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// KeyRangeService_ServiceDesc is the grpc.ServiceDesc for KeyRangeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KeyRangeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "spqr.KeyRangeService",
	HandlerType: (*KeyRangeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListKeyRange",
			Handler:    _KeyRangeService_ListKeyRange_Handler,
		},
		{
			MethodName: "LockKeyRange",
			Handler:    _KeyRangeService_LockKeyRange_Handler,
		},
		{
			MethodName: "AddKeyRange",
			Handler:    _KeyRangeService_AddKeyRange_Handler,
		},
		{
			MethodName: "DropKeyRange",
			Handler:    _KeyRangeService_DropKeyRange_Handler,
		},
		{
			MethodName: "DropAllKeyRanges",
			Handler:    _KeyRangeService_DropAllKeyRanges_Handler,
		},
		{
			MethodName: "UnlockKeyRange",
			Handler:    _KeyRangeService_UnlockKeyRange_Handler,
		},
		{
			MethodName: "SplitKeyRange",
			Handler:    _KeyRangeService_SplitKeyRange_Handler,
		},
		{
			MethodName: "MergeKeyRange",
			Handler:    _KeyRangeService_MergeKeyRange_Handler,
		},
		{
			MethodName: "MoveKeyRange",
			Handler:    _KeyRangeService_MoveKeyRange_Handler,
		},
		{
			MethodName: "ResolveKeyRange",
			Handler:    _KeyRangeService_ResolveKeyRange_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/key_range.proto",
}
