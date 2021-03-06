// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AncientServiceClient is the client API for AncientService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AncientServiceClient interface {
	GetAncient(ctx context.Context, in *GetAncientReq, opts ...grpc.CallOption) (*Ancient, error)
	PutAncient(ctx context.Context, in *PutAncientReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UpdateAncient(ctx context.Context, in *UpdateAncientReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	SearchAncient(ctx context.Context, in *SearchAncientReq, opts ...grpc.CallOption) (*SearchAncientRes, error)
}

type ancientServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAncientServiceClient(cc grpc.ClientConnInterface) AncientServiceClient {
	return &ancientServiceClient{cc}
}

func (c *ancientServiceClient) GetAncient(ctx context.Context, in *GetAncientReq, opts ...grpc.CallOption) (*Ancient, error) {
	out := new(Ancient)
	err := c.cc.Invoke(ctx, "/api.AncientService/GetAncient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ancientServiceClient) PutAncient(ctx context.Context, in *PutAncientReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/api.AncientService/PutAncient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ancientServiceClient) UpdateAncient(ctx context.Context, in *UpdateAncientReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/api.AncientService/UpdateAncient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ancientServiceClient) SearchAncient(ctx context.Context, in *SearchAncientReq, opts ...grpc.CallOption) (*SearchAncientRes, error) {
	out := new(SearchAncientRes)
	err := c.cc.Invoke(ctx, "/api.AncientService/SearchAncient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AncientServiceServer is the server API for AncientService service.
// All implementations must embed UnimplementedAncientServiceServer
// for forward compatibility
type AncientServiceServer interface {
	GetAncient(context.Context, *GetAncientReq) (*Ancient, error)
	PutAncient(context.Context, *PutAncientReq) (*emptypb.Empty, error)
	UpdateAncient(context.Context, *UpdateAncientReq) (*emptypb.Empty, error)
	SearchAncient(context.Context, *SearchAncientReq) (*SearchAncientRes, error)
	mustEmbedUnimplementedAncientServiceServer()
}

// UnimplementedAncientServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAncientServiceServer struct {
}

func (UnimplementedAncientServiceServer) GetAncient(context.Context, *GetAncientReq) (*Ancient, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAncient not implemented")
}
func (UnimplementedAncientServiceServer) PutAncient(context.Context, *PutAncientReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutAncient not implemented")
}
func (UnimplementedAncientServiceServer) UpdateAncient(context.Context, *UpdateAncientReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAncient not implemented")
}
func (UnimplementedAncientServiceServer) SearchAncient(context.Context, *SearchAncientReq) (*SearchAncientRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchAncient not implemented")
}
func (UnimplementedAncientServiceServer) mustEmbedUnimplementedAncientServiceServer() {}

// UnsafeAncientServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AncientServiceServer will
// result in compilation errors.
type UnsafeAncientServiceServer interface {
	mustEmbedUnimplementedAncientServiceServer()
}

func RegisterAncientServiceServer(s grpc.ServiceRegistrar, srv AncientServiceServer) {
	s.RegisterService(&AncientService_ServiceDesc, srv)
}

func _AncientService_GetAncient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAncientReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AncientServiceServer).GetAncient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.AncientService/GetAncient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AncientServiceServer).GetAncient(ctx, req.(*GetAncientReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AncientService_PutAncient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutAncientReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AncientServiceServer).PutAncient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.AncientService/PutAncient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AncientServiceServer).PutAncient(ctx, req.(*PutAncientReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AncientService_UpdateAncient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAncientReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AncientServiceServer).UpdateAncient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.AncientService/UpdateAncient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AncientServiceServer).UpdateAncient(ctx, req.(*UpdateAncientReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AncientService_SearchAncient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchAncientReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AncientServiceServer).SearchAncient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.AncientService/SearchAncient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AncientServiceServer).SearchAncient(ctx, req.(*SearchAncientReq))
	}
	return interceptor(ctx, in, info, handler)
}

// AncientService_ServiceDesc is the grpc.ServiceDesc for AncientService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AncientService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.AncientService",
	HandlerType: (*AncientServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAncient",
			Handler:    _AncientService_GetAncient_Handler,
		},
		{
			MethodName: "PutAncient",
			Handler:    _AncientService_PutAncient_Handler,
		},
		{
			MethodName: "UpdateAncient",
			Handler:    _AncientService_UpdateAncient_Handler,
		},
		{
			MethodName: "SearchAncient",
			Handler:    _AncientService_SearchAncient_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/ancient.proto",
}
