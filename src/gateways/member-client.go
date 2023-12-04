package gateways

import (
	"context"

	memberService "github.com/schema-creator/graph-gateway/pkg/grpc/member-service/v1"
)

type memberClient struct {
	client memberService.MemberClient
}

type MemberClient interface {
	CreateProjectMember(ctx context.Context, arg *memberService.MemberRequest) (*memberService.MemberResponse, error)
	GetProjectMembers(ctx context.Context, projectId string) ([]*memberService.MemberResponse, error)
	UpdateProjectMember(ctx context.Context, arg *memberService.MemberRequest) (*memberService.MemberResponse, error)
	DeleteProjectMember(ctx context.Context, arg *memberService.DeleteMemberRequest) (*memberService.DeleteMemberResponse, error)
}

func NewMemberClient(client memberService.MemberClient) MemberClient {
	return &memberClient{
		client: client,
	}
}

func (m *memberClient) CreateProjectMember(ctx context.Context, arg *memberService.MemberRequest) (*memberService.MemberResponse, error) {
	return m.client.CreateMember(ctx, arg)
}

func (m *memberClient) GetProjectMembers(ctx context.Context, projectId string) ([]*memberService.MemberResponse, error) {
	result, err := m.client.GetMembers(ctx, &memberService.GetMembersRequest{ProjectId: projectId})
	if err != nil {
		return nil, err
	}
	return result.Members, nil
}

func (m *memberClient) UpdateProjectMember(ctx context.Context, arg *memberService.MemberRequest) (*memberService.MemberResponse, error) {
	return m.client.UpdateAuthority(ctx, arg)
}

func (m *memberClient) DeleteProjectMember(ctx context.Context, arg *memberService.DeleteMemberRequest) (*memberService.DeleteMemberResponse, error) {
	return m.client.DeleteMember(ctx, arg)
}
