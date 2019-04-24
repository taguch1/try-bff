package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/taguch1/try-bff/apps/bff-server/application"
	"github.com/taguch1/try-bff/apps/bff-server/infrastructure/grpc"
	"github.com/taguch1/try-bff/apps/bff-server/infrastructure/log"
	"github.com/taguch1/try-bff/apps/bff-server/interfaces/handler"
	"github.com/taguch1/try-bff/apps/bff-server/interfaces/middleware"
	"github.com/taguch1/try-bff/apps/bff-server/interfaces/router"
)

const (
	version     = "v0.0.0"
	httpAddress = ":1323"
)

var revision string

func init() {
	log.Setup()
}
func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) > 0 && args[0] == "version" {
		fmt.Printf(`backends for frontends mock server. version:%s  revision: %s`, version, revision)
		os.Exit(0)
	}

	ctx := context.Background()
	httpServer := newServer(ctx)

	log.Info(ctx, "start http server")
	startServer(ctx, newServer(ctx))

	quitCh := make(chan os.Signal)
	signal.Notify(quitCh, syscall.SIGINT, syscall.SIGTERM)
	<-quitCh

	log.Info(ctx, "shutdown http server")
	shutdownServer(ctx, httpServer)
}

func newServer(ctx context.Context) *http.Server {

	grpcConfig, err := grpc.NewConf(grpc.ConfFileName)
	if err != nil {
		log.Fatalf(ctx, "failed to load grpc config. err:%s", err)
	}
	middlewareConfig, err := middleware.NewConf(middleware.ConfFileName)
	if err != nil {
		log.Fatalf(ctx, "failed to load grpc config. err:%s", err)
	}

	todoService, _ := grpc.NewTodoService(grpcConfig)
	todoApp := application.NewTodo(todoService)
	healthHandler := handler.NewHealth()
	todoHandler := handler.NewTodo(todoApp)

	r := router.NewHTTPRouter(
		middlewareConfig,
		healthHandler,
		todoHandler,
	)

	httpServer := &http.Server{Addr: httpAddress, Handler: r}
	return httpServer
}

func startServer(ctx context.Context, httpServer *http.Server) {
	// http server
	go func() {
		if err := httpServer.ListenAndServe(); err != nil {
			log.Fatalf(ctx, "failed to start the http server  %s", err)
		}
	}()
}

func shutdownServer(ctx context.Context, httpServer *http.Server) {

	//TODO
	// - Readinessチェックをフラグで制御
	// 1. unreadyにする
	// 2. ReadinessProbeの(failureThreshold * periodSeconds) + バッファ分待つ
	// 3. ReadinessチェックでPodが外れてからシャットダウン処理をする

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		if err := httpServer.Shutdown(ctx); err != nil {
			log.Errorf(ctx, "failed to shutdown the http server: %s", err)
		}
		wg.Done()
	}()
	wg.Wait()
}
