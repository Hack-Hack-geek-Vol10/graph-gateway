package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/schema-creator/graph-gateway/cmd/config"
	"github.com/schema-creator/graph-gateway/src/internal"
	"github.com/schema-creator/graph-gateway/src/middleware"
)

func Server() {
	e := echo.New()

	resolber, err := NewResolver()
	if err != nil {
		log.Fatal(err)
	}

	e.Use(echoMiddleware.CORSWithConfig(echoMiddleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000", "*", ""},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
		AllowHeaders: []string{middleware.TokenKey, "Content-Type"},
	}), middleware.FirebaseAuth())
	e.POST("/query", echo.WrapHandler(handler.NewDefaultServer(internal.NewExecutableSchema(internal.Config{Resolvers: resolber}))))

	srv := &http.Server{
		Addr:    ":" + config.Config.Server.Port,
		Handler: e,
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
