package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/rs/cors"
	"github.com/schema-creator/graph-gateway/cmd/config"
	"github.com/schema-creator/graph-gateway/src/internal"
	"github.com/schema-creator/graph-gateway/src/middleware"
)

func Server() {
	mux := http.NewServeMux()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*", "http://localhost:3000"},
		AllowedHeaders:   []string{"*", "Content-Type", "Authorization"},
		AllowedMethods:   []string{"POST"},
		AllowCredentials: true,
		Debug:            false,
	})

	app, err := newrelic.NewApplication(
		newrelic.ConfigFromEnvironment(),
	)
	if err != nil {
		log.Fatal(err)
	}

	resolver, err := NewResolver(app)
	if err != nil {
		log.Fatal(err)
	}

	mux.Handle(
		newrelic.WrapHandle(
			app,
			"/query",
			middleware.AccessLog(
				middleware.Recover(
					c.Handler(
						middleware.FirebaseAuth(
							handler.NewDefaultServer(
								internal.NewExecutableSchema(
									internal.Config{
										Resolvers: resolver,
									},
								),
							),
						),
					),
				),
			),
		),
	)

	srv := &http.Server{
		Addr:    ":" + config.Config.Server.Port,
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
