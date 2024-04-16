package container

import (
	"github.com/nofendian17/gostarterkit/internal/config"
	"github.com/nofendian17/gostarterkit/internal/infra/cache"
	"github.com/nofendian17/gostarterkit/internal/infra/database"
	"github.com/nofendian17/gostarterkit/internal/usecase"
	"github.com/nofendian17/gostarterkit/pkg/logger"
)

type Container struct {
	Config  *config.Config
	UseCase *usecase.UseCase
	Logger  logger.Logger
}

// New initializes and returns a new Container with the given configuration.
func New(cfg *config.Config, db *database.DB, c cache.Client, l logger.Logger) *Container {
	return &Container{
		Config:  cfg,
		UseCase: usecase.New(cfg, db, c),
		Logger:  l,
	}
}
