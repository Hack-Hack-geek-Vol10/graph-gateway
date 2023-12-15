package services

import (
	"context"

	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/schema-creator/graph-gateway/pkg/firebase"
	member "github.com/schema-creator/graph-gateway/pkg/grpc/member-service/v1"
	v1 "github.com/schema-creator/graph-gateway/pkg/grpc/token-service/v1"
	"github.com/schema-creator/graph-gateway/src/gateways"
	"github.com/schema-creator/graph-gateway/src/graph/model"
	"github.com/schema-creator/graph-gateway/src/infra/auth"
)

type memberService struct {
	memberClient gateways.MemberClient
	tokenClient  gateways.TokenClient
}

type MemberService interface {
	CreateMember(ctx context.Context, txn *newrelic.Transaction, token string) (*model.ProjectMember, error)
	GetMembers(ctx context.Context, txn *newrelic.Transaction, projectID string) ([]*model.ProjectMember, error)
	UpdateMember(ctx context.Context, txn *newrelic.Transaction, projectID, userID string, authority *model.Auth) (*model.ProjectMember, error)
	DeleteMember(ctx context.Context, txn *newrelic.Transaction, projectID, userID string) (*string, error)
}

func NewMemberService(memberClient gateways.MemberClient, tokenClient gateways.TokenClient) MemberService {
	return &memberService{
		memberClient: memberClient,
		tokenClient:  tokenClient,
	}
}

func (m *memberService) CreateMember(ctx context.Context, txn *newrelic.Transaction, token string) (*model.ProjectMember, error) {
	defer txn.StartSegment("CreateMember-service").End()

	payload := ctx.Value(auth.TokenKey).(*firebase.CustomClaims)

	response, err := m.tokenClient.GetToken(ctx, txn, &v1.GetTokenRequest{
		Token: token,
	})
	if err != nil {
		return nil, err
	}

	result, err := m.memberClient.CreateProjectMember(ctx, txn, &member.MemberRequest{
		UserId:    payload.UserId,
		ProjectId: response.ProjectId,
		Authority: response.Authority,
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

func (m *memberService) GetMembers(ctx context.Context, txn *newrelic.Transaction, projectID string) ([]*model.ProjectMember, error) {
	defer txn.StartSegment("GetMembers-service").End()

	result, err := m.memberClient.GetProjectMembers(ctx, txn, projectID)
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

func (m *memberService) UpdateMember(ctx context.Context, txn *newrelic.Transaction, projectID, userID string, authority *model.Auth) (*model.ProjectMember, error) {
	defer txn.StartSegment("UpdateMember-service").End()

	result, err := m.memberClient.UpdateProjectMember(ctx, txn, &member.MemberRequest{
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

func (m *memberService) DeleteMember(ctx context.Context, txn *newrelic.Transaction, projectID, userID string) (*string, error) {
	defer txn.StartSegment("DeleteMember-service").End()

	result, err := m.memberClient.DeleteProjectMember(ctx, txn, &member.DeleteMemberRequest{
		ProjectId: projectID,
		UserId:    userID,
	})
	if err != nil {
		return nil, err
	}
	return &result.Message, nil
}
