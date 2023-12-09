package graph

import (
	"sync"

	"github.com/schema-creator/graph-gateway/src/graph/model"
	"github.com/schema-creator/graph-gateway/src/services"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	UserService     services.UserService
	ProjectService  services.ProjectService
	MemberService   services.MemberService
	SaveService     services.SaveService
	ChannelsByMatID map[int64][]chan<- *model.CreateSaveInput
	Mutex           sync.Mutex
}

// func newResolver(saveClient *saveClient.Client) *Resolver {}
