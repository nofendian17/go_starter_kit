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
	"os"
	"os/signal"
	"syscall"
	"time"
)

type ServerInterface interface {
	Start(ctx context.Context, port int) error
}

// Server represents the HTTP server.
type Server struct {
	router  *http.ServeMux
	handler *handler.Handler
}

// Start starts the HTTP server.
func (s *Server) Start(ctx context.Context, port int) error {
	stack := middleware.Stack(
		middleware.Cors,
		middleware.RequestID,
	)

	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: stack(s.router),
	}

	// Start the HTTP server in a separate goroutine
	go func() {
		if err := httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Errorf("Failed to start HTTP server: %v", err)
		}
	}()

	// Wait for an OS interrupt signal
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	signal.Notify(ch, syscall.SIGTERM)
	<-ch

	// Create a context with a timeout for graceful shutdown
	ctxShutDown, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	// Attempt to gracefully shut down the HTTP server
	if err := httpServer.Shutdown(ctxShutDown); err != nil {
		slog.Errorf("Failed to gracefully shutdown HTTP server: %v", err)
		return err
	}

	slog.Printf("HTTP server shutdown completed.")
	return nil
}

// New creates a new instance of the HTTP server.
func New(c *container.Container) ServerInterface {
	srv := &Server{
		router:  http.NewServeMux(),
		handler: handler.New(c),
	}
	srv.routes()
	return srv
}
