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
	"github.com/defry256/pokemon-helper/config/env"
	"github.com/defry256/pokemon-helper/internal/httpserver"
	"github.com/defry256/pokemon-helper/internal/logger"
	queue "github.com/defryheryanto/job-queuer"
)

func main() {
	env.LoadEnv()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	queuer := queue.NewQueuer(config.MAX_QUEUE_WORKER())
	queuer.Run(context.Background())
	log.Println("queuer successfully running")

	var appserver *http.Server
	go func() {
		redisClient := setupRedis()
		app := BuildApp(redisClient, queuer)
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
