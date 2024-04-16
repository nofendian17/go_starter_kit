package cmd

import (
	"context"
	"fmt"
	"github.com/nofendian17/gostarterkit/internal/config"
	"github.com/nofendian17/gostarterkit/internal/infra/cache"
	"github.com/nofendian17/gostarterkit/internal/infra/database"
	"github.com/nofendian17/gostarterkit/pkg/logger"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/nofendian17/gostarterkit/internal/container"
	"github.com/nofendian17/gostarterkit/internal/delivery/rest"
)

// Run starts the application.
func Run() error {
	ctx := context.Background()

	// Initialize config
	cfg := config.New()

	// Initialize log
	l := logger.New(logger.Config{
		Output:  cfg.Logger.Output,
		Level:   cfg.Logger.Level,
		Service: cfg.Application.Name,
		Version: cfg.Application.Version,
	})

	// Initialize db
	db, err := database.New(cfg, l)
	if err != nil {
		return err
	}

	// Initialize cache
	c, err := cache.New(cfg)
	if err != nil {
		return err
	}

	cntr := container.New(cfg, db, c, l)
	restServer := rest.New(cntr)

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
	err = <-errCh
	if err != nil {
		l.Error(ctx, "Got error signal", err)
	}

	// Create a context for graceful shutdown
	ctxCancel, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	// Stop the REST server
	if err := restServer.Stop(ctxCancel); err != nil {
		return fmt.Errorf("failed to stop server: %v", err)
	}

	return nil
}
