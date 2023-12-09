package services

import (
	"context"
	"fmt"

	"github.com/schema-creator/graph-gateway/src/gateways"
	"github.com/schema-creator/graph-gateway/src/graph/model"
	"github.com/schema-creator/graph-gateway/src/middleware"
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
	payload := ctx.Value(middleware.TokenKey{}).(*middleware.CustomClaims)
	if payload == nil {
		return fmt.Errorf("no token found")
	}

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
	result, err := s.saveClient.GetSave(ctx, projectID)
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
