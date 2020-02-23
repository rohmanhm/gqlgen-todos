// Package resolver provides a schema resolvers.
// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
package resolver

import (
	"github.com/rohmanhm/gqlgen-todos/graph/generated"
	"github.com/rohmanhm/gqlgen-todos/graph/model"
)

// Root is resolver.
type Root struct {
	todos []*model.Todo
	users []*model.User
}

var _ generated.ResolverRoot = &Root{}

// NewRoot Root resolver.
func NewRoot() *Root {
	return &Root{}
}
