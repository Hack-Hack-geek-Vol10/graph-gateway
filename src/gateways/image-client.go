package gateways

import (
	"context"
	"io"

	"github.com/99designs/gqlgen/graphql"
	"github.com/newrelic/go-agent/v3/newrelic"
	imageService "github.com/schema-creator/graph-gateway/pkg/grpc/image-service/v1"
)

type imageClient struct {
	client imageService.ImageClient
}

type ImageClient interface {
	UploadImage(ctx context.Context, txn *newrelic.Transaction, name string, arg *graphql.Upload) (*imageService.UploadImageResponse, error)
	DeleteImage(ctx context.Context, txn *newrelic.Transaction, key string) (*imageService.DeleteImageResponse, error)
}

func NewImageClient(client imageService.ImageClient) ImageClient {
	return &imageClient{
		client: client,
	}
}

func (i *imageClient) UploadImage(ctx context.Context, txn *newrelic.Transaction, name string, arg *graphql.Upload) (*imageService.UploadImageResponse, error) {
	defer txn.StartSegment("UploadImage-client").End()
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

func (i *imageClient) DeleteImage(ctx context.Context, txn *newrelic.Transaction, key string) (*imageService.DeleteImageResponse, error) {
	defer txn.StartSegment("DeleteImage-client").End()
	return i.client.DeleteImage(ctx, &imageService.DeleteImageRequest{Key: key})
}
