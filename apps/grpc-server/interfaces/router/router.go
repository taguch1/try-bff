package router

import (
	"time"

	"github.com/taguch1/try-bff/apps/grpc-server/infrastructure/log"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	pb "github.com/taguch1/try-bff/apps/grpc-server/proto"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

// NewGPRCRouter constructor
func NewGPRCRouter(
	healthServer healthpb.HealthServer,
	todoApp pb.TodoServer,
) *grpc.Server {

	// TODO: move interface/middleware?
	opts := []grpc_zap.Option{
		grpc_zap.WithDurationField(func(duration time.Duration) zapcore.Field {
			return zap.Int64("grpc.time_ns", duration.Nanoseconds())
		}),
		grpc_zap.WithDecider(func(fullMethodName string, err error) bool {
			// TODO: grpc
			if fullMethodName == "/grpc.health.v1.Health/Check" {
				return false
			}
			return true
		}),
	}
	grpc_zap.ReplaceGrpcLogger(log.Logger)

	s := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			grpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			grpc_zap.UnaryServerInterceptor(log.Logger, opts...),
		),
	)

	healthpb.RegisterHealthServer(s, healthServer)
	pb.RegisterTodoServer(s, todoApp)
	reflection.Register(s)
	return s
}
