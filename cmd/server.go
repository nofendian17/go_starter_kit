package cmd

import (
	"context"
	"github.com/nofendian17/gostarterkit/internal/config"
	"github.com/nofendian17/gostarterkit/internal/container"
	"github.com/nofendian17/gostarterkit/internal/delivery/rest"
)

func Run() error {
	cfg := config.New()
	c := container.New(cfg)

	restServer := rest.New(c)
	err := restServer.Start(context.Background(), cfg.Application.Port)
	if err != nil {
		return err
	}
	return nil
}
