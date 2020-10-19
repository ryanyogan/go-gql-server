package handlers

import (
	"github.com/99designs/gqlgen/handler"
	"github.com/gin-gonic/gin"
	"yogan.dev/go-gql-server/internal/gql"
	"yogan.dev/go-gql-server/internal/gql/resolvers"
	"yogan.dev/go-gql-server/internal/orm"
)

func GraphqlHandler(orm *orm.ORM) gin.HandlerFunc {
	c := gql.Config{
		Resolvers: &resolvers.Resolver{
			ORM: orm,
		},
	}

	h := handler.GraphQL(gql.NewExecutableSchema(c))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func PlaygroundHandler(path string) gin.HandlerFunc {
	h := handler.Playground("Go Go Gadget", path)
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
