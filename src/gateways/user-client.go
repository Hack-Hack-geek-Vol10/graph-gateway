package gateways

import (
	"context"

	userService "github.com/Hack-Hack-geek-Vol10/graph-gateway/pkg/grpc/user-service/v1"
)

type userClient struct {
	client userService.UserServiceClient
}

type UserClient interface {
	CreateUser(ctx context.Context, arg *userService.CreateUserParams) (*userService.UserDetail, error)
	GetOneUser(ctx context.Context, userId string) (*userService.UserDetail, error)
}

func NewUserClien(client userService.UserServiceClient) UserClient {
	return &userClient{
		client: client,
	}
}

func (u *userClient) CreateUser(ctx context.Context, arg *userService.CreateUserParams) (*userService.UserDetail, error) {
	return u.client.CreateUser(ctx, arg)
}

func (u *userClient) GetOneUser(ctx context.Context, userId string) (*userService.UserDetail, error) {
	return u.client.GetUser(ctx, &userService.GetUserParams{UserId: userId})
}
