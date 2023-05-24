package routers

import (
	_ "backend/api/docs" // docs is generated by Swag CLI
	"backend/api/middleware"
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger/v2"
)

// UseMw applies custom middlewares to the router.
func UseMw(handler http.Handler) http.Handler {
	return ApplyMiddlewares(handler,
		// add your middlewares here
		middleware.Logger,
		middleware.SetHeaders,
		middleware.RequestSize(1024*10), //10 KB body
	)
}

// SetupRouters sets up the routes for the server's HTTP router.
func (s *Server) SetupRouters() {
	s.Router.HandleFunc("/api/locations/show", Handle(s.Handlers.ShowLocations).Methods("GET"))
	s.Router.HandleFunc("/api/cars/showreservedcars", Handle(s.Handlers.ShowReservedCars).Methods("GET"))
	s.Router.HandleFunc("/api/cars/reserve", Handle(s.Handlers.ReserveCar).Methods("POST"))
	s.Router.HandleFunc("/api/cars/available", Handle(s.Handlers.AvailableCars).Methods("GET"))
	s.Router.HandleFunc("/api/swagger/", Handle(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8000/api/swagger/doc.json"))).Methods("GET"))
}