package services

import (
	"context"

	"github.com/Hack-Hack-geek-Vol10/graph-gateway/cmd/config"
	domain "github.com/Hack-Hack-geek-Vol10/graph-gateway/pkg/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func CreateUser(ctx context.Context, arg *domain.CreateUserParams) (string, error) {
	conn, err := grpc.Dial(
		config.Config.Service.UserServiceAddr,

		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		return "", err
	}

	defer conn.Close()

	client := domain.NewUserServiceClient(conn)

	res, err := client.CreateUser(ctx, arg)
	if err != nil {
		return "", err
	}

	return res.GetId(), nil
}

func GetOneUser(ctx context.Context, userId string) (*domain.UserDetail, error) {
	conn, err := grpc.Dial(
		config.Config.Service.UserServiceAddr,

		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	client := domain.NewUserServiceClient(conn)

	res, err := client.GetUser(ctx, &domain.GetUserParams{Id: userId})
	if err != nil {
		return nil, err
	}

	return res, nil
}
