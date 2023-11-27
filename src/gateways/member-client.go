package gateways

import (
	"context"

	memberService "github.com/Hack-Hack-geek-Vol10/graph-gateway/pkg/grpc/member-service"
)

type memberClient struct {
	client memberService.MemberServiceClient
}

type MemberClient interface {
	CreateProjectMember(ctx context.Context, arg *memberService.MemberRequest) (*memberService.Member, error)
	GetProjectMembers(ctx context.Context, projectId string) ([]*memberService.Member, error)
	UpdateProjectMember(ctx context.Context, arg *memberService.MemberRequest) (*memberService.Member, error)
	DeleteProjectMember(ctx context.Context, arg *memberService.DeleteMemberRequest) (*memberService.DeleteMemberResponse, error)
}

func NewMemberClient(client memberService.MemberServiceClient) MemberClient {
	return &memberClient{
		client: client,
	}
}

func (m *memberClient) CreateProjectMember(ctx context.Context, arg *memberService.MemberRequest) (*memberService.Member, error) {
	return m.client.AddMember(ctx, arg)
}

func (m *memberClient) GetProjectMembers(ctx context.Context, projectId string) ([]*memberService.Member, error) {
	result, err := m.client.ReadMembers(ctx, &memberService.ReadMembersRequest{ProjectId: projectId})
	if err != nil {
		return nil, err
	}
	return result.Members, nil
}

func (m *memberClient) UpdateProjectMember(ctx context.Context, arg *memberService.MemberRequest) (*memberService.Member, error) {
	return m.client.UpdateAuthority(ctx, arg)
}

func (m *memberClient) DeleteProjectMember(ctx context.Context, arg *memberService.DeleteMemberRequest) (*memberService.DeleteMemberResponse, error) {
	return m.client.DeleteMember(ctx, arg)
}
