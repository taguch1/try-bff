package router

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/taguch1/try-bff/apps/bff-grpc/interfaces/handler"
)

// NewHTTPRouter constructor
func NewHTTPRouter(
	healthHandler handler.Health,
	todoHandler handler.Todo,
) http.Handler {
	r := chi.NewRouter()
	r.Get("/health", healthHandler.Get)
	r.Route("/todos", func(r chi.Router) {
		r.Post("/", todoHandler.Save)
		r.Get("/", todoHandler.List)
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", todoHandler.Get)
			r.Patch("/", todoHandler.Update)
			r.Delete("/", todoHandler.Delete)
		})
	})
	return r
}
