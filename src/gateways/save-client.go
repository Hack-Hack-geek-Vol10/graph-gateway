package gateways

import (
	"context"

	saveService "github.com/schema-creator/graph-gateway/pkg/grpc/save-service/v1"
)

type saveClient struct {
	client saveService.SaveClient
}

type SaveClient interface {

}

func NewSaveClient(client saveService.SaveClient) SaveClient {
	return &saveClient{
		client: client,
	}
}

func (s *saveClient) CreateSave(ctx context.Context, arg *saveService.CreateSaveParams) (*saveService.SaveDetail, error) {
	return s.client.CreateSave(ctx, arg)
}

func (s *saveClient) 