package main

import (
	"context"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/eclipse-orbital-systems/mission-control-api/api/health"
	"github.com/eclipse-orbital-systems/mission-control-api/api/v1"
	"github.com/eclipse-orbital-systems/mission-control-api/config"
	"github.com/eclipse-orbital-systems/mission-control-api/dal"
	"github.com/palantir/stacktrace"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	log.Println("Starting service...")

	initValidators()

	dalImpl := dal.New(config.MySql)

	apiEngine := gin.New()
	healthEngine := gin.New()

	corsMiddleware := cors.New(config.Cors)
	apiEngine.Use(corsMiddleware)

	healthRouter := healthEngine.Group("/health")
	v1Router := apiEngine.Group("/v1")
	v1Router.Use(gin.Logger(), gin.Recovery())

	health.New(healthRouter, config.MySql)
	v1.New(v1Router, dalImpl)

	runGroup, runGroupCtx := errgroup.WithContext(context.Background())
	runGroup.Go(func() error {
		return HandleSignals(runGroupCtx)
	})
	runGroup.Go(func() error {
		return ListenAndServeWithContext(runGroupCtx, "8081", healthEngine)
	})
	runGroup.Go(func() error {
		return ListenAndServeWithContext(runGroupCtx, "8080", apiEngine)
	})

	log.Print("listening...")
	<-runGroupCtx.Done()
	log.Print("shutting down...")
	log.Print(runGroup.Wait().Error())
	log.Print("shutdown complete")
}

func HandleSignals(ctx context.Context) error {
	sigintc := make(chan os.Signal, 1)
	signal.Notify(sigintc, syscall.SIGINT)

	sigtermc := make(chan os.Signal, 1)
	signal.Notify(sigtermc, syscall.SIGTERM)

	select {
	case <-sigintc:
		return stacktrace.NewError("sigint received")
	case <-sigtermc:
		return stacktrace.NewError("sigterm received")
	case <-ctx.Done():
		return nil
	}
}

func ListenAndServeWithContext(ctx context.Context, port string, handler http.Handler) error {
	server := &http.Server{Addr: "0.0.0.0:" + port, Handler: handler}

	serverErr := make(chan error)
	go func() {
		serverErr <- server.ListenAndServe()
	}()

	for {
		select {
		case err := <-serverErr:
			return err
		case <-ctx.Done():
			server.Shutdown(context.Background())
		}
	}
}
