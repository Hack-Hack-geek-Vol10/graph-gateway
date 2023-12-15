package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/schema-creator/graph-gateway/cmd/config"
	"github.com/schema-creator/graph-gateway/src/infra/echo"
)

func Server() {
	app, err := newrelic.NewApplication(
		newrelic.ConfigFromEnvironment(),
	)
	if err != nil {
		log.Fatal(err)
	}

	handler := echo.NewRouter(app)

	srv := &http.Server{
		Addr:    ":" + config.Config.Server.Port,
		Handler: handler,
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
