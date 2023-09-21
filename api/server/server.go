package server

import (
	"backend/api/handlers"
	"backend/api/routers"
	"backend/database"
	Mongo "backend/database/mongo"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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

// gracefulShutDown gracefully shuts down the HTTP server.
func gracefulShutDown(serverStopCh <-chan os.Signal, server *http.Server) {
	<-serverStopCh
	log.Println("Shutting down the server...")
	// Create a context with a timeout to allow existing connections to finish
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Shutdown the server gracefully
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Error during server shutdown: %v\n", err)
	}
	log.Println("Server gracefully shut down")
}

// Start initializes and starts the server.
func Start(port string) {
	r, d := NewAuth()
	s := NewServer(r, d)
	s.SetupRouters()
	handler := routers.UseMw(s.Router)

	server := &http.Server{
		Addr:    port,
		Handler: handler,
	}
	// Create a channel to listen for OS signals
	serverStopCh := make(chan os.Signal, 1)
	signal.Notify(serverStopCh, syscall.SIGINT, syscall.SIGTERM)
	// Start the server in a separate goroutine
	go func() {
		log.Println("Server is listening on", port)
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("Error: %v\n", err)
		}
	}()
	gracefulShutDown(serverStopCh, server)
}
