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

type Server interface {
	StartRestServer(ctx context.Context) error
}

type server struct {
	cntr *container.Container
	rest rest.Server
}

// StartRestServer starts the rest server.
func (s *server) StartRestServer(ctx context.Context) error {

	// Start REST server in a separate goroutine
	go func() {
		err := s.rest.Start(s.cntr.Config.Application.Port)
		if err != nil {
			s.cntr.Logger.Error(ctx, "Failed to start server", err)
		}
	}()

	// Handle OS signals for graceful shutdown
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

	sig := <-signalCh
	s.cntr.Logger.Info(ctx, fmt.Sprintf("Received signal: %s", sig), nil)

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
