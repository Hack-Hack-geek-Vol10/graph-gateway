package main

import (
	"github.com/schema-creator/graph-gateway/cmd/config"
	"github.com/schema-creator/graph-gateway/pkg/google"
	"github.com/schema-creator/graph-gateway/src/server"
)

func init() {
	google.ParseGoogleJWKs("./jwks.json")
	config.LoadEnv()
}

func main() {
	server.Server()
}
