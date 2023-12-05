package gateways

import (
	"context"

	"github.com/newrelic/go-agent/v3/newrelic"
	userService "github.com/schema-creator/graph-gateway/pkg/grpc/user-service/v1"
)

type userClient struct {
	client userService.UserClient
}

type UserClient interface {
	CreateUser(ctx context.Context, txn *newrelic.Transaction, arg *userService.CreateUserParams) (*userService.UserDetail, error)
	GetOneUser(ctx context.Context, txn *newrelic.Transaction, userId string) (*userService.UserDetail, error)
}

func NewUserClien(client userService.UserClient) UserClient {
	return &userClient{
		client: client,
	}
}

func (u *userClient) CreateUser(ctx context.Context, txn *newrelic.Transaction, arg *userService.CreateUserParams) (*userService.UserDetail, error) {
	defer txn.StartSegment("CreateUser-client").End()
	return u.client.CreateUser(ctx, arg)
}

func (u *userClient) GetOneUser(ctx context.Context, txn *newrelic.Transaction, userId string) (*userService.UserDetail, error) {
	defer txn.StartSegment("GetOneUser-client").End()
	return u.client.GetUser(ctx, &userService.GetUserParams{UserId: userId})
}
