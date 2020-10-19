package server

import (
	"log"

	"github.com/gin-gonic/gin"
	"yogan.dev/go-gql-server/internal/handlers"
	"yogan.dev/go-gql-server/pkg/utils"
)

var (
	host string
	port string
)

func init() {
	host = utils.MustGet("GQL_SERVER_HOST")
	port = utils.MustGet("GQL_SERVER_PORT")
}

// Run the web server
func Run() {
	r := gin.Default()
	r.GET("/ping", handlers.Ping())

	log.Println("Running @ http://" + host + ":" + port)
	log.Fatalln(r.Run(host + ":" + port))
}
