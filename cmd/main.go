package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/defry256/pokemon-helper/config"
	"github.com/defry256/pokemon-helper/config/env"
	"github.com/defry256/pokemon-helper/internal/httpserver"
	"github.com/defry256/pokemon-helper/internal/logger"
)

func main() {
	env.LoadEnv()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	var appserver *http.Server
	go func() {
		app := BuildApp()
		appserver = &http.Server{
			Addr:    fmt.Sprintf("%s:%s", config.HOST_URL(), config.HOST_PORT()),
			Handler: httpserver.HandleRoutes(app),
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
