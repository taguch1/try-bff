package grpc

import (
	"context"
	"time"

	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/taguch1/try-bff/apps/bff-server/domain/model"
	"github.com/taguch1/try-bff/apps/bff-server/domain/service"
	pb "github.com/taguch1/try-bff/apps/bff-server/proto"
	"google.golang.org/grpc"
)

// TODO: move config
// const targetAddress = "grpc-server:50051"
const targetAddress = "localhost:50051"
const timeoutMillis = 3000

type todoSrviceImpl struct {
	timeout    time.Duration
	todoClient pb.TodoClient
}

// NewTodoService todo service constructor
func NewTodoService(config *Config) (service.Todo, *prometheus.Registry, error) {

	reg := prometheus.NewRegistry()
	grpcMetrics := grpc_prometheus.NewClientMetrics()
	reg.MustRegister(grpcMetrics)

	conn, err := grpc.Dial(
		config.TargetAddress,
		grpc.WithUnaryInterceptor(grpcMetrics.UnaryClientInterceptor()),
		grpc.WithStreamInterceptor(grpcMetrics.StreamClientInterceptor()),
		grpc.WithInsecure(),
	)
	if err != nil {
		return nil, nil, err
	}
	timeout := time.Duration(config.TimeoutMillis) * time.Millisecond
	todoClient := pb.NewTodoClient(conn)

	return &todoSrviceImpl{timeout, todoClient}, reg, nil
}

func (s *todoSrviceImpl) Save(ctx context.Context, title string) (*model.Todo, error) {
	timeoutCtx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()
	in := &pb.TodoSaveRequest{Title: title}
	res, err := s.todoClient.Save(timeoutCtx, in)
	if err != nil {
		return nil, err
	}
	return &model.Todo{
		ID:    res.Id,
		Title: res.Title,
	}, nil
}

func (s *todoSrviceImpl) Get(ctx context.Context, id string) (*model.Todo, error) {
	in := &pb.TodoGetRequest{Id: id}
	timeoutCtx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()
	res, err := s.todoClient.Get(timeoutCtx, in)
	if err != nil {
		return nil, err
	}
	return &model.Todo{
		ID:    res.Id,
		Title: res.Title,
	}, nil
}

func (s *todoSrviceImpl) List(ctx context.Context, offset, limit int64) ([]*model.Todo, error) {
	in := &pb.TodoListRequest{
		Offset: offset,
		Limit:  limit,
	}
	timeoutCtx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()
	res, err := s.todoClient.List(timeoutCtx, in)
	if err != nil {
		return nil, err
	}

	todos := make([]*model.Todo, len(res.Todos))
	for i, todo := range res.Todos {
		todos[i] = &model.Todo{
			ID:    todo.Id,
			Title: todo.Title,
		}
	}
	return todos, nil
}

func (s *todoSrviceImpl) Update(ctx context.Context, id, title string) (*model.Todo, error) {
	in := &pb.TodoUpdateRequest{Id: id, Title: title}
	timeoutCtx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()
	res, err := s.todoClient.Update(timeoutCtx, in)
	if err != nil {
		return nil, err
	}
	return &model.Todo{
		ID:    res.Id,
		Title: res.Title,
	}, nil
}

func (s *todoSrviceImpl) Delete(ctx context.Context, id string) error {
	in := &pb.TodoDeleteRequest{Id: id}
	timeoutCtx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()
	_, err := s.todoClient.Delete(timeoutCtx, in)
	return err
}
