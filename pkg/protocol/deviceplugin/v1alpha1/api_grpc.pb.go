// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.4
// source: v1alpha1/api.proto

package v1alpha1

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

// RegistrationClient is the client API for Registration service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RegistrationClient interface {
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*Empty, error)
}

type registrationClient struct {
	cc grpc.ClientConnInterface
}

func NewRegistrationClient(cc grpc.ClientConnInterface) RegistrationClient {
	return &registrationClient{cc}
}

func (c *registrationClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/deviceplugin.v1alpha1.Registration/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RegistrationServer is the server API for Registration service.
// All implementations must embed UnimplementedRegistrationServer
// for forward compatibility
type RegistrationServer interface {
	Register(context.Context, *RegisterRequest) (*Empty, error)
	mustEmbedUnimplementedRegistrationServer()
}

// UnimplementedRegistrationServer must be embedded to have forward compatible implementations.
type UnimplementedRegistrationServer struct {
}

func (UnimplementedRegistrationServer) Register(context.Context, *RegisterRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedRegistrationServer) mustEmbedUnimplementedRegistrationServer() {}

// UnsafeRegistrationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RegistrationServer will
// result in compilation errors.
type UnsafeRegistrationServer interface {
	mustEmbedUnimplementedRegistrationServer()
}

func RegisterRegistrationServer(s grpc.ServiceRegistrar, srv RegistrationServer) {
	s.RegisterService(&Registration_ServiceDesc, srv)
}

func _Registration_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistrationServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/deviceplugin.v1alpha1.Registration/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistrationServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Registration_ServiceDesc is the grpc.ServiceDesc for Registration service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Registration_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "deviceplugin.v1alpha1.Registration",
	HandlerType: (*RegistrationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _Registration_Register_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1alpha1/api.proto",
}

// DevicePluginClient is the client API for DevicePlugin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DevicePluginClient interface {
	ListAndWatch(ctx context.Context, in *Empty, opts ...grpc.CallOption) (DevicePlugin_ListAndWatchClient, error)
}

type devicePluginClient struct {
	cc grpc.ClientConnInterface
}

func NewDevicePluginClient(cc grpc.ClientConnInterface) DevicePluginClient {
	return &devicePluginClient{cc}
}

func (c *devicePluginClient) ListAndWatch(ctx context.Context, in *Empty, opts ...grpc.CallOption) (DevicePlugin_ListAndWatchClient, error) {
	stream, err := c.cc.NewStream(ctx, &DevicePlugin_ServiceDesc.Streams[0], "/deviceplugin.v1alpha1.DevicePlugin/ListAndWatch", opts...)
	if err != nil {
		return nil, err
	}
	x := &devicePluginListAndWatchClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DevicePlugin_ListAndWatchClient interface {
	Recv() (*ListAndWatchResponse, error)
	grpc.ClientStream
}

type devicePluginListAndWatchClient struct {
	grpc.ClientStream
}

func (x *devicePluginListAndWatchClient) Recv() (*ListAndWatchResponse, error) {
	m := new(ListAndWatchResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DevicePluginServer is the server API for DevicePlugin service.
// All implementations must embed UnimplementedDevicePluginServer
// for forward compatibility
type DevicePluginServer interface {
	ListAndWatch(*Empty, DevicePlugin_ListAndWatchServer) error
	mustEmbedUnimplementedDevicePluginServer()
}

// UnimplementedDevicePluginServer must be embedded to have forward compatible implementations.
type UnimplementedDevicePluginServer struct {
}

func (UnimplementedDevicePluginServer) ListAndWatch(*Empty, DevicePlugin_ListAndWatchServer) error {
	return status.Errorf(codes.Unimplemented, "method ListAndWatch not implemented")
}
func (UnimplementedDevicePluginServer) mustEmbedUnimplementedDevicePluginServer() {}

// UnsafeDevicePluginServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DevicePluginServer will
// result in compilation errors.
type UnsafeDevicePluginServer interface {
	mustEmbedUnimplementedDevicePluginServer()
}

func RegisterDevicePluginServer(s grpc.ServiceRegistrar, srv DevicePluginServer) {
	s.RegisterService(&DevicePlugin_ServiceDesc, srv)
}

func _DevicePlugin_ListAndWatch_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DevicePluginServer).ListAndWatch(m, &devicePluginListAndWatchServer{stream})
}

type DevicePlugin_ListAndWatchServer interface {
	Send(*ListAndWatchResponse) error
	grpc.ServerStream
}

type devicePluginListAndWatchServer struct {
	grpc.ServerStream
}

func (x *devicePluginListAndWatchServer) Send(m *ListAndWatchResponse) error {
	return x.ServerStream.SendMsg(m)
}

// DevicePlugin_ServiceDesc is the grpc.ServiceDesc for DevicePlugin service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DevicePlugin_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "deviceplugin.v1alpha1.DevicePlugin",
	HandlerType: (*DevicePluginServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListAndWatch",
			Handler:       _DevicePlugin_ListAndWatch_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "v1alpha1/api.proto",
}
