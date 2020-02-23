package resolver

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/rohmanhm/gqlgen-todos/graph/model"
)

// CreateTodo mutation.
func (r *Mutation) CreateTodo(ctx context.Context, input model.TodoInput) (*model.Todo, error) {
	todo := &model.Todo{
		Text: input.Text,
		ID:   fmt.Sprintf("T%d", rand.Int()),
		User: &model.User{ID: input.UserID, Name: "user " + input.UserID},
	}
	r.todos = append(r.todos, todo)
	return todo, nil
}

// UpdateTodo Mutation.
func (r *Mutation) UpdateTodo(ctx context.Context, input model.TodoInput) (*model.Todo, error) {
	panic("not implemented")
}

// Todos query resolver.
func (r *Query) Todos(ctx context.Context) ([]*model.Todo, error) {
	return r.todos, nil
}
