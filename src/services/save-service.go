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
	WsEditor(ctx context.Context, id string) (<-chan *model.Save, error)
}

func NewSaveService(saveClient gateways.SaveClient) SaveService {
	return &saveService{
		saveClient:          saveClient,
		ChannelsByProjectID: make(map[string][]chan<- *model.Save),
		Mutex:               sync.Mutex{},
	}
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
	// チャネルに送信
	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	for _, ch := range s.ChannelsByProjectID[arg.ProjectID] {
		ch <- &model.Save{
			SaveID: saveID.String(),
			Editor: arg.Editor,
			Object: arg.Object,
		}
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

	return &model.Save{
		SaveID: result.SaveId,
		Editor: result.Editor,
		Object: result.Object,
	}, nil
}

func (s *saveService) WsEditor(ctx context.Context, pid string) (<-chan *model.Save, error) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()

	ch := make(chan *model.Save, 1)
	s.ChannelsByProjectID[pid] = append(s.ChannelsByProjectID[pid], ch)

	// コネクション終了時にチャネルを削除
	go func() {
		<-ctx.Done()
		s.Mutex.Lock()
		defer s.Mutex.Unlock()
		for i, c := range s.ChannelsByProjectID[pid] {
			if c == ch {
				s.ChannelsByProjectID[pid] = append(s.ChannelsByProjectID[pid][:i], s.ChannelsByProjectID[pid][i+1:]...)
				break
			}
		}
	}()

	return ch, nil
}
