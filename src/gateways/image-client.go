package gateways

import (
	"context"
	"io"

	"github.com/99designs/gqlgen/graphql"
	imageService "github.com/Hack-Hack-geek-Vol10/graph-gateway/pkg/grpc/image-service"
)

type imageClient struct {
	client imageService.ImageServiceClient
}

type ImageClient interface {
	UploadImage(ctx context.Context, name string, arg *graphql.Upload) (*imageService.UploadImageResponse, error)
	DeleteImage(ctx context.Context, key string) (*imageService.DeleteImageResponse, error)
}

func NewImageClient(client imageService.ImageServiceClient) ImageClient {
	return &imageClient{
		client: client,
	}
}

func (i *imageClient) UploadImage(ctx context.Context, name string, arg *graphql.Upload) (*imageService.UploadImageResponse, error) {
	data, err := io.ReadAll(arg.File)
	if err != nil {
		return nil, err
	}
	return i.client.UploadImage(ctx, &imageService.UploadImageRequest{
		Key:         name,
		ContentType: arg.ContentType,
		Data:        data,
	})
}

func (i *imageClient) DeleteImage(ctx context.Context, key string) (*imageService.DeleteImageResponse, error) {
	return i.client.DeleteImage(ctx, &imageService.DeleteImageRequest{Key: key})
}
