package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.41

import (
	"context"
	"fmt"
	"log"

	"github.com/99designs/gqlgen/graphql"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/schema-creator/graph-gateway/src/graph/model"
	"github.com/schema-creator/graph-gateway/src/internal"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, name string) (*model.User, error) {
	txn := newrelic.FromContext(ctx)
	defer txn.StartSegment("CreateUser").End()

	res, err := r.UserService.CreateUser(ctx, txn, name)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return res, nil
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
	txn := newrelic.FromContext(ctx)
	defer txn.StartSegment("CreateProject").End()

	res, err := r.ProjectService.CreateProject(ctx, txn, title)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return res, nil
}

// UpdateProject is the resolver for the updateProject field.
func (r *mutationResolver) UpdateProject(ctx context.Context, projectID string, title *string, lastImage *graphql.Upload) (*model.Project, error) {
	txn := newrelic.FromContext(ctx)
	defer txn.StartSegment("UpdateProject").End()

	res, err := r.ProjectService.UpdateProject(ctx, txn, projectID, *title, lastImage)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return res, nil
}

// DeleteProject is the resolver for the deleteProject field.
func (r *mutationResolver) DeleteProject(ctx context.Context, projectID string) (*string, error) {
	txn := newrelic.FromContext(ctx)
	defer txn.StartSegment("DeleteProject").End()

	res, err := r.ProjectService.DeleteProject(ctx, txn, projectID)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return res, nil
}

// CreateInviteLink is the resolver for the createInviteLink field.
func (r *mutationResolver) CreateInviteLink(ctx context.Context, projectID string, authority model.Auth) (*string, error) {
	txn := newrelic.FromContext(ctx)
	defer txn.StartSegment("CreateInviteLink").End()

	res, err := r.ProjectService.CreateInviteLink(ctx, txn, projectID, authority)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return res, nil
}

// CreateProjectMember is the resolver for the createProjectMember field.
func (r *mutationResolver) CreateProjectMember(ctx context.Context, token string) (*model.ProjectMember, error) {
	txn := newrelic.FromContext(ctx)
	defer txn.StartSegment("CreateProjectMember").End()

	res, err := r.MemberService.CreateMember(ctx, txn, token)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return res, nil
}

// UpdateProjectMember is the resolver for the updateProjectMember field.
func (r *mutationResolver) UpdateProjectMember(ctx context.Context, projectID string, userID string, authority *model.Auth) (*model.ProjectMember, error) {
	txn := newrelic.FromContext(ctx)
	defer txn.StartSegment("UpdateProjectMember").End()

	res, err := r.MemberService.UpdateMember(ctx, txn, projectID, userID, authority)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return res, nil
}

// DeleteProjectMember is the resolver for the deleteProjectMember field.
func (r *mutationResolver) DeleteProjectMember(ctx context.Context, projectID string, userID string) (*string, error) {
	txn := newrelic.FromContext(ctx)
	defer txn.StartSegment("DeleteProjectMember").End()

	res, err := r.MemberService.DeleteMember(ctx, txn, projectID, userID)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return res, nil
}

// CreateSave is the resolver for the createSave field.
func (r *mutationResolver) CreateSave(ctx context.Context, input model.CreateSaveInput) (*string, error) {
	param, err := r.SaveService.CreateSave(ctx, &input)
	return &param.SaveId, err
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, userID string) (*model.User, error) {
	txn := newrelic.FromContext(ctx)
	defer txn.StartSegment("GetUser").End()

	res, err := r.UserService.GetUser(ctx, txn, userID)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return res, nil
}

// Project is the resolver for the project field.
func (r *queryResolver) Project(ctx context.Context, projectID string) (*model.Project, error) {
	txn := newrelic.FromContext(ctx)
	defer txn.StartSegment("GetProject").End()

	res, err := r.ProjectService.GetProject(ctx, txn, projectID)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return res, nil
}

// Projects is the resolver for the projects field.
func (r *queryResolver) Projects(ctx context.Context, userID string) ([]*model.Project, error) {
	txn := newrelic.FromContext(ctx)
	defer txn.StartSegment("GetProjects").End()

	res, err := r.ProjectService.GetProjects(ctx, txn, userID)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return res, nil
}

// ProjectMembers is the resolver for the projectMembers field.
func (r *queryResolver) ProjectMembers(ctx context.Context, projectID string) ([]*model.ProjectMember, error) {
	txn := newrelic.FromContext(ctx)
	defer txn.StartSegment("GetProjectMembers").End()

	res, err := r.MemberService.GetMembers(ctx, txn, projectID)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return res, nil
}

// Save is the resolver for the save field.
func (r *queryResolver) Save(ctx context.Context, projectID string) (*model.Save, error) {
	return r.SaveService.GetSave(ctx, projectID)
}

// PostEditor is the resolver for the postEditor field.
func (r *subscriptionResolver) PostEditor(ctx context.Context, projectID string) (<-chan *model.Save, error) {
	return r.SaveService.WsEditor(ctx, projectID)
}

// Mutation returns internal.MutationResolver implementation.
func (r *Resolver) Mutation() internal.MutationResolver { return &mutationResolver{r} }

// Query returns internal.QueryResolver implementation.
func (r *Resolver) Query() internal.QueryResolver { return &queryResolver{r} }

// Subscription returns internal.SubscriptionResolver implementation.
func (r *Resolver) Subscription() internal.SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
