package rest

import (
	"context"
	"errors"
	"fmt"
	"github.com/gookit/slog"
	"github.com/nofendian17/gostarterkit/internal/container"
	"github.com/nofendian17/gostarterkit/internal/delivery/rest/handler"
	"github.com/nofendian17/gostarterkit/internal/delivery/rest/middleware"
	"net/http"
	"time"
)

// ServerInterface defines the methods for an HTTP server.
type ServerInterface interface {
	Start(port int) error
	Stop(ctx context.Context) error
}

// server represents the HTTP server.
type server struct {
	router     *http.ServeMux
	handler    *handler.Handler
	httpServer *http.Server
}

// Start starts the HTTP server.
func (s *server) Start(port int) error {
	stack := middleware.Stack(
		middleware.Cors,
		middleware.RequestID,
	)

	s.httpServer = &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: stack(s.router),
	}

	if err := s.httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		slog.Errorf("Failed to start HTTP server: %v", err)
		return err
	}

	return nil
}

// Stop gracefully shuts down the HTTP server.
func (s *server) Stop(ctx context.Context) error {
	// Create a context with a timeout for graceful shutdown
	ctxShutDown, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	// Attempt to gracefully shut down the HTTP server
	if err := s.httpServer.Shutdown(ctxShutDown); err != nil {
		slog.Errorf("Failed to gracefully shutdown HTTP server: %v", err)
		return err
	}

	slog.Printf("HTTP server shutdown completed.")
	return nil
}

// New creates a new instance of the HTTP server.
func New(c *container.Container) ServerInterface {
	srv := &server{
		router:  http.NewServeMux(),
		handler: handler.New(c),
	}
	srv.routes()
	return srv
}
