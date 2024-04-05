package healthcheck

import (
	"context"
	"github.com/nofendian17/gostarterkit/internal/config"
	"github.com/nofendian17/gostarterkit/internal/delivery/rest/model/response"
	cacheClient "github.com/nofendian17/gostarterkit/internal/infra/cache"
	"github.com/nofendian17/gostarterkit/internal/infra/database"
	"time"
)

type UseCase interface {
	Heath(ctx context.Context) (*response.HealthResponse, error)
	Readiness(ctx context.Context) (*response.ReadinessResponse, error)
}

type useCase struct {
	startAt     time.Time
	config      *config.Config
	db          *database.DB
	cacheClient cacheClient.Client
}

func New(t time.Time, c *config.Config, db *database.DB, cacheClient cacheClient.Client) UseCase {
	return &useCase{
		startAt:     t,
		config:      c,
		db:          db,
		cacheClient: cacheClient,
	}
}
