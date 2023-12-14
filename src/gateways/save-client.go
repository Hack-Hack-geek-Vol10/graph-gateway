package gateways

import (
	"context"

	saveService "github.com/schema-creator/graph-gateway/pkg/grpc/save-service/v1"
)

type saveClient struct {
	client saveService.SaveClient
}

type SaveClient interface {
	CreateSave(ctx context.Context, arg *saveService.CreateSaveRequest) (*saveService.CreateSaveResponse, error)
	GetSave(ctx context.Context, arg *saveService.GetSaveRequest) (*saveService.GetSaveResponse, error)
}

func NewSaveClient(client saveService.SaveClient) SaveClient {
	return &saveClient{
		client: client,
	}
}

func (s *saveClient) CreateSave(ctx context.Context, arg *saveService.CreateSaveRequest) (*saveService.CreateSaveResponse, error) {
	return s.client.CreateSave(ctx, arg)
}

func (s *saveClient) GetSave(ctx context.Context, arg *saveService.GetSaveRequest) (*saveService.GetSaveResponse, error) {
	return s.client.GetSave(ctx, arg)
}
