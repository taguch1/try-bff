package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/taguch1/try-bff/apps/bff-grpc/application"
	"github.com/taguch1/try-bff/apps/bff-grpc/domain/model"
)

// Todo handler
type Todo interface {
	Ctx(next http.Handler) http.Handler
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

const todoIDParam string = "todoID"

type key int

const (
	todoCtxKey key = iota
)

func (h *todoImpl) Ctx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		id := chi.URLParam(r, todoIDParam)
		res, err := h.todoApp.Get(ctx, &model.TodoGetRequest{ID: id})
		if err != nil {
			handleError(w, r, err)
			return
		}
		newCtx := context.WithValue(ctx, todoCtxKey, res)
		next.ServeHTTP(w, r.WithContext(newCtx))
	})
}

// TODO: render
func (h *todoImpl) Save(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	req := &model.TodoSaveRequest{}

	if err := render.Bind(r, req); err != nil {
		handleError(w, r, err)
		return
	}

	res, err := h.todoApp.Save(ctx, req)
	if err != nil {
		handleError(w, r, err)
		return
	}

	resJSON, err := json.Marshal(res)
	if err != nil {
		handleError(w, r, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(resJSON)
}

func (h *todoImpl) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	todo := ctx.Value(todoCtxKey).(*model.TodoResponse)

	resJSON, err := json.Marshal(todo)
	if err != nil {
		handleError(w, r, err)
		return
	}

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
	todo := ctx.Value(todoCtxKey).(*model.TodoResponse)

	req := &model.TodoUpdateRequest{ID: todo.ID}
	if err := render.Bind(r, req); err != nil {
		handleError(w, r, err)
		return
	}

	res, err := h.todoApp.Update(ctx, req)
	if err != nil {
		handleError(w, r, err)
		return
	}

	resJSON, _ := json.Marshal(res)
	w.WriteHeader(http.StatusOK)
	w.Write(resJSON)
}

func (h *todoImpl) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	todo := ctx.Value(todoCtxKey).(*model.TodoResponse)
	req := &model.TodoDeleteRequest{ID: todo.ID}

	if err := h.todoApp.Delete(ctx, req); err != nil {
		handleError(w, r, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
