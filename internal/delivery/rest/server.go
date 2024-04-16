package rest

import (
	"context"
	"errors"
	"fmt"
	"github.com/nofendian17/gostarterkit/internal/container"
	"github.com/nofendian17/gostarterkit/internal/delivery/rest/handler"
	"github.com/nofendian17/gostarterkit/internal/delivery/rest/middleware"
	"github.com/nofendian17/gostarterkit/pkg/logger"
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
	logger     logger.Logger
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
		s.logger.Error(fmt.Sprintf("Failed to start HTTP server: %v", err), nil)
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
		s.logger.Error(fmt.Sprintf("Failed to gracefully shutdown HTTP server: %v", err), nil)
		return err
	}
	s.logger.Info("HTTP server shutdown completed.", nil)
	return nil
}

// New creates a new instance of the HTTP server.
func New(c *container.Container) ServerInterface {
	srv := &server{
		router:  http.NewServeMux(),
		handler: handler.New(c),
		logger:  c.Logger,
	}
	srv.routes()
	return srv
}
