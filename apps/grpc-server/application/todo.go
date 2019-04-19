package application

import (
	"context"

	"github.com/taguch1/try-bff/apps/grpc-server/domain/repository"
	pb "github.com/taguch1/try-bff/apps/grpc-server/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type todoImpl struct {
	todoRepo repository.Todo
}

// NewTodo constructor
func NewTodo(todoRepo repository.Todo) pb.TodoServer {
	return &todoImpl{todoRepo}
}

func (app *todoImpl) Save(ctx context.Context, req *pb.TodoSaveRequest) (*pb.TodoResponse, error) {

	row, err := app.todoRepo.Save(ctx, app.todoRepo.NextID(ctx), req.Title)
	if err != nil {
		return nil, err
	}
	return &pb.TodoResponse{Id: row.ID, Title: row.Title}, nil
}

func (app *todoImpl) Get(ctx context.Context, req *pb.TodoGetRequest) (*pb.TodoResponse, error) {
	row, err := app.todoRepo.Get(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	if row == nil {
		return nil, status.Errorf(codes.NotFound, "not found todo")
	}
	return &pb.TodoResponse{Id: row.ID, Title: row.Title}, nil
}

func (app *todoImpl) List(ctx context.Context, req *pb.TodoListRequest) (*pb.TodoListResponse, error) {
	rows, err := app.todoRepo.List(ctx, req.Offset, req.Limit)
	if err != nil {
		return nil, err
	}
	todos := make([]*pb.TodoResponse, len(rows))
	for i, row := range rows {
		todos[i] = &pb.TodoResponse{
			Id:    row.ID,
			Title: row.Title,
		}
	}
	return &pb.TodoListResponse{
		Todos: todos,
	}, nil
}

func (app *todoImpl) Update(ctx context.Context, req *pb.TodoUpdateRequest) (*pb.TodoResponse, error) {
	row, err := app.todoRepo.Update(ctx, req.Id, req.Title)
	if err != nil {
		return nil, err
	}
	return &pb.TodoResponse{Id: row.ID, Title: row.Title}, nil
}

func (app *todoImpl) Delete(ctx context.Context, req *pb.TodoDeleteRequest) (*pb.TodoEmptyResponse, error) {
	err := app.todoRepo.Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.TodoEmptyResponse{}, nil
}
