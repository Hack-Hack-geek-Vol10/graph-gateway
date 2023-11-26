package services

import (
	"context"

	grpc "github.com/Hack-Hack-geek-Vol10/graph-gateway/pkg/grpc/v1"
	"github.com/Hack-Hack-geek-Vol10/graph-gateway/pkg/token"
	"github.com/Hack-Hack-geek-Vol10/graph-gateway/src/gateways"
	"github.com/Hack-Hack-geek-Vol10/graph-gateway/src/graph/model"
	"github.com/Hack-Hack-geek-Vol10/graph-gateway/src/middleware"
)

type memberService struct {
	client gateways.GRPCClient
	maker  token.Maker
}

type MemberService interface {
	CreateMember(ctx context.Context, token string) (*model.ProjectMember, error)
	GetMembers(ctx context.Context, projectID string) ([]*model.ProjectMember, error)
	UpdateMember(ctx context.Context, projectID, userID string, authority *model.Auth) (*model.ProjectMember, error)
	DeleteMember(ctx context.Context, projectID, userID string) (*string, error)
}

func NewMemberService(client gateways.GRPCClient) MemberService {
	return &memberService{
		client: client,
	}
}

func (m *memberService) CreateMember(ctx context.Context, token string) (*model.ProjectMember, error) {
	payload := ctx.Value(middleware.TokenKey{}).(*middleware.CustomClaims)

	claims, err := m.maker.VerifyToken(token)
	if err != nil {
		return nil, err
	}

	result, err := m.client.CreateProjectMember(ctx, &grpc.MemberRequest{
		UserId:    payload.UserId,
		ProjectId: claims.ProjectID,
		Authority: claims.Authority.String(),
	})
	if err != nil {
		return nil, err
	}

	return &model.ProjectMember{
		ProjectID: result.ProjectId,
		UserID:    result.UserId,
		Authority: model.Auth(result.Authority),
	}, nil
}

func (m *memberService) GetMembers(ctx context.Context, projectID string) ([]*model.ProjectMember, error) {
	result, err := m.client.GetProjectMembers(ctx, projectID)
	if err != nil {
		return nil, err
	}

	members := make([]*model.ProjectMember, len(result))
	for i, member := range result {
		members[i] = &model.ProjectMember{
			ProjectID: member.ProjectId,
			UserID:    member.UserId,
			Authority: model.Auth(member.Authority),
		}
	}

	return members, nil
}

func (m *memberService) UpdateMember(ctx context.Context, projectID, userID string, authority *model.Auth) (*model.ProjectMember, error) {
	result, err := m.client.UpdateProjectMember(ctx, &grpc.MemberRequest{
		ProjectId: projectID,
		UserId:    userID,
		Authority: authority.String(),
	})
	if err != nil {
		return nil, err
	}

	return &model.ProjectMember{
		ProjectID: result.ProjectId,
		UserID:    result.UserId,
		Authority: model.Auth(result.Authority),
	}, nil
}

func (m *memberService) DeleteMember(ctx context.Context, projectID, userID string) (*string, error) {
	result, err := m.client.DeleteProjectMember(ctx, &grpc.DeleteMemberRequest{
		ProjectId: projectID,
		UserId:    userID,
	})
	if err != nil {
		return nil, err
	}
	return &result.Message, nil
}
