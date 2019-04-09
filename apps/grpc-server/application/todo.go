package application

import (
	"context"

	pb "github.com/taguch1/try-bff/apps/grpc-server/proto"
)

type todoAppImpl struct {
}

// NewTodoApp Constructor
func NewTodoApp() pb.TodoServer {
	return &todoAppImpl{}
}

func (app *todoAppImpl) Save(ctx context.Context, req *pb.TodoSaveRequest) (*pb.TodoResponse, error) {
	return &pb.TodoResponse{Id: "1", Title: "TitleA"}, nil
}

func (app *todoAppImpl) Get(ctx context.Context, req *pb.TodoGetRequest) (*pb.TodoResponse, error) {
	return &pb.TodoResponse{Id: req.Id, Title: "TitleX"}, nil
}

func (app *todoAppImpl) List(ctx context.Context, req *pb.TodoListRequest) (*pb.TodoListResponse, error) {
	return &pb.TodoListResponse{
		Todos: []*pb.TodoResponse{
			{Id: "1", Title: "TitleA"},
			{Id: "2", Title: "TitleB"},
		},
	}, nil
}

func (app *todoAppImpl) Update(ctx context.Context, req *pb.TodoUpdateRequest) (*pb.TodoResponse, error) {
	return &pb.TodoResponse{Id: req.Id, Title: req.Title}, nil
}

func (app *todoAppImpl) Delete(ctx context.Context, req *pb.TodoDeleteRequest) (*pb.TodoEmptyResponse, error) {
	return &pb.TodoEmptyResponse{}, nil
	// return nil, status.Errorf(codes.InvalidArgument, "Ouch!")
}
