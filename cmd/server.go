package cmd

import (
	"context"
	"fmt"
	"github.com/nofendian17/gostarterkit/internal/container"
	"github.com/nofendian17/gostarterkit/internal/delivery/rest"
	"os"
	"os/signal"
	"syscall"
)

type Server interface {
	StartRestServer(ctx context.Context) error
}

type server struct {
	cntr *container.Container
	rest rest.Server
}

// StartRestServer starts the rest server.
func (s *server) StartRestServer(ctx context.Context) error {
	// Channel to catch errors
	errCh := make(chan error)

	// Start REST server in a separate goroutine
	go func() {
		errCh <- s.rest.Start(s.cntr.Config.Application.Port)
	}()

	// Handle OS signals for graceful shutdown
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

	select {
	case sig := <-signalCh:
		errCh <- fmt.Errorf("received signal %v", sig)
	case err := <-errCh:
		if err != nil {
			s.cntr.Logger.Error(ctx, "Got error signal", err)
		}
	}

	// Stop the REST server
	if err := s.rest.Stop(context.Background()); err != nil {
		return fmt.Errorf("failed to stop server: %v", err)
	}

	return nil
}

func New(cntr *container.Container, rest rest.Server) Server {
	return &server{
		cntr: cntr,
		rest: rest,
	}
}
