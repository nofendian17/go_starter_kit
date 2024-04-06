package container

import (
	"github.com/nofendian17/gostarterkit/internal/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	cfg := config.New()

	c := New(cfg)

	// Ensure that the container is not nil
	assert.NotNil(t, c, "expected container to be initialized")

	// Ensure that the configuration in the container matches the input configuration
	assert.Equal(t, cfg, c.Config, "expected configuration to match")

	// Ensure that the UseCase dependency is initialized
	assert.NotNil(t, c.UseCase, "expected UseCase to be initialized")

}
