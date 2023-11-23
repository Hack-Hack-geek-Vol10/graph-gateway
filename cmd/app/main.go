package main

import (
	"github.com/Hack-Hack-geek-Vol10/graph-gateway/cmd/config"
	"github.com/Hack-Hack-geek-Vol10/graph-gateway/src/server"
	"github.com/Hack-Hack-geek-Vol10/graph-gateway/utils/google"
)

func init() {
	google.GetGoogleJWKs()
	config.LoadEnv()
}

func main() {
	server.Server()
}
