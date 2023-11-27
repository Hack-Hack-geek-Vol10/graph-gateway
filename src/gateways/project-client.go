package gateways

import (
	"context"

	projectService "github.com/Hack-Hack-geek-Vol10/graph-gateway/pkg/grpc/project-service"
)

type projectClient struct {
	client projectService.ProjectServiceClient
}

type ProjectClient interface {
	CreateProject(ctx context.Context, arg *projectService.CreateProjectRequest) (*projectService.ProjectDetails, error)
	GetOneProject(ctx context.Context, projectId string) (*projectService.ProjectDetails, error)
	GetProjects(ctx context.Context, userId string) ([]*projectService.ProjectDetails, error)
	UpdateProjectImage(ctx context.Context, projectID string, path string) (*projectService.ProjectDetails, error)
	UpdateProjectTitle(ctx context.Context, arg *projectService.UpdateTitleRequest) (*projectService.ProjectDetails, error)
	DeleteProject(ctx context.Context, projectId string) (string, error)
}

func NewProjectClient(client projectService.ProjectServiceClient) ProjectClient {
	return &projectClient{
		client: client,
	}
}

func (p *projectClient) CreateProject(ctx context.Context, arg *projectService.CreateProjectRequest) (*projectService.ProjectDetails, error) {
	return p.client.CreateProject(ctx, arg)
}

func (p *projectClient) GetOneProject(ctx context.Context, projectId string) (*projectService.ProjectDetails, error) {
	return p.client.GetProject(ctx, &projectService.GetProjectRequest{ProjectId: projectId})
}

func (p *projectClient) GetProjects(ctx context.Context, userId string) ([]*projectService.ProjectDetails, error) {
	result, err := p.client.ListProjects(ctx, &projectService.ListProjectsRequest{UserId: userId})
	if err != nil {
		return nil, err
	}
	return result.Projects, nil
}

func (p *projectClient) UpdateProjectImage(ctx context.Context, projectID string, path string) (*projectService.ProjectDetails, error) {
	return p.client.UpdateImage(ctx, &projectService.UpdateImageRequest{ProjectId: projectID, LastImage: path})
}

func (p *projectClient) UpdateProjectTitle(ctx context.Context, arg *projectService.UpdateTitleRequest) (*projectService.ProjectDetails, error) {
	return p.client.UpdateTitle(ctx, arg)
}

func (p *projectClient) DeleteProject(ctx context.Context, projectId string) (string, error) {
	result, err := p.client.DeleteProject(ctx, &projectService.DeleteProjectRequest{ProjectId: projectId})
	return result.ProjectId, err
}
