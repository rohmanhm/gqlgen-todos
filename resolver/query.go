package resolver

import (
	"github.com/rohmanhm/gqlgen-todos/graph/generated"
)

var _ generated.QueryResolver = &Query{}

// Query resolver
type Query struct{ *Root }

// Query resolver.
func (r *Root) Query() generated.QueryResolver {
	return &Query{r}
}
