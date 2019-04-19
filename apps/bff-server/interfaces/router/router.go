package router

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/taguch1/try-bff/apps/bff-server/infrastructure/log"
	"github.com/taguch1/try-bff/apps/bff-server/interfaces/handler"
)

// NewHTTPRouter constructor
func NewHTTPRouter(
	healthHandler handler.Health,
	todoHandler handler.Todo,
) http.Handler {

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.RequestLogger(log.NewLogFormatter(log.Logger)))
	r.Use(middleware.Recoverer)
	r.Get("/health", healthHandler.Get)
	r.Route("/todos", func(r chi.Router) {
		r.Post("/", todoHandler.Save)
		r.Get("/", todoHandler.List)
		r.Route("/{todoID}", func(r chi.Router) {
			r.Use(todoHandler.Ctx)
			r.Get("/", todoHandler.Get)
			r.Patch("/", todoHandler.Update)
			r.Delete("/", todoHandler.Delete)
		})
	})
	return r
}