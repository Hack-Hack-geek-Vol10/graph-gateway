package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Hack-Hack-geek-Vol10/graph-gateway/cmd/config"
	"github.com/Hack-Hack-geek-Vol10/graph-gateway/src/internal"
	"github.com/Hack-Hack-geek-Vol10/graph-gateway/src/middleware"
)

func Server() {
	mux := http.NewServeMux()

	resolber, err := NewResolver()
	if err != nil {
		log.Fatal(err)
	}

	mux.Handle("/", playground.Handler("GraphQL playground", "/query"))
	mux.Handle("/query", middleware.FirebaseAuth(handler.NewDefaultServer(internal.NewExecutableSchema(internal.Config{Resolvers: resolber}))))

	srv := &http.Server{
		Addr:    config.Config.Server.Port,
		Handler: mux,
	}

	go func() {
		log.Printf("start graphQL server port: %v \n playground -> http://localhost%v", config.Config.Server.Port, config.Config.Server.Port)
		if err := srv.ListenAndServe(); err != nil || err == http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit

	log.Println("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown: %s", err)
	}
}
