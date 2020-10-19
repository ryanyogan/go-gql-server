package server

import (
	"log"

	"github.com/gin-gonic/gin"
	"yogan.dev/go-gql-server/internal/handlers"
	"yogan.dev/go-gql-server/pkg/utils"
)

var (
	host        string
	port        string
	gqlPath     string
	gqlPgPath   string
	isPgEnabled bool
)

func init() {
	host = utils.MustGet("GQL_SERVER_HOST")
	port = utils.MustGet("GQL_SERVER_PORT")
	gqlPath = utils.MustGet("GQL_SERVER_GRAPHQL_PATH")
	gqlPgPath = utils.MustGet("GQL_SERVER_GRAPHQL_PLAYGROUND_PATH")
	isPgEnabled = utils.MustGetBool("GQL_SERVER_GQL_PG_ENABLED")
}

// Run the web server
func Run() {
	endpoint := "http://" + host + ":" + port

	r := gin.Default()
	r.GET("/ping", handlers.Ping())

	if isPgEnabled {
		r.GET(gqlPgPath, handlers.PlaygroundHandler(gqlPath))
		log.Println("GraphQL Playground @ " + endpoint + gqlPgPath)
	}

	r.POST(gqlPath, handlers.GraphqlHandler())
	log.Println("GraphQL @ " + endpoint + gqlPath)

	log.Println("Running @ http://" + host + ":" + port)
	log.Fatalln(r.Run(host + ":" + port))
}
