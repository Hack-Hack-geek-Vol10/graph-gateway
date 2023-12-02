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
	VerifyToken(ctx context.Context, arg *tokenService.VerifyTokenRequest) (*tokenService.VerifyTokenResponse, error)
}

func NewTokenClient(client tokenService.TokenServiceClient) TokenClient {
	return &tokenClient{
		client: client,
	}
}

func (t *tokenClient) CreateToken(ctx context.Context, arg *tokenService.CreateTokenRequest) (*tokenService.CreateTokenResponse, error) {
	return t.client.CreateToken(ctx, arg)
}

func (t *tokenClient) VerifyToken(ctx context.Context, arg *tokenService.VerifyTokenRequest) (*tokenService.VerifyTokenResponse, error) {
	return t.client.VerifyToken(ctx, arg)
}
