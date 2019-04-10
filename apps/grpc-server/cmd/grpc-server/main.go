package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/taguch1/try-bff/apps/grpc-server/application"
	"github.com/taguch1/try-bff/apps/grpc-server/interfaces/router"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
)

const (
	version     = "v0.0.0"
	grpcAddress = ":50051"
)

var revision string

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) > 0 && args[0] == "version" {
		fmt.Printf(`gRPC mock server. version:%s  revision: %s`, version, revision)
		os.Exit(0)
	}

	ctx := context.Background()
	gPRCRouter := newServer(ctx)
	startServer(ctx, gPRCRouter)

	quitCh := make(chan os.Signal)
	signal.Notify(quitCh, syscall.SIGINT, syscall.SIGTERM)
	<-quitCh

	shutdownServer(ctx, gPRCRouter)
}

func newServer(ctx context.Context) *grpc.Server {
	healthServer := health.NewServer()
	todoApp := application.NewTodoApp()
	gPRCRouter := router.NewGPRCRouter(healthServer, todoApp)
	return gPRCRouter
}

func startServer(ctx context.Context, gPRCRouter *grpc.Server) {
	// gRPC server
	lis, err := net.Listen("tcp", grpcAddress)
	if err != nil {
		log.Fatalf("error grpc server start %s", err)
	}

	go func() {
		if err := gPRCRouter.Serve(lis); err != nil {
			log.Printf("shutting down the grpc server %s", err)
		}
	}()
}

func shutdownServer(ctx context.Context, gPRCRouter *grpc.Server) {

	//TODO
	// - Readinessチェックをフラグで制御
	// 1. unreadyにする
	// 2. ReadinessProbeの(failureThreshold * periodSeconds) + バッファ分待つ
	// 3. ReadinessチェックでPodが外れてからシャットダウン処理をする

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		gPRCRouter.GracefulStop()
		wg.Done()
	}()
	wg.Wait()
}
