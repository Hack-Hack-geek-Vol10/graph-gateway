package services

import (
	"context"

	user "github.com/Hack-Hack-geek-Vol10/graph-gateway/pkg/grpc/user-service"
	"github.com/Hack-Hack-geek-Vol10/graph-gateway/src/gateways"
	"github.com/Hack-Hack-geek-Vol10/graph-gateway/src/graph/model"
	"github.com/Hack-Hack-geek-Vol10/graph-gateway/src/middleware"
)

type userService struct {
	userClient gateways.UserClient
}

type UserService interface {
	CreateUser(ctx context.Context, name string) (*model.User, error)
	GetUser(ctx context.Context, userId string) (*model.User, error)
}

func NewUserService(userClient gateways.UserClient) UserService {
	return &userService{
		userClient: userClient,
	}
}

func (u *userService) CreateUser(ctx context.Context, name string) (*model.User, error) {
	payload := ctx.Value(middleware.TokenKey{}).(*middleware.CustomClaims)

	result, err := u.userClient.CreateUser(ctx, &user.CreateUserParams{
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
	result, err := u.userClient.GetOneUser(ctx, userId)
	if err != nil {
		return nil, err
	}

	return &model.User{
		UserID: result.UserId,
		Email:  result.Email,
		Name:   result.Name,
	}, nil
}
