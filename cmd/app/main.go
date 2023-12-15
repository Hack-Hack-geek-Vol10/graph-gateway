package main

import (
	"log"

	"github.com/schema-creator/graph-gateway/cmd/config"
	"github.com/schema-creator/graph-gateway/pkg/google"
	"github.com/schema-creator/graph-gateway/src/infra/server"
)

func init() {
	google.GetGoogleJWKs()
	config.LoadEnv()
}

func main() {
	log.Println("start graph-gateway server")
	server.Server()
}
