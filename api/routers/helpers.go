package routers

import (
	"backend/api/middleware"
	"net/http"
)

type methods struct {
	h http.HandlerFunc
}

// Methods restricts an HTTP handler function to specified HTTP methods.
func (m *methods) Methods(s ...string) http.HandlerFunc {
	return middleware.MethodRestricted(m.h, s)
}

// Handle wraps a handler function with a method restriction middleware.
func Handle(h http.HandlerFunc) *methods {
	return &methods{h: h}
}

// ApplyMiddlewares is a function that applies multiple middlewares to an HTTP handler.
func ApplyMiddlewares(handler http.Handler,
	middlewares ...func(http.Handler) http.Handler) http.Handler {
	// Wrap the handler with the specified middlewares
	for _, mw := range middlewares {
		handler = mw(handler)
	}
	return handler
}
