package usecase

import (
	"github.com/nofendian17/gostarterkit/internal/config"
	"github.com/nofendian17/gostarterkit/internal/infra/cache"
	"github.com/nofendian17/gostarterkit/internal/infra/database"
	"github.com/nofendian17/gostarterkit/internal/usecase/healthcheck"
	"time"
)

type UseCase struct {
	Health healthcheck.UseCase
}

// New creates a new instance of the UseCase struct, initializing it with the provided configuration and database.
func New(cfg *config.Config, db *database.DB, cache cache.Client) *UseCase {
	healthUseCase := healthcheck.New(time.Now(), cfg, db, cache)
	return &UseCase{
		Health: healthUseCase,
	}
}
