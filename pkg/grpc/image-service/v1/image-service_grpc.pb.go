// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.1
// source: image-service/v1/image-service.proto

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

// ImageServiceClient is the client API for ImageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ImageServiceClient interface {
	UploadImage(ctx context.Context, in *UploadImageRequest, opts ...grpc.CallOption) (*UploadImageResponse, error)
	DeleteImage(ctx context.Context, in *DeleteImageRequest, opts ...grpc.CallOption) (*DeleteImageResponse, error)
}

type imageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewImageServiceClient(cc grpc.ClientConnInterface) ImageServiceClient {
	return &imageServiceClient{cc}
}

func (c *imageServiceClient) UploadImage(ctx context.Context, in *UploadImageRequest, opts ...grpc.CallOption) (*UploadImageResponse, error) {
	out := new(UploadImageResponse)
	err := c.cc.Invoke(ctx, "/image_service.v1.ImageService/UploadImage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imageServiceClient) DeleteImage(ctx context.Context, in *DeleteImageRequest, opts ...grpc.CallOption) (*DeleteImageResponse, error) {
	out := new(DeleteImageResponse)
	err := c.cc.Invoke(ctx, "/image_service.v1.ImageService/DeleteImage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ImageServiceServer is the server API for ImageService service.
// All implementations must embed UnimplementedImageServiceServer
// for forward compatibility
type ImageServiceServer interface {
	UploadImage(context.Context, *UploadImageRequest) (*UploadImageResponse, error)
	DeleteImage(context.Context, *DeleteImageRequest) (*DeleteImageResponse, error)
	mustEmbedUnimplementedImageServiceServer()
}

// UnimplementedImageServiceServer must be embedded to have forward compatible implementations.
type UnimplementedImageServiceServer struct {
}

func (UnimplementedImageServiceServer) UploadImage(context.Context, *UploadImageRequest) (*UploadImageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadImage not implemented")
}
func (UnimplementedImageServiceServer) DeleteImage(context.Context, *DeleteImageRequest) (*DeleteImageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteImage not implemented")
}
func (UnimplementedImageServiceServer) mustEmbedUnimplementedImageServiceServer() {}

// UnsafeImageServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ImageServiceServer will
// result in compilation errors.
type UnsafeImageServiceServer interface {
	mustEmbedUnimplementedImageServiceServer()
}

func RegisterImageServiceServer(s grpc.ServiceRegistrar, srv ImageServiceServer) {
	s.RegisterService(&ImageService_ServiceDesc, srv)
}

func _ImageService_UploadImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageServiceServer).UploadImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/image_service.v1.ImageService/UploadImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageServiceServer).UploadImage(ctx, req.(*UploadImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImageService_DeleteImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageServiceServer).DeleteImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/image_service.v1.ImageService/DeleteImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageServiceServer).DeleteImage(ctx, req.(*DeleteImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ImageService_ServiceDesc is the grpc.ServiceDesc for ImageService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ImageService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "image_service.v1.ImageService",
	HandlerType: (*ImageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UploadImage",
			Handler:    _ImageService_UploadImage_Handler,
		},
		{
			MethodName: "DeleteImage",
			Handler:    _ImageService_DeleteImage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "image-service/v1/image-service.proto",
}
