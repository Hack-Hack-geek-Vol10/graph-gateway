package gateways

import (
	"context"

	tokenService "github.com/schema-creator/graph-gateway/pkg/grpc/token-service/v1"
)

type tokenClient struct {
	client tokenService.TokenServiceClient
}

type TokenClient interface {
	CreateToken(ctx context.Context, arg *tokenService.CreateTokenRequest) (*tokenService.CreateTokenResponse, error)
	GetToken(ctx context.Context, arg *tokenService.GetTokenRequest) (*tokenService.GetTokenResponse, error)
}

func NewTokenClient(client tokenService.TokenServiceClient) TokenClient {
	return &tokenClient{
		client: client,
	}
}

func (t *tokenClient) CreateToken(ctx context.Context, arg *tokenService.CreateTokenRequest) (*tokenService.CreateTokenResponse, error) {
	return t.client.CreateToken(ctx, arg)
}

func (t *tokenClient) GetToken(ctx context.Context, arg *tokenService.GetTokenRequest) (*tokenService.GetTokenResponse, error) {
	return t.client.GetToken(ctx, arg)
}
