package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/defryheryanto/pokemon-helper/config"
	"github.com/defryheryanto/pokemon-helper/internal/httpserver"
	"github.com/defryheryanto/pokemon-helper/internal/logger"
	otel "go.opentelemetry.io/otel"
)

func main() {
	config.Load()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	tracer := otel.GetTracerProvider().Tracer("")
	if config.TracingEnabled() {
		tracerProvider := setupTracer()
		otel.SetTracerProvider(tracerProvider)
		tracer = tracerProvider.Tracer("pokemon-helper")
	}

	var appserver *http.Server
	go func() {
		redisClient := setupRedis()
		app := BuildApp(redisClient, tracer)
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
