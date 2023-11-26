package gateways

import (
	"context"
	"io"

	"github.com/99designs/gqlgen/graphql"
	domain "github.com/Hack-Hack-geek-Vol10/graph-gateway/pkg/grpc/v1"
	"google.golang.org/grpc"
)

type gRPCClient struct {
	conn *grpc.ClientConn
}

type GRPCClient interface {
	CreateUser(ctx context.Context, arg *domain.CreateUserParams) (*domain.UserDetail, error)
	GetOneUser(ctx context.Context, userId string) (*domain.UserDetail, error)
	CreateProject(ctx context.Context, arg *domain.CreateProjectRequest) (*domain.ProjectDetails, error)
	GetOneProject(ctx context.Context, projectId string) (*domain.ProjectDetails, error)
	GetProjects(ctx context.Context, userId string) ([]*domain.ProjectDetails, error)
	UpdateProjectImage(ctx context.Context, projectID string, image *graphql.Upload) (*domain.ProjectDetails, error)
	UpdateProjectTitle(ctx context.Context, arg *domain.UpdateTitleRequest) (*domain.ProjectDetails, error)
	DeleteProject(ctx context.Context, projectId string) (string, error)
	CreateProjectMember(ctx context.Context, arg *domain.MemberRequest) (*domain.Member, error)
	GetProjectMembers(ctx context.Context, projectId string) ([]*domain.Member, error)
	UpdateProjectMember(ctx context.Context, arg *domain.MemberRequest) (*domain.Member, error)
	DeleteProjectMember(ctx context.Context, arg *domain.DeleteMemberRequest) (*domain.DeleteMemberResponse, error)

	UploadImage(ctx context.Context, arg graphql.Upload) (*domain.UploadImageResponse, error)
	DeleteImage(ctx context.Context, key string) (*domain.DeleteImageResponse, error)
}

func NewgRPCClient(conn *grpc.ClientConn) GRPCClient {
	return &gRPCClient{
		conn: conn,
	}
}

func (u *gRPCClient) CreateUser(ctx context.Context, arg *domain.CreateUserParams) (*domain.UserDetail, error) {
	return domain.NewUserServiceClient(u.conn).CreateUser(ctx, arg)
}

func (u *gRPCClient) GetOneUser(ctx context.Context, userId string) (*domain.UserDetail, error) {
	return domain.NewUserServiceClient(u.conn).GetUser(ctx, &domain.GetUserParams{UserId: userId})
}

func (u *gRPCClient) CreateProject(ctx context.Context, arg *domain.CreateProjectRequest) (*domain.ProjectDetails, error) {
	return domain.NewProjectServiceClient(u.conn).CreateProject(ctx, arg)
}

func (u *gRPCClient) GetOneProject(ctx context.Context, projectId string) (*domain.ProjectDetails, error) {
	return domain.NewProjectServiceClient(u.conn).GetProject(ctx, &domain.GetProjectRequest{ProjectId: projectId})
}

func (u *gRPCClient) GetProjects(ctx context.Context, userId string) ([]*domain.ProjectDetails, error) {
	result, err := domain.NewProjectServiceClient(u.conn).ListProjects(ctx, &domain.ListProjectsRequest{UserId: userId})
	if err != nil {
		return nil, err
	}
	return result.Projects, nil
}

func (u *gRPCClient) UpdateProjectImage(ctx context.Context, projectID string, image *graphql.Upload) (*domain.ProjectDetails, error) {
	// TODO : upload image to storage
	imagepath := ""

	return domain.NewProjectServiceClient(u.conn).UpdateImage(ctx, &domain.UpdateImageRequest{ProjectId: projectID, LastImage: imagepath})
}

func (u *gRPCClient) UpdateProjectTitle(ctx context.Context, arg *domain.UpdateTitleRequest) (*domain.ProjectDetails, error) {
	return domain.NewProjectServiceClient(u.conn).UpdateTitle(ctx, arg)
}

func (u *gRPCClient) DeleteProject(ctx context.Context, projectId string) (string, error) {
	result, err := domain.NewProjectServiceClient(u.conn).DeleteProject(ctx, &domain.DeleteProjectRequest{ProjectId: projectId})
	return result.ProjectId, err
}

func (u *gRPCClient) CreateProjectMember(ctx context.Context, arg *domain.MemberRequest) (*domain.Member, error) {
	return domain.NewMemberServiceClient(u.conn).AddMember(ctx, arg)
}

func (u *gRPCClient) GetProjectMembers(ctx context.Context, projectId string) ([]*domain.Member, error) {
	result, err := domain.NewMemberServiceClient(u.conn).ReadMembers(ctx, &domain.ReadMembersRequest{ProjectId: projectId})
	if err != nil {
		return nil, err
	}
	return result.Members, nil
}

func (u *gRPCClient) UpdateProjectMember(ctx context.Context, arg *domain.MemberRequest) (*domain.Member, error) {
	return domain.NewMemberServiceClient(u.conn).UpdateAuthority(ctx, arg)
}

func (u *gRPCClient) DeleteProjectMember(ctx context.Context, arg *domain.DeleteMemberRequest) (*domain.DeleteMemberResponse, error) {
	return domain.NewMemberServiceClient(u.conn).DeleteMember(ctx, arg)
}

func (u *gRPCClient) UploadImage(ctx context.Context, arg graphql.Upload) (*domain.UploadImageResponse, error) {
	data, err := io.ReadAll(arg.File)
	if err != nil {
		return nil, err
	}
	return domain.NewImageServiceClient(u.conn).UploadImage(ctx, &domain.UploadImageRequest{
		Key:         arg.Filename,
		ContentType: arg.ContentType,
		Data:        data,
	})
}

func (u *gRPCClient) DeleteImage(ctx context.Context, key string) (*domain.DeleteImageResponse, error) {
	return domain.NewImageServiceClient(u.conn).DeleteImage(ctx, &domain.DeleteImageRequest{Key: key})
}
