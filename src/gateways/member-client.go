package gateways

import (
	"context"

	"github.com/newrelic/go-agent/v3/newrelic"
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
	DeleteAllProjectMember(ctx context.Context, arg *memberService.DeleteAllMemberRequest) (*memberService.DeleteMemberResponse, error)
}

func NewMemberClient(client memberService.MemberClient) MemberClient {
	return &memberClient{
		client: client,
	}
}

func (m *memberClient) CreateProjectMember(ctx context.Context, arg *memberService.MemberRequest) (*memberService.MemberResponse, error) {
	defer newrelic.FromContext(ctx).StartSegment("CreateProjectMember-client").End()
	return m.client.CreateMember(ctx, arg)
}

func (m *memberClient) GetProjectMembers(ctx context.Context, projectId string) ([]*memberService.MemberResponse, error) {
	defer newrelic.FromContext(ctx).StartSegment("GetProjectMembers-client").End()
	result, err := m.client.GetMembers(ctx, &memberService.GetMembersRequest{ProjectId: projectId})
	if err != nil {
		return nil, err
	}
	return result.Members, nil
}

func (m *memberClient) UpdateProjectMember(ctx context.Context, arg *memberService.MemberRequest) (*memberService.MemberResponse, error) {
	defer newrelic.FromContext(ctx).StartSegment("UpdateProjectMember-client").End()
	return m.client.UpdateAuthority(ctx, arg)
}

func (m *memberClient) DeleteProjectMember(ctx context.Context, arg *memberService.DeleteMemberRequest) (*memberService.DeleteMemberResponse, error) {
	defer newrelic.FromContext(ctx).StartSegment("DeleteProjectMember-client").End()
	return m.client.DeleteMember(ctx, arg)
}

func (m *memberClient) DeleteAllProjectMember(ctx context.Context, arg *memberService.DeleteAllMemberRequest) (*memberService.DeleteMemberResponse, error) {
	defer newrelic.FromContext(ctx).StartSegment("DeleteAllProjectMember-client").End()
	return m.client.DeleteAllMembers(ctx, arg)
}
