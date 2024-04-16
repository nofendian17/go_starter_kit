package healthcheck

import (
	"github.com/nofendian17/gostarterkit/internal/config"
	"github.com/nofendian17/gostarterkit/internal/container"
	"github.com/nofendian17/gostarterkit/internal/usecase/healthcheck"
	"github.com/nofendian17/gostarterkit/pkg/logger"
	"net/http"
)

type Handler interface {
	Ping() http.HandlerFunc
	Readiness() http.HandlerFunc
	Health() http.HandlerFunc
}

type handler struct {
	config  *config.Config
	useCase healthcheck.UseCase
	logger  logger.Logger
}

func New(c *container.Container) Handler {
	return &handler{
		config:  c.Config,
		useCase: c.UseCase.Health,
		logger:  c.Logger,
	}
}
