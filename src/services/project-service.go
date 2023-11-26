package services

import (
	"context"
	"fmt"
	"time"

	"github.com/99designs/gqlgen/graphql"
	grpc "github.com/Hack-Hack-geek-Vol10/graph-gateway/pkg/grpc/v1"
	"github.com/Hack-Hack-geek-Vol10/graph-gateway/src/gateways"
	"github.com/Hack-Hack-geek-Vol10/graph-gateway/src/graph/model"
	"github.com/Hack-Hack-geek-Vol10/graph-gateway/src/middleware"
)

type projectService struct {
	client gateways.GRPCClient
}

type ProjectService interface {
	CreateProject(ctx context.Context, title string) (*model.Project, error)
	GetProject(ctx context.Context, projectID string) (*model.Project, error)
	GetProjects(ctx context.Context, userID string) ([]*model.Project, error)
	UpdateProject(ctx context.Context, projectID, title string, image *graphql.Upload) (*model.Project, error)
	DeleteProject(ctx context.Context, projectId string) (*string, error)
}

func NewProjectService(client gateways.GRPCClient) ProjectService {
	return &projectService{
		client: client,
	}
}

func (p *projectService) CreateProject(ctx context.Context, title string) (*model.Project, error) {
	if len(title) == 0 {
		title = "untitled"
	}

	payload := ctx.Value(middleware.TokenKey{}).(*middleware.CustomClaims)

	result, err := p.client.CreateProject(ctx, &grpc.CreateProjectRequest{
		Title:  title,
		UserId: payload.UserId,
	})

	if err != nil {
		return nil, err
	}

	return &model.Project{
		ProjectID: result.ProjectId,
		Title:     result.Title,
		LastImage: result.LastImage,
		UpdatedAt: time.Now().String(),
	}, nil
}

func (p *projectService) GetProject(ctx context.Context, projectID string) (*model.Project, error) {
	result, err := p.client.GetOneProject(ctx, projectID)
	if err != nil {
		return nil, err
	}

	return &model.Project{
		ProjectID: result.ProjectId,
		Title:     result.Title,
		LastImage: result.LastImage,
		UpdatedAt: time.Now().String(),
	}, nil
}

func (p *projectService) GetProjects(ctx context.Context, userId string) ([]*model.Project, error) {
	result, err := p.client.GetProjects(ctx, userId)
	if err != nil {
		return nil, err
	}

	projects := make([]*model.Project, 0, len(result))
	for _, project := range result {
		projects = append(projects, &model.Project{
			ProjectID: project.ProjectId,
			Title:     project.Title,
			LastImage: project.LastImage,
			UpdatedAt: time.Now().String(),
		})
	}

	return projects, nil
}

func (p *projectService) UpdateProject(ctx context.Context, projectID, title string, image *graphql.Upload) (*model.Project, error) {
	var (
		result *grpc.ProjectDetails
		err    error
	)

	switch {
	case len(title) != 0 && image == nil:
		result, err = p.client.UpdateProjectTitle(ctx, &grpc.UpdateTitleRequest{
			ProjectId: projectID,
			Title:     title,
		})
	case len(title) == 0 && image != nil:
		result, err = p.client.UpdateProjectImage(ctx, projectID, image)
	case len(title) != 0 && image != nil:
		_, err = p.client.UpdateProjectTitle(ctx, &grpc.UpdateTitleRequest{
			ProjectId: projectID,
			Title:     title,
		})
		if err != nil {
			return nil, err
		}

		result, err = p.client.UpdateProjectImage(ctx, projectID, image)
	default:
		return nil, fmt.Errorf("title or image must be specified")
	}

	if err != nil {
		return nil, err
	}

	return &model.Project{
		ProjectID: result.ProjectId,
		Title:     result.Title,
		LastImage: result.LastImage,
		UpdatedAt: time.Now().String(),
	}, nil
}

func (p *projectService) DeleteProject(ctx context.Context, projectId string) (*string, error) {
	project, err := p.client.DeleteProject(ctx, projectId)
	if err != nil {
		return nil, err
	}

	return &project, nil
}
