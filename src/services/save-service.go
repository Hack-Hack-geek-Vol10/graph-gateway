package services

import (
	"context"

	save "github.com/schema-creator/graph-gateway/pkg/grpc/save-service/v1"
	"github.com/schema-creator/graph-gateway/src/gateways"
	"github.com/schema-creator/graph-gateway/src/graph/model"
)

type saveService struct {
	saveClient gateways.SaveClient
}

type SaveService interface {
	CreateSave(ctx context.Context, arg *model.CreateSaveInput) error
	GetSave(ctx context.Context, projectID string) (*model.Save, error)
}

func NewSaveService() SaveService {
	return &saveService{}
}

// Subscriptions
func (s *saveService) CreateSave(ctx context.Context, arg *model.CreateSaveInput) error {
	s.Mutex.Lock()

	_, err := s.saveClient.CreateSave(ctx, &save.CreateSaveRequest{
		ProjectId: arg.ProjectID,
		Editor:    arg.Editor,
		Object:    arg.Object,
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *saveService) GetSave(ctx context.Context, projectID string) (*model.Save, error) {
	result, err := s.saveClient.GetSave(ctx, &save.GetSaveRequest{
		ProjectId: projectID,
	})
	if err != nil {
		return nil, err
	}

	return &model.Save{
		SaveID: result.SaveId,
		Editor: result.Editor,
		Object: result.Object,
	}, nil
}
