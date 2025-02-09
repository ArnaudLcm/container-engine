// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.12.4
// source: google/rpc/service.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ContainerDaemonService_GetContainers_FullMethodName   = "/daemon.ContainerDaemonService/GetContainers"
	ContainerDaemonService_CreateContainer_FullMethodName = "/daemon.ContainerDaemonService/CreateContainer"
)

// ContainerDaemonServiceClient is the client API for ContainerDaemonService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ContainerDaemonServiceClient interface {
	GetContainers(ctx context.Context, in *GetContainersRequest, opts ...grpc.CallOption) (*GetContainersResponse, error)
	CreateContainer(ctx context.Context, in *CreateContainerRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[CreateContainerResponse], error)
}

type containerDaemonServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewContainerDaemonServiceClient(cc grpc.ClientConnInterface) ContainerDaemonServiceClient {
	return &containerDaemonServiceClient{cc}
}

func (c *containerDaemonServiceClient) GetContainers(ctx context.Context, in *GetContainersRequest, opts ...grpc.CallOption) (*GetContainersResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetContainersResponse)
	err := c.cc.Invoke(ctx, ContainerDaemonService_GetContainers_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *containerDaemonServiceClient) CreateContainer(ctx context.Context, in *CreateContainerRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[CreateContainerResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &ContainerDaemonService_ServiceDesc.Streams[0], ContainerDaemonService_CreateContainer_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[CreateContainerRequest, CreateContainerResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ContainerDaemonService_CreateContainerClient = grpc.ServerStreamingClient[CreateContainerResponse]

// ContainerDaemonServiceServer is the server API for ContainerDaemonService service.
// All implementations must embed UnimplementedContainerDaemonServiceServer
// for forward compatibility.
type ContainerDaemonServiceServer interface {
	GetContainers(context.Context, *GetContainersRequest) (*GetContainersResponse, error)
	CreateContainer(*CreateContainerRequest, grpc.ServerStreamingServer[CreateContainerResponse]) error
	mustEmbedUnimplementedContainerDaemonServiceServer()
}

// UnimplementedContainerDaemonServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedContainerDaemonServiceServer struct{}

func (UnimplementedContainerDaemonServiceServer) GetContainers(context.Context, *GetContainersRequest) (*GetContainersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetContainers not implemented")
}
func (UnimplementedContainerDaemonServiceServer) CreateContainer(*CreateContainerRequest, grpc.ServerStreamingServer[CreateContainerResponse]) error {
	return status.Errorf(codes.Unimplemented, "method CreateContainer not implemented")
}
func (UnimplementedContainerDaemonServiceServer) mustEmbedUnimplementedContainerDaemonServiceServer() {
}
func (UnimplementedContainerDaemonServiceServer) testEmbeddedByValue() {}

// UnsafeContainerDaemonServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ContainerDaemonServiceServer will
// result in compilation errors.
type UnsafeContainerDaemonServiceServer interface {
	mustEmbedUnimplementedContainerDaemonServiceServer()
}

func RegisterContainerDaemonServiceServer(s grpc.ServiceRegistrar, srv ContainerDaemonServiceServer) {
	// If the following call pancis, it indicates UnimplementedContainerDaemonServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ContainerDaemonService_ServiceDesc, srv)
}

func _ContainerDaemonService_GetContainers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetContainersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContainerDaemonServiceServer).GetContainers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContainerDaemonService_GetContainers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContainerDaemonServiceServer).GetContainers(ctx, req.(*GetContainersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContainerDaemonService_CreateContainer_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(CreateContainerRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ContainerDaemonServiceServer).CreateContainer(m, &grpc.GenericServerStream[CreateContainerRequest, CreateContainerResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ContainerDaemonService_CreateContainerServer = grpc.ServerStreamingServer[CreateContainerResponse]

// ContainerDaemonService_ServiceDesc is the grpc.ServiceDesc for ContainerDaemonService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ContainerDaemonService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "daemon.ContainerDaemonService",
	HandlerType: (*ContainerDaemonServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetContainers",
			Handler:    _ContainerDaemonService_GetContainers_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "CreateContainer",
			Handler:       _ContainerDaemonService_CreateContainer_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "google/rpc/service.proto",
}
