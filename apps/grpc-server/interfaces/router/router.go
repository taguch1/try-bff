package router

import (
	pb "github.com/taguch1/try-bff/apps/grpc-server/proto"
	"google.golang.org/grpc"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

// NewGPRCRouter constructor
func NewGPRCRouter(
	healthServer healthpb.HealthServer,
	todoApp pb.TodoServer,
) *grpc.Server {
	s := grpc.NewServer()
	healthpb.RegisterHealthServer(s, healthServer)
	pb.RegisterTodoServer(s, todoApp)
	reflection.Register(s)
	return s
}
