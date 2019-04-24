package middleware

import (
	"github.com/go-chi/cors"
)

// NewCors Cors middleware
func NewCors(config *CorsConfig) *cors.Cors {

	return cors.New(cors.Options{
		AllowedOrigins: config.AllowedOrigins,
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods: config.AllowedMethods,
		AllowedHeaders: config.AllowedHeaders,
		// ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
}
