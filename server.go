package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

	"github.com/Hack-Hack-geek-Vol10/graph-gateway/google"
	"github.com/Hack-Hack-geek-Vol10/graph-gateway/graph"
	"github.com/Hack-Hack-geek-Vol10/graph-gateway/internal"
	"github.com/Hack-Hack-geek-Vol10/graph-gateway/middleware"
)

const defaultPort = "8080"

func init() {
	google.GetGoogleJWKs()
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(internal.NewExecutableSchema(internal.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", middleware.FirebaseAuth(srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
