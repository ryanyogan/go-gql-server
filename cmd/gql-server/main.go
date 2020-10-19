package main

import (
	"log"

	"yogan.dev/go-gql-server/internal/orm"
	"yogan.dev/go-gql-server/pkg/server"
)

func main() {
	orm, err := orm.Factory()
	if err != nil {
		log.Panic(err)
	}

	server.Run(orm)
}
