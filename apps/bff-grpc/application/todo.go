package application

import (
	"context"

	"github.com/taguch1/try-bff/apps/bff-grpc/domain/model"
)

// Todo todo application service
type Todo interface {
	Save(ctx context.Context, req *model.TodoSaveRequest) (*model.TodoResponse, error)
	Get(ctx context.Context, req *model.TodoGetRequest) (*model.TodoResponse, error)
	List(ctx context.Context, req *model.TodoListRequest) (*model.TodoListResponse, error)
	Update(ctx context.Context, req *model.TodoUpdateRequest) (*model.TodoResponse, error)
	Delete(ctx context.Context, req *model.TodoDeleteRequest) error
}

type todoImpl struct {
}

// NewTodo Constructor
func NewTodo() Todo {
	return &todoImpl{}
}

func (app *todoImpl) Save(ctx context.Context, req *model.TodoSaveRequest) (*model.TodoResponse, error) {
	return &model.TodoResponse{ID: "1", Title: "TitleA"}, nil
}

func (app *todoImpl) Get(ctx context.Context, req *model.TodoGetRequest) (*model.TodoResponse, error) {
	return &model.TodoResponse{ID: req.ID, Title: "TitleX"}, nil
}

func (app *todoImpl) List(ctx context.Context, req *model.TodoListRequest) (*model.TodoListResponse, error) {
	return &model.TodoListResponse{
		Todos: []*model.TodoResponse{
			{ID: "1", Title: "TitleA"},
			{ID: "2", Title: "TitleB"},
		},
	}, nil
}

func (app *todoImpl) Update(ctx context.Context, req *model.TodoUpdateRequest) (*model.TodoResponse, error) {
	return &model.TodoResponse{ID: req.ID, Title: req.Title}, nil
}

func (app *todoImpl) Delete(ctx context.Context, req *model.TodoDeleteRequest) error {
	return nil
}
