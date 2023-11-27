package gateways

import (
	"context"

	tokenService "github.com/Hack-Hack-geek-Vol10/graph-gateway/pkg/grpc/token-service"
)

type tokenClient struct {
	client tokenService.TokenServiceClient
}

type TokenClient interface {
	CreateToken(ctx context.Context, arg *tokenService.CreateTokenRequest) (*tokenService.CreateTokenResponse, error)
	VerifyToken(ctx context.Context, arg *tokenService.ValidateTokenRequest) (*tokenService.ValidateTokenResponse, error)
}

func NewTokenClient(client tokenService.TokenServiceClient) TokenClient {
	return &tokenClient{
		client: client,
	}
}

func (t *tokenClient) CreateToken(ctx context.Context, arg *tokenService.CreateTokenRequest) (*tokenService.CreateTokenResponse, error) {
	return t.client.CreateToken(ctx, arg)
}

func (t *tokenClient) VerifyToken(ctx context.Context, arg *tokenService.ValidateTokenRequest) (*tokenService.ValidateTokenResponse, error) {
	return t.client.ValidateToken(ctx, arg)
}
