package services

import (
	"context"

	grpc "github.com/Hack-Hack-geek-Vol10/graph-gateway/pkg/grpc/v1"
	"github.com/Hack-Hack-geek-Vol10/graph-gateway/src/gateways"
	"github.com/Hack-Hack-geek-Vol10/graph-gateway/src/graph/model"
	"github.com/Hack-Hack-geek-Vol10/graph-gateway/src/middleware"
)

type userService struct {
	client gateways.GRPCClient
}

type UserService interface {
	CreateUser(ctx context.Context, name string) (*model.User, error)
	GetUser(ctx context.Context, userId string) (*model.User, error)
}

func NewUserService(client gateways.GRPCClient) UserService {
	return &userService{
		client: client,
	}
}

func (u *userService) CreateUser(ctx context.Context, name string) (*model.User, error) {
	payload := ctx.Value(middleware.TokenKey{}).(*middleware.CustomClaims)

	result, err := u.client.CreateUser(ctx, &grpc.CreateUserParams{
		UserId: payload.UserId,
		Email:  payload.Email,
		Name:   name,
	})

	if err != nil {
		return nil, err
	}

	return &model.User{
		UserID: result.UserId,
		Email:  result.Email,
		Name:   result.Name,
	}, nil
}

func (u *userService) GetUser(ctx context.Context, userId string) (*model.User, error) {
	result, err := u.client.GetOneUser(ctx, userId)
	if err != nil {
		return nil, err
	}

	return &model.User{
		UserID: result.UserId,
		Email:  result.Email,
		Name:   result.Name,
	}, nil
}
