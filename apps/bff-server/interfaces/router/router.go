package router

import (
	"net/http"

	"github.com/go-chi/chi"
	chi_middleware "github.com/go-chi/chi/middleware"
	"github.com/taguch1/try-bff/apps/bff-server/infrastructure/log"
	"github.com/taguch1/try-bff/apps/bff-server/interfaces/handler"
	"github.com/taguch1/try-bff/apps/bff-server/interfaces/middleware"
)

const apiVersion = "v1"

// NewHTTPRouter constructor
func NewHTTPRouter(
	config *middleware.Config,
	metricsHandler http.Handler,
	healthHandler handler.Health,
	todoHandler handler.Todo,
) http.Handler {

	r := chi.NewRouter()
	r.Use(chi_middleware.Recoverer)
	r.Use(chi_middleware.RequestLogger(log.NewLogFormatter(log.Logger)))
	r.Use(chi_middleware.RequestID)
	r.Use(chi_middleware.RealIP)
	r.Get("/health", healthHandler.Get)
	r.Get("/metrics", http.HandlerFunc(metricsHandler.ServeHTTP))

	// version1 apis
	r.Route(`/v1`, func(r chi.Router) {
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
	})

	return r
}
