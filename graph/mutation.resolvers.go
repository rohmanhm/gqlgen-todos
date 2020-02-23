// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
package graph

import (
	"context"
	"fmt"

	"github.com/rohmanhm/gqlgen-todos/graph/generated"
	"github.com/rohmanhm/gqlgen-todos/graph/model"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.UserInput) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.TodoInput) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateTodo(ctx context.Context, input model.TodoInput) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
