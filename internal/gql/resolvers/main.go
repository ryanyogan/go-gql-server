package resolvers

import (
	"yogan.dev/go-gql-server/internal/gql"
	"yogan.dev/go-gql-server/internal/orm"
)

// Resolver is a struct that can used to pass on props use in resolvers, such as a DB
type Resolver struct {
	ORM *orm.ORM
}

// Mutation exposes mutation methods
func (r *Resolver) Mutation() gql.MutationResolver {
	return &mutationResolver{r}
}

// Query exposes query methods
func (r *Resolver) Query() gql.QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct {
	*Resolver
}

type queryResolver struct {
	*Resolver
}
