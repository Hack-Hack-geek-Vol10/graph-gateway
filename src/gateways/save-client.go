package gateways

import (
	"context"

	saveService "github.com/schema-creator/graph-gateway/pkg/grpc/save-service/v1"
	"github.com/schema-creator/graph-gateway/src/graph/model"
)

type saveClient struct {
	client saveService.SaveClient
}

type SaveClient interface {
	CreateSave(ctx context.Context, arg *saveService.CreateSaveParams) (*saveService.SaveDetail, error)
}

func NewSaveClient(client saveService.SaveClient) SaveClient {
	return &saveClient{
		client: client,
	}
}

func (s *saveClient) CreateSave(ctx context.Context, arg *saveService.CreateSaveParams) (*saveService.SaveDetail, error) {
	return s.client.CreateSave(ctx, arg)
}

func (s *saveClient) GetSave(ctx context.Context, projectId string) (*model.Save, error) {
	result, err := s.client.GetSave(ctx, &saveService.GetSaveRequest{ProjectId: projectId})
	if err != nil {
		return nil, err
	}

	return &model.Save{
		SaveID:    result.Save.SaveId,
		ProjectID: result.Save.ProjectId,
		Editor:    result.Save.Editor,
		Object:    result.Save.Object,
	}, nil
}
