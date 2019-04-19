package application

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/taguch1/try-bff/apps/bff-server/domain/model"
	mock_service "github.com/taguch1/try-bff/apps/bff-server/test/mock/domain/service"
)

var todos = []*model.Todo{
	&model.Todo{ID: "1", Title: "TitleA"},
	&model.Todo{ID: "2", Title: "TitleB"},
	&model.Todo{ID: "3", Title: "TitleC"},
	&model.Todo{ID: "4", Title: "TitleD"},
}

var todoResponses = []*model.TodoResponse{
	&model.TodoResponse{ID: "1", Title: "TitleA"},
	&model.TodoResponse{ID: "2", Title: "TitleB"},
	&model.TodoResponse{ID: "3", Title: "TitleC"},
	&model.TodoResponse{ID: "4", Title: "TitleD"},
}

var todoListResponses = []*model.TodoListResponse{
	&model.TodoListResponse{Todos: todoResponses[0:2]},
	&model.TodoListResponse{Todos: todoResponses[2:]},
}

func TestTodo_SaveSuccess(t *testing.T) {
	ctx := context.Background()
	srv := &mock_service.TodoService{}
	srv.On("Save", ctx, todos[0].Title).Return(todos[0], nil)
	srv.On("Save", ctx, todos[1].Title).Return(todos[1], nil)
	app := NewTodo(srv)

	cases := []struct {
		input    *model.TodoSaveRequest
		expected *model.TodoResponse
	}{
		{&model.TodoSaveRequest{Title: todos[0].Title}, todoResponses[0]},
		{&model.TodoSaveRequest{Title: todos[1].Title}, todoResponses[1]},
	}

	for _, c := range cases {
		res, err := app.Save(ctx, c.input)
		assert.Nil(t, err)
		assert.Equal(t, c.expected, res)
	}
}
func TestTodo_GetSuccess(t *testing.T) {
	ctx := context.Background()
	srv := &mock_service.TodoService{}
	srv.On("Get", ctx, todos[0].ID).Return(todos[0], nil)
	srv.On("Get", ctx, todos[1].ID).Return(todos[1], nil)
	app := NewTodo(srv)

	cases := []struct {
		input    *model.TodoGetRequest
		expected *model.TodoResponse
	}{
		{&model.TodoGetRequest{ID: todos[0].ID}, todoResponses[0]},
		{&model.TodoGetRequest{ID: todos[1].ID}, todoResponses[1]},
	}

	for _, c := range cases {
		res, err := app.Get(ctx, c.input)
		assert.Nil(t, err)
		assert.Equal(t, c.expected, res)
	}

}
func TestTodo_ListSuccess(t *testing.T) {
	ctx := context.Background()
	srv := &mock_service.TodoService{}
	srv.On("List", ctx, int64(0), int64(2)).Return(todos[0:2], nil)
	srv.On("List", ctx, int64(2), int64(2)).Return(todos[2:4], nil)
	app := NewTodo(srv)

	cases := []struct {
		input    *model.TodoListRequest
		expected *model.TodoListResponse
	}{
		{&model.TodoListRequest{Offset: 0, Limit: 2}, todoListResponses[0]},
		{&model.TodoListRequest{Offset: 2, Limit: 2}, todoListResponses[1]},
	}
	for _, c := range cases {
		res, err := app.List(ctx, c.input)
		assert.Nil(t, err)
		assert.Equal(t, c.expected, res)
	}

}
func TestTodo_UpdateSuccess(t *testing.T) {
	ctx := context.Background()
	srv := &mock_service.TodoService{}
	srv.On("Update", ctx, todos[0].ID, todos[0].Title).Return(todos[0], nil)
	srv.On("Update", ctx, todos[1].ID, todos[1].Title).Return(todos[1], nil)
	app := NewTodo(srv)

	cases := []struct {
		input    *model.TodoUpdateRequest
		expected *model.TodoResponse
	}{
		{&model.TodoUpdateRequest{ID: todos[0].ID, Title: todos[0].Title}, todoResponses[0]},
		{&model.TodoUpdateRequest{ID: todos[1].ID, Title: todos[1].Title}, todoResponses[1]},
	}
	for _, c := range cases {
		res, err := app.Update(ctx, c.input)
		assert.Nil(t, err)
		assert.Equal(t, c.expected, res)
	}

}
func TestTodo_DeleteSuccess(t *testing.T) {
	ctx := context.Background()
	srv := &mock_service.TodoService{}
	srv.On("Delete", ctx, todos[0].ID).Return(nil)
	srv.On("Delete", ctx, todos[1].ID).Return(nil)
	app := NewTodo(srv)

	cases := []struct {
		input    *model.TodoDeleteRequest
		expected error
	}{
		{&model.TodoDeleteRequest{ID: todos[0].ID}, nil},
		{&model.TodoDeleteRequest{ID: todos[1].ID}, nil},
	}
	for _, c := range cases {
		err := app.Delete(ctx, c.input)
		assert.Equal(t, c.expected, err)
	}
}
