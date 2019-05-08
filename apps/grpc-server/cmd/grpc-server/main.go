package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/taguch1/try-bff/apps/grpc-server/application"
	"github.com/taguch1/try-bff/apps/grpc-server/infrastructure/log"
	"github.com/taguch1/try-bff/apps/grpc-server/infrastructure/mysql"
	"github.com/taguch1/try-bff/apps/grpc-server/infrastructure/persistence"
	"github.com/taguch1/try-bff/apps/grpc-server/interfaces/router"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
)

const (
	version     = "v0.0.0"
	grpcAddress = ":50051"
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
		fmt.Printf(`gRPC mock server. version:%s  revision: %s`, version, revision)
		os.Exit(0)
	}

	ctx := context.Background()
	gPRCRouter, httpServer := newServer(ctx)

	log.Info(ctx, "start grpc server")
	startServer(ctx, gPRCRouter, httpServer)

	quitCh := make(chan os.Signal)
	signal.Notify(quitCh, syscall.SIGINT, syscall.SIGTERM)
	<-quitCh

	log.Info(ctx, "shutdown grpc server")
	shutdownServer(ctx, gPRCRouter, httpServer)
}

func newServer(ctx context.Context) (*grpc.Server, *http.Server) {
	healthServer := health.NewServer()
	mysqlConfig, err := mysql.NewConf(mysql.ConfFileName)
	if err != nil {
		log.Fatalf(ctx, "failed to load mysql config  %s", err)
	}
	db, err := mysql.Open(mysqlConfig)
	if err != nil {
		log.Fatalf(ctx, "failed to open mysql connection  %s", err)
	}
	todoRepo := persistence.NewTodo(db)
	todoApp := application.NewTodo(todoRepo)
	gPRCRouter := router.NewGPRCRouter(healthServer, todoApp)

	handler := http.NewServeMux()
	handler.Handle("/metrics", promhttp.Handler())
	httpServer := &http.Server{Addr: httpAddress, Handler: handler}

	return gPRCRouter, httpServer
}

func startServer(ctx context.Context, gPRCRouter *grpc.Server, httpServer *http.Server) {
	// gRPC server
	lis, err := net.Listen("tcp", grpcAddress)
	if err != nil {
		log.Fatalf(ctx, "failed to listen grpc server port  %s", err)
	}

	go func() {
		if err := gPRCRouter.Serve(lis); err != nil {
			log.Fatalf(ctx, "failed to start the grpc server  %s", err)
		}
	}()

	go func() {
		if err := httpServer.ListenAndServe(); err != nil {
			log.Fatalf(ctx, "failed to start the http server  %s", err)
		}
	}()

}

func shutdownServer(ctx context.Context, gPRCRouter *grpc.Server, httpServer *http.Server) {

	//TODO
	// - Readinessチェックをフラグで制御
	// 1. unreadyにする
	// 2. ReadinessProbeの(failureThreshold * periodSeconds) + バッファ分待つ
	// 3. ReadinessチェックでPodが外れてからシャットダウン処理をする

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		gPRCRouter.GracefulStop()
		wg.Done()
	}()

	go func() {
		if err := httpServer.Shutdown(ctx); err != nil {
			log.Errorf(ctx, "failed to shutdown the http server: %s", err)
		}
		wg.Done()
	}()

	wg.Wait()
}
