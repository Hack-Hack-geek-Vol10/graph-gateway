package services

import (
	"context"

	member "github.com/Hack-Hack-geek-Vol10/graph-gateway/pkg/grpc/member-service/v1"
	v1 "github.com/Hack-Hack-geek-Vol10/graph-gateway/pkg/grpc/token-service/v1"
	"github.com/Hack-Hack-geek-Vol10/graph-gateway/src/gateways"
	"github.com/Hack-Hack-geek-Vol10/graph-gateway/src/graph/model"
	"github.com/Hack-Hack-geek-Vol10/graph-gateway/src/middleware"
)

type memberService struct {
	memberClient gateways.MemberClient
	tokenClient  gateways.TokenClient
}

type MemberService interface {
	CreateMember(ctx context.Context, token string) (*model.ProjectMember, error)
	GetMembers(ctx context.Context, projectID string) ([]*model.ProjectMember, error)
	UpdateMember(ctx context.Context, projectID, userID string, authority *model.Auth) (*model.ProjectMember, error)
	DeleteMember(ctx context.Context, projectID, userID string) (*string, error)
}

func NewMemberService(memberClient gateways.MemberClient, tokenClient gateways.TokenClient) MemberService {
	return &memberService{
		memberClient: memberClient,
		tokenClient:  tokenClient,
	}
}

func (m *memberService) CreateMember(ctx context.Context, token string) (*model.ProjectMember, error) {
	payload := ctx.Value(middleware.TokenKey{}).(*middleware.CustomClaims)

	response, err := m.tokenClient.VerifyToken(ctx, &v1.VerifyTokenRequest{
		Token: token,
	})
	if err != nil {
		return nil, err
	}

	result, err := m.memberClient.CreateProjectMember(ctx, &member.MemberRequest{
		UserId:    payload.UserId,
		ProjectId: response.ProjectId,
		Authority: response.Authority.String(),
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
	result, err := m.memberClient.GetProjectMembers(ctx, projectID)
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
	result, err := m.memberClient.UpdateProjectMember(ctx, &member.MemberRequest{
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
	result, err := m.memberClient.DeleteProjectMember(ctx, &member.DeleteMemberRequest{
		ProjectId: projectID,
		UserId:    userID,
	})
	if err != nil {
		return nil, err
	}
	return &result.Message, nil
}
