package gateways

import (
	"context"

	"github.com/newrelic/go-agent/v3/newrelic"
	projectService "github.com/schema-creator/graph-gateway/pkg/grpc/project-service/v1"
)

type projectClient struct {
	client projectService.ProjectClient
}

type ProjectClient interface {
	CreateProject(ctx context.Context, txn *newrelic.Transaction, arg *projectService.CreateProjectRequest) (*projectService.ProjectDetails, error)
	GetOneProject(ctx context.Context, txn *newrelic.Transaction, projectId string) (*projectService.ProjectDetails, error)
	GetProjects(ctx context.Context, txn *newrelic.Transaction, userId string) ([]*projectService.ProjectDetails, error)
	UpdateProjectImage(ctx context.Context, txn *newrelic.Transaction, projectID string, path string) (*projectService.ProjectDetails, error)
	UpdateProjectTitle(ctx context.Context, txn *newrelic.Transaction, arg *projectService.UpdateTitleRequest) (*projectService.ProjectDetails, error)
	DeleteProject(ctx context.Context, txn *newrelic.Transaction, projectId string) (string, error)
}

func NewProjectClient(client projectService.ProjectClient) ProjectClient {
	return &projectClient{
		client: client,
	}
}

func (p *projectClient) CreateProject(ctx context.Context, txn *newrelic.Transaction, arg *projectService.CreateProjectRequest) (*projectService.ProjectDetails, error) {
	defer txn.StartSegment("CreateProject-client").End()
	return p.client.CreateProject(ctx, arg)
}

func (p *projectClient) GetOneProject(ctx context.Context, txn *newrelic.Transaction, projectId string) (*projectService.ProjectDetails, error) {
	defer txn.StartSegment("GetOneProject-client").End()
	return p.client.GetProject(ctx, &projectService.GetProjectRequest{ProjectId: projectId})
}

func (p *projectClient) GetProjects(ctx context.Context, txn *newrelic.Transaction, userId string) ([]*projectService.ProjectDetails, error) {
	defer txn.StartSegment("GetProjects-client").End()
	result, err := p.client.ListProjects(ctx, &projectService.ListProjectsRequest{UserId: userId})
	if err != nil {
		return nil, err
	}
	return result.Projects, nil
}

func (p *projectClient) UpdateProjectImage(ctx context.Context, txn *newrelic.Transaction, projectID string, path string) (*projectService.ProjectDetails, error) {
	defer txn.StartSegment("UpdateProjectImage-client").End()
	return p.client.UpdateImage(ctx, &projectService.UpdateImageRequest{ProjectId: projectID, LastImage: path})
}

func (p *projectClient) UpdateProjectTitle(ctx context.Context, txn *newrelic.Transaction, arg *projectService.UpdateTitleRequest) (*projectService.ProjectDetails, error) {
	defer txn.StartSegment("UpdateProjectTitle-client").End()
	return p.client.UpdateTitle(ctx, arg)
}

func (p *projectClient) DeleteProject(ctx context.Context, txn *newrelic.Transaction, projectId string) (string, error) {
	defer txn.StartSegment("DeleteProject-client").End()
	result, err := p.client.DeleteProject(ctx, &projectService.DeleteProjectRequest{ProjectId: projectId})
	return result.ProjectId, err
}
