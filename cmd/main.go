package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/defry256/pokemon-helper/config"
	"github.com/defry256/pokemon-helper/internal/httpserver"
	"github.com/defry256/pokemon-helper/internal/logger"
	queue "github.com/defryheryanto/job-queuer"
	otel "go.opentelemetry.io/otel"
)

func main() {
	config.Load()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	queuer := queue.NewQueuer(config.MaxQueueWorker())
	queuer.Run(context.Background())
	log.Println("queuer successfully running")

	tracerProvider := setupTracer()
	otel.SetTracerProvider(tracerProvider)
	tracer := tracerProvider.Tracer("pokemon-helper")

	var appserver *http.Server
	go func() {
		redisClient := setupRedis()
		app := BuildApp(redisClient, queuer, tracer)
		appserver = &http.Server{
			Addr:    fmt.Sprintf(":%s", config.HostPort()),
			Handler: httpserver.HandleRoutes(app, tracer),
		}
		logger.Print(fmt.Sprintf("Application Server listening on %s", appserver.Addr))
		err := appserver.ListenAndServe()
		if err != nil {
			logger.Error(fmt.Sprintf("error listen - %v", err), err)
		}
	}()

	<-quit
	shutdownServer(60*time.Second, appserver)
}

func shutdownServer(timeout time.Duration, server *http.Server) {
	cto, cancel := context.WithTimeout(context.Background(), timeout)
	if e := server.Shutdown(cto); e != nil && e != http.ErrServerClosed {
		logger.Error(fmt.Sprintf("Shutdown failed for server in address: %s, %v", server.Addr, e), e)
	}
	cancel()
}
