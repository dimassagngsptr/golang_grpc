// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v4.23.3
// source: common/model/address.proto

package model

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Addresses_LIST_ADDRESS_FullMethodName = "/model.Addresses/LIST_ADDRESS"
	Addresses_ADD_ADDRESS_FullMethodName  = "/model.Addresses/ADD_ADDRESS"
)

// AddressesClient is the client API for Addresses service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AddressesClient interface {
	LIST_ADDRESS(ctx context.Context, in *AddressUserId, opts ...grpc.CallOption) (*AddressList, error)
	ADD_ADDRESS(ctx context.Context, in *AddressAndUserId, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type addressesClient struct {
	cc grpc.ClientConnInterface
}

func NewAddressesClient(cc grpc.ClientConnInterface) AddressesClient {
	return &addressesClient{cc}
}

func (c *addressesClient) LIST_ADDRESS(ctx context.Context, in *AddressUserId, opts ...grpc.CallOption) (*AddressList, error) {
	// cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddressList)
	err := c.cc.Invoke(ctx, Addresses_LIST_ADDRESS_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addressesClient) ADD_ADDRESS(ctx context.Context, in *AddressAndUserId, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	// cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Addresses_ADD_ADDRESS_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AddressesServer is the server API for Addresses service.
// All implementations must embed UnimplementedAddressesServer
// for forward compatibility.
type AddressesServer interface {
	LIST_ADDRESS(context.Context, *AddressUserId) (*AddressList, error)
	ADD_ADDRESS(context.Context, *AddressAndUserId) (*emptypb.Empty, error)
	mustEmbedUnimplementedAddressesServer()
}

// UnimplementedAddressesServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAddressesServer struct{}

func (UnimplementedAddressesServer) LIST_ADDRESS(context.Context, *AddressUserId) (*AddressList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LIST_ADDRESS not implemented")
}
func (UnimplementedAddressesServer) ADD_ADDRESS(context.Context, *AddressAndUserId) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ADD_ADDRESS not implemented")
}
func (UnimplementedAddressesServer) mustEmbedUnimplementedAddressesServer() {}
func (UnimplementedAddressesServer) testEmbeddedByValue()                   {}

// UnsafeAddressesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AddressesServer will
// result in compilation errors.
type UnsafeAddressesServer interface {
	mustEmbedUnimplementedAddressesServer()
}

func RegisterAddressesServer(s grpc.ServiceRegistrar, srv AddressesServer) {
	// If the following call pancis, it indicates UnimplementedAddressesServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Addresses_ServiceDesc, srv)
}

func _Addresses_LIST_ADDRESS_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddressUserId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddressesServer).LIST_ADDRESS(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Addresses_LIST_ADDRESS_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddressesServer).LIST_ADDRESS(ctx, req.(*AddressUserId))
	}
	return interceptor(ctx, in, info, handler)
}

func _Addresses_ADD_ADDRESS_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddressAndUserId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddressesServer).ADD_ADDRESS(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Addresses_ADD_ADDRESS_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddressesServer).ADD_ADDRESS(ctx, req.(*AddressAndUserId))
	}
	return interceptor(ctx, in, info, handler)
}

// Addresses_ServiceDesc is the grpc.ServiceDesc for Addresses service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Addresses_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "model.Addresses",
	HandlerType: (*AddressesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LIST_ADDRESS",
			Handler:    _Addresses_LIST_ADDRESS_Handler,
		},
		{
			MethodName: "ADD_ADDRESS",
			Handler:    _Addresses_ADD_ADDRESS_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "common/model/address.proto",
}