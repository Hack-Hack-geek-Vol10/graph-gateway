package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.40

import (
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/Hack-Hack-geek-Vol10/graph-gateway/src/graph/model"
	"github.com/Hack-Hack-geek-Vol10/graph-gateway/src/internal"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, name string) (*model.User, error) {
	return r.user.CreateUser(ctx, name)
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, userID string, name *string) (*model.User, error) {
	return nil, fmt.Errorf("not implemented: UpdateUser - updateUser")
}

// DeleteUser is the resolver for the deleteUser field.
func (r *mutationResolver) DeleteUser(ctx context.Context, userID string) (*model.User, error) {
	return nil, fmt.Errorf("not implemented: DeleteUser - deleteUser")
}

// CreateProject is the resolver for the createProject field.
func (r *mutationResolver) CreateProject(ctx context.Context, title string) (*model.Project, error) {
	return r.project.CreateProject(ctx, title)
}

// UpdateProject is the resolver for the updateProject field.
func (r *mutationResolver) UpdateProject(ctx context.Context, projectID string, title *string, lastImage *graphql.Upload) (*model.Project, error) {
	return r.project.UpdateProject(ctx, projectID, *title, lastImage)
}

// DeleteProject is the resolver for the deleteProject field.
func (r *mutationResolver) DeleteProject(ctx context.Context, projectID string) (*string, error) {
	return r.project.DeleteProject(ctx, projectID)
}

// CreateInviteLink is the resolver for the createInviteLink field.
func (r *mutationResolver) CreateInviteLink(ctx context.Context, projectID string, authority model.Auth) (*string, error) {
	return nil, fmt.Errorf("not implemented: CreateInviteLink - createInviteLink")
}

// CreateProjectMember is the resolver for the createProjectMember field.
func (r *mutationResolver) CreateProjectMember(ctx context.Context, projectID string, userID string, authority model.Auth) (*model.ProjectMember, error) {
	return nil, nil
}

// UpdateProjectMember is the resolver for the updateProjectMember field.
func (r *mutationResolver) UpdateProjectMember(ctx context.Context, projectID string, userID string, authority *model.Auth) (*model.ProjectMember, error) {
	return nil, nil
}

// DeleteProjectMember is the resolver for the deleteProjectMember field.
func (r *mutationResolver) DeleteProjectMember(ctx context.Context, projectID string, userID string) (*model.ProjectMember, error) {
	return nil, nil
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, userID string) (*model.User, error) {
	return r.user.GetUser(ctx, userID)
}

// Project is the resolver for the project field.
func (r *queryResolver) Project(ctx context.Context, projectID string) (*model.Project, error) {
	return r.project.GetProject(ctx, projectID)
}

// Projects is the resolver for the projects field.
func (r *queryResolver) Projects(ctx context.Context, userID string) ([]*model.Project, error) {
	return r.project.GetProjects(ctx, userID)
}

// ProjectMembers is the resolver for the projectMembers field.
func (r *queryResolver) ProjectMembers(ctx context.Context, projectID string) ([]*model.ProjectMember, error) {
	return nil, nil
}

// Mutation returns internal.MutationResolver implementation.
func (r *Resolver) Mutation() internal.MutationResolver { return &mutationResolver{r} }

// Query returns internal.QueryResolver implementation.
func (r *Resolver) Query() internal.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
