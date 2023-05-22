package routers

import (
	"backend/api/handlers"
	"net/http"
)

// Server represents an HTTP server.
type Server struct {
	Router   *http.ServeMux
	Handlers *handlers.Server
}
