package rest

import (
	"context"
	"fmt"
	"github.com/nofendian17/gostarterkit/internal/config"
	mockCacheClient "github.com/nofendian17/gostarterkit/internal/mocks/infra/cache"
	"github.com/nofendian17/gostarterkit/internal/usecase"
	"github.com/nofendian17/gostarterkit/pkg/logger"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/nofendian17/gostarterkit/internal/container"
	"github.com/stretchr/testify/assert"
)

func TestServer_StartStop(t *testing.T) {
	cfg := config.New()
	l := logger.New(logger.Config{
		Output:  cfg.Logger.Output,
		Level:   cfg.Logger.Level,
		Service: cfg.Application.Name,
		Version: cfg.Application.Version,
	})

	c := &mockCacheClient.Client{}
	u := usecase.New(cfg, nil, c)

	// Create a container with necessary dependencies
	cntr := &container.Container{
		Config:  cfg,
		UseCase: u,
		Logger:  l,
	}

	// Create a new HTTP server instance
	srv := New(cntr)

	// Start the server on a random available port
	go func() {
		err := srv.Start(cfg.Application.Port)
		assert.NoError(t, err, "expected no error when starting server")
	}()

	// Wait for the server to start
	time.Sleep(100 * time.Millisecond)

	// Create a new HTTP request to test the server
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://%s:%d/ping", cfg.Application.Address, cfg.Application.Port), nil)
	assert.NoError(t, err, "expected no error when creating request")

	// Create a new HTTP recorder to capture the response
	rec := httptest.NewRecorder()

	// Perform a dummy request to the server
	srv.(*server).router.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code, "expected 200 response")

	// Stop the server
	err = srv.Stop(context.Background())
	assert.NoError(t, err, "expected no error when stopping server")

	//// Try to stop the server again (should return an error)
	//err = srv.Stop(context.Background())
	//assert.Error(t, err, "expected error when stopping already stopped server")
	//assert.True(t, errors.Is(err, http.ErrServerClosed), "expected specific error when stopping already stopped server")
}
