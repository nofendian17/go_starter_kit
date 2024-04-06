package container

import (
	"github.com/nofendian17/gostarterkit/internal/config"
	"github.com/nofendian17/gostarterkit/internal/infra/cache"
	"github.com/nofendian17/gostarterkit/internal/infra/database"
	"github.com/nofendian17/gostarterkit/internal/usecase"
)

type Container struct {
	Config   *config.Config
	UseCase  *usecase.UseCase
	Database *database.DB
	Cache    cache.Client
}

// New initializes and returns a new Container with the given configuration.
func New(cfg *config.Config) *Container {
	db := database.New(cfg)
	c := cache.New(cfg)
	return &Container{
		Config:   cfg,
		UseCase:  usecase.New(cfg, db, c),
		Database: db,
		Cache:    c,
	}
}
