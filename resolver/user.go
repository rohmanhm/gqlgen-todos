package resolver

import (
	"context"

	"github.com/rohmanhm/gqlgen-todos/graph/model"
)

// CreateUser Mutation.
func (r *Mutation) CreateUser(ctx context.Context, input model.UserInput) (*model.User, error) {
	panic("not implemented")
}

// Users query resolver.
func (r *Query) Users(ctx context.Context) ([]*model.User, error) {
	return r.users, nil
}

// User query resolver.
func (r *Query) User(ctx context.Context) (*model.User, error) {
	return r.users[0], nil
}
