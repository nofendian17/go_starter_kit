package cmd

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/nofendian17/gostarterkit/internal/container"
	"github.com/nofendian17/gostarterkit/internal/delivery/rest"
)

// Server is responsible for starting and stopping the REST server.
type Server interface {
	StartRestServer(ctx context.Context) error
}

type server struct {
	cntr *container.Container
	rest rest.Server
}

// New creates a new instance of the Server.
func New(cntr *container.Container, rest rest.Server) Server {
	return &server{
		cntr: cntr,
		rest: rest,
	}
}

// StartRestServer starts the REST server and handles graceful shutdown.
func (s *server) StartRestServer(ctx context.Context) error {
	// Start the REST server in a separate goroutine
	errCh := make(chan error, 1)
	go func() {
		errCh <- s.rest.Start(s.cntr.Config.Application.Port)
	}()

	// Handle OS signals for graceful shutdown
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

	select {
	case sig := <-signalCh:
		return fmt.Errorf("received signal %v", sig)
	case err := <-errCh:
		if err != nil {
			s.cntr.Logger.Error(ctx, "Failed to start server", err)
			return err
		}
	}

	// Stop the REST server
	if err := s.rest.Stop(ctx); err != nil {
		return fmt.Errorf("failed to stop server: %v", err)
	}

	return nil
}
