package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/schema-creator/graph-gateway/cmd/config"
	"github.com/schema-creator/graph-gateway/src/internal"
	"github.com/schema-creator/graph-gateway/src/middleware"
)

func Server() {
	g := gin.Default()

	resolber, err := NewResolver()
	if err != nil {
		log.Fatal(err)
	}

	c := cors.DefaultConfig()
	c.AllowAllOrigins = true
	c.AllowHeaders = append(c.AllowHeaders, middleware.AuthTokenKey)
	g.Use(cors.New(c), middleware.FirebaseAuth(), gin.Recovery())
	g.POST("/query", gin.WrapH(handler.NewDefaultServer(internal.NewExecutableSchema(internal.Config{Resolvers: resolber}))))

	srv := &http.Server{
		Addr:    ":" + config.Config.Server.Port,
		Handler: g,
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
