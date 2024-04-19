package main

import (
	"context"
	"github.com/nofendian17/gostarterkit/cmd"
	"github.com/nofendian17/gostarterkit/internal/config"
	"github.com/nofendian17/gostarterkit/internal/container"
	"github.com/nofendian17/gostarterkit/internal/delivery/rest"
	"github.com/nofendian17/gostarterkit/internal/infra/cache"
	"github.com/nofendian17/gostarterkit/internal/infra/database"
	"github.com/nofendian17/gostarterkit/pkg/logger"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
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
		panic(err)
	}

	// Initialize cache
	c, err := cache.New(cfg)
	if err != nil {
		panic(err)
	}

	// Initialize container
	cntr := container.New(cfg, db, c, l)

	// Initialize rest server
	restServer := rest.New(cntr)
	srv := cmd.New(cntr, restServer)
	err = srv.StartRestServer(context.Background())
	if err != nil {
		panic(err)
	}
}
