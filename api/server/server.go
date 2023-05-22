package server

import (
	"backend/api/handlers"
	"backend/api/routers"
	"backend/database"
	Mongo "backend/database/mongo"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

// NewAuth creates a new authentication instance for the server.
func NewAuth() (*http.ServeMux, *mongo.Client) {
	r := http.NewServeMux()
	m := Mongo.ConnectMongoDB()
	return r, m
}

// NewServer creates a new server instance.
func NewServer(r *http.ServeMux, m *mongo.Client) *routers.Server {
	s := &routers.Server{
		Router: r,
		Handlers: &handlers.Server{
			DB: &database.DB{
				Mongo: &Mongo.Server{
					Client: m,
				},
			},
		},
	}
	return s
}

// Start initializes and starts the server.
func Start(port string) error {
	r, d := NewAuth()
	s := NewServer(r, d)
	s.SetupRouters()
	server := routers.UseMw(s.Router)
	log.Println("API Listen ON", port)
	return http.ListenAndServe(port, server)
}
