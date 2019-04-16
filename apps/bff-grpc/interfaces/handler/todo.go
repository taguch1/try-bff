package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
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

const todoIDParam string = "id"

// TODO: render
func (h *todoImpl) Save(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	req := &model.TodoSaveRequest{}
	if err := render.Bind(r, req); err != nil {
		handleError(w, r, err)
		return
	}
	res, _ := h.todoApp.Save(ctx, req)
	resJSON, _ := json.Marshal(res)
	w.WriteHeader(http.StatusCreated)
	w.Write(resJSON)
}

func (h *todoImpl) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := chi.URLParam(r, todoIDParam)
	req := &model.TodoGetRequest{ID: id}
	if err := render.Bind(r, req); err != nil {
		handleError(w, r, err)
		return
	}
	res, _ := h.todoApp.Get(ctx, req)
	resJSON, _ := json.Marshal(res)
	w.WriteHeader(http.StatusOK)
	w.Write(resJSON)
}

func (h *todoImpl) List(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	req := &model.TodoListRequest{}
	if err := render.Bind(r, req); err != nil {
		handleError(w, r, err)
		return
	}
	res, err := h.todoApp.List(ctx, req)
	if err != nil {
		handleError(w, r, err)
		return
	}
	resJSON, _ := json.Marshal(res)
	w.WriteHeader(http.StatusOK)
	w.Write(resJSON)
}

func (h *todoImpl) Update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := chi.URLParam(r, todoIDParam)
	req := &model.TodoUpdateRequest{ID: id}
	if err := render.Bind(r, req); err != nil {
		handleError(w, r, err)
		return
	}
	res, _ := h.todoApp.Update(ctx, req)
	resJSON, _ := json.Marshal(res)
	w.WriteHeader(http.StatusOK)
	w.Write(resJSON)
}

func (h *todoImpl) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := chi.URLParam(r, todoIDParam)
	req := &model.TodoDeleteRequest{ID: id}
	if err := render.Bind(r, req); err != nil {
		handleError(w, r, err)
		return
	}
	h.todoApp.Delete(ctx, req)
	w.WriteHeader(http.StatusNoContent)
}
