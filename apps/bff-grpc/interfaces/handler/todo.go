package handler

import (
	"encoding/json"
	"net/http"

	"github.com/taguch1/try-bff/apps/bff-grpc/application"
	"github.com/taguch1/try-bff/apps/bff-grpc/domain/model"
)

// Todo handler
type Todo interface {
	Get(w http.ResponseWriter, r *http.Request)
	List(w http.ResponseWriter, r *http.Request)
	Save(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

type todoImpl struct {
	todoApp application.Todo
}

//NewTodo  handler constructor
func NewTodo(todoApp application.Todo) Todo {
	return &todoImpl{todoApp}
}

// TODO: render
// TODO: bind

func (h *todoImpl) Save(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	req := &model.TodoSaveRequest{}
	res, _ := h.todoApp.Save(ctx, req)
	resJSON, _ := json.Marshal(res)
	w.WriteHeader(http.StatusCreated)
	w.Write(resJSON)
}

func (h *todoImpl) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	req := &model.TodoGetRequest{}
	res, _ := h.todoApp.Get(ctx, req)
	resJSON, _ := json.Marshal(res)
	w.WriteHeader(http.StatusOK)
	w.Write(resJSON)
}

func (h *todoImpl) List(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	req := &model.TodoListRequest{}
	res, _ := h.todoApp.List(ctx, req)
	resJSON, _ := json.Marshal(res)
	w.WriteHeader(http.StatusOK)
	w.Write(resJSON)
}

func (h *todoImpl) Update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	req := &model.TodoUpdateRequest{}
	res, _ := h.todoApp.Update(ctx, req)
	resJSON, _ := json.Marshal(res)
	w.WriteHeader(http.StatusOK)
	w.Write(resJSON)
}

func (h *todoImpl) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	req := &model.TodoDeleteRequest{}
	h.todoApp.Delete(ctx, req)
	w.WriteHeader(http.StatusNoContent)
}
