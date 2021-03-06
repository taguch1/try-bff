package model

import (
	"net/http"
)

// Todo todo
type Todo struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

// TodoSaveRequest save request
type TodoSaveRequest struct {
	Title string `json:"title"`
}

// TodoGetRequest get request
type TodoGetRequest struct {
	ID string `json:"id"`
}

// TodoListRequest list request
type TodoListRequest struct {
	Offset int64 `json:"offset"`
	Limit  int64 `json:"limit"`
}

// TodoUpdateRequest update request
type TodoUpdateRequest struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

// TodoDeleteRequest delete request
type TodoDeleteRequest struct {
	ID string `json:"id"`
}

// TodoResponse response
type TodoResponse struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

// TodoListResponse list response
type TodoListResponse struct {
	Todos []*TodoResponse `json:"todos"`
}

// NewTodoResponse new response
func NewTodoResponse(todo *Todo) *TodoResponse {
	return &TodoResponse{ID: todo.ID, Title: todo.Title}
}

// NewTodoListResponse new response
func NewTodoListResponse(todos []*Todo) *TodoListResponse {

	todoList := make([]*TodoResponse, len(todos))
	for i, todo := range todos {
		todoList[i] = &TodoResponse{
			ID:    todo.ID,
			Title: todo.Title,
		}
	}
	return &TodoListResponse{
		Todos: todoList,
	}
}

// Bind validate request
func (req *TodoSaveRequest) Bind(r *http.Request) error {
	return nil
}

// Bind validate request
func (req *TodoGetRequest) Bind(r *http.Request) error {
	return nil
}

// Bind validate request
func (req *TodoListRequest) Bind(r *http.Request) error {
	return nil
}

// Bind validate request
func (req *TodoUpdateRequest) Bind(r *http.Request) error {
	return nil
}

// Bind validate request
func (req *TodoDeleteRequest) Bind(r *http.Request) error {
	return nil
}
