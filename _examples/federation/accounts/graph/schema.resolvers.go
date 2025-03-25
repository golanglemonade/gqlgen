package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.69

import (
	"context"

	"github.com/99designs/gqlgen/_examples/federation/accounts/graph/model"
)

// Me is the resolver for the me field.
func (r *queryResolver) Me(ctx context.Context) (*model.User, error) {
	return &model.User{
		ID: "1234",
		Host: &model.EmailHost{
			ID:   "4567",
			Name: "Email Host 4567",
		},
		Email:    "me@example.com",
		Username: "Me",
	}, nil
}

// Host is the resolver for the host field.
func (r *userResolver) Host(ctx context.Context, obj *model.User) (*model.EmailHost, error) {
	return r.HostForUserID(obj.ID)
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// User returns UserResolver implementation.
func (r *Resolver) User() UserResolver { return &userResolver{r} }

type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
