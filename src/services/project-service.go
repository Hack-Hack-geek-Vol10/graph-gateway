package services

import (
	"context"
	"fmt"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/newrelic/go-agent/v3/newrelic"
	member "github.com/schema-creator/graph-gateway/pkg/grpc/member-service/v1"
	project "github.com/schema-creator/graph-gateway/pkg/grpc/project-service/v1"
	token "github.com/schema-creator/graph-gateway/pkg/grpc/token-service/v1"
	"github.com/schema-creator/graph-gateway/src/gateways"
	"github.com/schema-creator/graph-gateway/src/graph/model"
	"github.com/schema-creator/graph-gateway/src/middleware"
)

type projectService struct {
	projectClient gateways.ProjectClient
	memberClient  gateways.MemberClient
	imageClient   gateways.ImageClient
	tokenClient   gateways.TokenClient
}

type ProjectService interface {
	CreateProject(ctx context.Context, txn *newrelic.Transaction, title string) (*model.Project, error)
	GetProject(ctx context.Context, txn *newrelic.Transaction, projectID string) (*model.Project, error)
	GetProjects(ctx context.Context, txn *newrelic.Transaction, userID string) ([]*model.Project, error)
	UpdateProject(ctx context.Context, txn *newrelic.Transaction, projectID, title string, image *graphql.Upload) (*model.Project, error)
	DeleteProject(ctx context.Context, txn *newrelic.Transaction, projectId string) (*string, error)
	CreateInviteLink(ctx context.Context, txn *newrelic.Transaction, projectID string, authority model.Auth) (*string, error)
}

func NewProjectService(projectClient gateways.ProjectClient, memberClient gateways.MemberClient, tokenClient gateways.TokenClient, imageClient gateways.ImageClient) ProjectService {
	return &projectService{
		projectClient: projectClient,
		memberClient:  memberClient,
		tokenClient:   tokenClient,
		imageClient:   imageClient,
	}
}

func (p *projectService) CreateProject(ctx context.Context, txn *newrelic.Transaction, title string) (*model.Project, error) {
	defer txn.StartSegment("CreateProject-service").End()
	if len(title) == 0 {
		title = "untitled"
	}

	payload := ctx.Value(middleware.TokenKey{}).(*middleware.CustomClaims)

	result, err := p.projectClient.CreateProject(ctx, txn, &project.CreateProjectRequest{
		Title:  title,
		UserId: payload.UserId,
	})

	if err != nil {
		return nil, err
	}

	_, err = p.memberClient.CreateProjectMember(ctx, txn, &member.MemberRequest{
		ProjectId: result.ProjectId,
		UserId:    payload.UserId,
		Authority: member.Auth.Enum(member.Auth_owner).String(),
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

func (p *projectService) GetProject(ctx context.Context, txn *newrelic.Transaction, projectID string) (*model.Project, error) {
	defer txn.StartSegment("GetProject-service").End()

	result, err := p.projectClient.GetOneProject(ctx, txn, projectID)
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

func (p *projectService) GetProjects(ctx context.Context, txn *newrelic.Transaction, userId string) ([]*model.Project, error) {
	defer txn.StartSegment("GetProjects-service").End()

	result, err := p.projectClient.GetProjects(ctx, txn, userId)
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

func (p *projectService) UpdateProject(ctx context.Context, txn *newrelic.Transaction, projectID, title string, image *graphql.Upload) (*model.Project, error) {
	defer txn.StartSegment("UpdateProject-service").End()

	var (
		result *project.ProjectDetails
		err    error
	)

	switch {
	case len(title) != 0 && image == nil:
		result, err = p.projectClient.UpdateProjectTitle(ctx, txn, &project.UpdateTitleRequest{
			ProjectId: projectID,
			Title:     title,
		})
	case len(title) == 0 && image != nil:
		result, err = p.imageUpload(ctx, txn, projectID, image)
	case len(title) != 0 && image != nil:
		_, err = p.projectClient.UpdateProjectTitle(ctx, txn, &project.UpdateTitleRequest{
			ProjectId: projectID,
			Title:     title,
		})
		if err != nil {
			return nil, err
		}

		result, err = p.imageUpload(ctx, txn, projectID, image)
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

func (p *projectService) imageUpload(ctx context.Context, txn *newrelic.Transaction, projectID string, image *graphql.Upload) (*project.ProjectDetails, error) {
	defer txn.StartSegment("imageUpload-service").End()

	response, err := p.imageClient.UploadImage(ctx, txn, projectID, image)
	if err != nil {
		return nil, err
	}

	return p.projectClient.UpdateProjectImage(ctx, txn, projectID, response.Path)
}

func (p *projectService) DeleteProject(ctx context.Context, txn *newrelic.Transaction, projectId string) (*string, error) {
	defer txn.StartSegment("DeleteProject-service").End()

	project, err := p.projectClient.DeleteProject(ctx, txn, projectId)
	if err != nil {
		return nil, err
	}

	return &project, nil
}

func (p *projectService) CreateInviteLink(ctx context.Context, txn *newrelic.Transaction, projectID string, authority model.Auth) (*string, error) {
	defer txn.StartSegment("CreateInviteLink-service").End()

	result, err := p.tokenClient.CreateToken(ctx, txn, &token.CreateTokenRequest{
		ProjectId: projectID,
		Authority: authority.String(),
	})
	if err != nil {
		return nil, err
	}
	return &result.Token, nil
}
