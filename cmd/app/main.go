package main

import (
	"github.com/Hack-Hack-geek-Vol10/graph-gateway/cmd/config"
	"github.com/Hack-Hack-geek-Vol10/graph-gateway/pkg/google"
	"github.com/Hack-Hack-geek-Vol10/graph-gateway/src/server"
)

func init() {
	google.GetGoogleJWKs()
	config.LoadEnv()
}

func main() {
	server.Server()
}
