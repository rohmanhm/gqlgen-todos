package resolver

import (
	"github.com/rohmanhm/gqlgen-todos/graph/generated"
)

var _ generated.MutationResolver = &Mutation{}

// Mutation resolver
type Mutation struct{ *Root }

// Mutation resolver.
func (r *Root) Mutation() generated.MutationResolver {
	return &Mutation{r}
}
