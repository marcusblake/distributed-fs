// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package chunkserver

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

// ChunkserverClient is the client API for Chunkserver service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChunkserverClient interface {
	ServeClientFileRequest(ctx context.Context, in *FileIORequest, opts ...grpc.CallOption) (*FileIOResponse, error)
}

type chunkserverClient struct {
	cc grpc.ClientConnInterface
}

func NewChunkserverClient(cc grpc.ClientConnInterface) ChunkserverClient {
	return &chunkserverClient{cc}
}

func (c *chunkserverClient) ServeClientFileRequest(ctx context.Context, in *FileIORequest, opts ...grpc.CallOption) (*FileIOResponse, error) {
	out := new(FileIOResponse)
	err := c.cc.Invoke(ctx, "/chunkserver.Chunkserver/ServeClientFileRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChunkserverServer is the server API for Chunkserver service.
// All implementations must embed UnimplementedChunkserverServer
// for forward compatibility
type ChunkserverServer interface {
	ServeClientFileRequest(context.Context, *FileIORequest) (*FileIOResponse, error)
	mustEmbedUnimplementedChunkserverServer()
}

// UnimplementedChunkserverServer must be embedded to have forward compatible implementations.
type UnimplementedChunkserverServer struct {
}

func (UnimplementedChunkserverServer) ServeClientFileRequest(context.Context, *FileIORequest) (*FileIOResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ServeClientFileRequest not implemented")
}
func (UnimplementedChunkserverServer) mustEmbedUnimplementedChunkserverServer() {}

// UnsafeChunkserverServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChunkserverServer will
// result in compilation errors.
type UnsafeChunkserverServer interface {
	mustEmbedUnimplementedChunkserverServer()
}

func RegisterChunkserverServer(s grpc.ServiceRegistrar, srv ChunkserverServer) {
	s.RegisterService(&Chunkserver_ServiceDesc, srv)
}

func _Chunkserver_ServeClientFileRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileIORequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChunkserverServer).ServeClientFileRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chunkserver.Chunkserver/ServeClientFileRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChunkserverServer).ServeClientFileRequest(ctx, req.(*FileIORequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Chunkserver_ServiceDesc is the grpc.ServiceDesc for Chunkserver service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Chunkserver_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chunkserver.Chunkserver",
	HandlerType: (*ChunkserverServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ServeClientFileRequest",
			Handler:    _Chunkserver_ServeClientFileRequest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/chunkserver/chunkserver.proto",
}
