package application

import (
	"context"

	"github.com/taguch1/try-bff/apps/bff-grpc/domain/model"
	"github.com/taguch1/try-bff/apps/bff-grpc/domain/service"
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
	todoService service.Todo
}

// NewTodo Constructor
func NewTodo(todoService service.Todo) Todo {
	return &todoImpl{todoService}
}

func (app *todoImpl) Save(ctx context.Context, req *model.TodoSaveRequest) (*model.TodoResponse, error) {
	todo, err := app.todoService.Save(ctx, req.Title)
	if err != nil {
		return nil, err
	}
	return model.NewTodoResponse(todo), nil
}

func (app *todoImpl) Get(ctx context.Context, req *model.TodoGetRequest) (*model.TodoResponse, error) {
	todo, err := app.todoService.Get(ctx, req.ID)
	if err != nil {
		return nil, err
	}
	return model.NewTodoResponse(todo), nil
}

func (app *todoImpl) List(ctx context.Context, req *model.TodoListRequest) (*model.TodoListResponse, error) {
	todos, err := app.todoService.List(ctx, req.Offset, req.Limit)
	if err != nil {
		return nil, err
	}
	return model.NewTodoListResponse(todos), nil
}

func (app *todoImpl) Update(ctx context.Context, req *model.TodoUpdateRequest) (*model.TodoResponse, error) {
	todo, err := app.todoService.Update(ctx, req.ID, req.Title)
	if err != nil {
		return nil, err
	}
	return model.NewTodoResponse(todo), nil
}

func (app *todoImpl) Delete(ctx context.Context, req *model.TodoDeleteRequest) error {
	return app.todoService.Delete(ctx, req.ID)
}
