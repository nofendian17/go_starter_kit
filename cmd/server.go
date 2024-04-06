package cmd

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gookit/slog"
	"github.com/nofendian17/gostarterkit/internal/config"
	"github.com/nofendian17/gostarterkit/internal/container"
	"github.com/nofendian17/gostarterkit/internal/delivery/rest"
)

// Run starts the application.
func Run() error {
	cfg := config.New()
	c := container.New(cfg)
	restServer := rest.New(c)

	// Channel to catch errors
	errCh := make(chan error)

	// Start REST server in a separate goroutine
	go func() {
		errCh <- restServer.Start(cfg.Application.Port)
	}()

	// Handle OS signals for graceful shutdown
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

	go func() {
		sig := <-signalCh
		errCh <- fmt.Errorf("received signal %v", sig)
	}()

	// Wait for an error to occur
	err := <-errCh
	if err != nil {
		slog.Errorf("Got error signal: %v", err)
	}

	// Create a context for graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Stop the REST server
	if err := restServer.Stop(ctx); err != nil {
		return fmt.Errorf("failed to stop server: %v", err)
	}

	return nil
}
