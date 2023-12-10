package services

import (
	"context"
	"sync"

	save "github.com/schema-creator/graph-gateway/pkg/grpc/save-service/v1"
	"github.com/schema-creator/graph-gateway/src/gateways"
	"github.com/schema-creator/graph-gateway/src/graph/model"
)

type saveService struct {
	saveClient          gateways.SaveClient
	ChannelsByProjectID map[string][]chan<- *model.Save
	Mutex               sync.Mutex
}

type SaveService interface {
	CreateSave(ctx context.Context, arg *model.CreateSaveInput) (*save.CreateSaveResponse, error)
	GetSave(ctx context.Context, projectID string) (*model.Save, error)
}

func NewSaveService() SaveService {
	return &saveService{}
}

// Subscriptions
func (s *saveService) CreateSave(ctx context.Context, arg *model.CreateSaveInput) (*save.CreateSaveResponse, error) {
	saveID, err := s.saveClient.CreateSave(ctx, &save.CreateSaveRequest{
		ProjectId: arg.ProjectID,
		Editor:    arg.Editor,
		Object:    arg.Object,
	})
	if err != nil {
		return nil, err
	}

	return saveID, err
}

func (s *saveService) GetSave(ctx context.Context, projectID string) (*model.Save, error) {
	result, err := s.saveClient.GetSave(ctx, &save.GetSaveRequest{
		ProjectId: projectID,
	})
	if err != nil {
		return nil, err
	}

	// チャネルに送信
	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	for _, ch := range s.ChannelsByProjectID[projectID] {
		ch <- &model.Save{
			SaveID: result.SaveId,
			Editor: result.Editor,
			Object: result.Object,
		}
	}

	return &model.Save{
		SaveID: result.SaveId,
		Editor: result.Editor,
		Object: result.Object,
	}, nil
}

func (s *saveService) WsEditor(ctx context.Context, id string) (<-chan *model.Save, error) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()

	ch := make(chan *model.Save, 1)
	s.ChannelsByProjectID[id] = append(s.ChannelsByProjectID[id], ch)

	// コネクション終了時にチャネルを削除
	go func() {
		<-ctx.Done()
		s.Mutex.Lock()
		defer s.Mutex.Unlock()
		for i, c := range s.ChannelsByProjectID[id] {
			if c == ch {
				s.ChannelsByProjectID[id] = append(s.ChannelsByProjectID[id][:i], s.ChannelsByProjectID[id][i+1:]...)
				break
			}
		}
	}()

	return ch, nil
}
