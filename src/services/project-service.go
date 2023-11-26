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

func (p *projectService) DeleteProject(ctx context.Context, projectId string) (*model.ProjectMember, error) {
	result, err := p.client.DeleteProject(ctx, projectId)
	if err != nil {
		return nil, err
	}

	return &model.ProjectMember{
		ProjectID: result,
	}, nil
}
