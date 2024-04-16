package healthcheck

import (
	"github.com/nofendian17/gostarterkit/internal/config"
	"github.com/nofendian17/gostarterkit/internal/container"
	mockCacheClient "github.com/nofendian17/gostarterkit/internal/mocks/infra/cache"
	"github.com/nofendian17/gostarterkit/internal/usecase"
	"github.com/nofendian17/gostarterkit/pkg/logger"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewHandler(t *testing.T) {
	cfg := config.New()

	c := &mockCacheClient.Client{}
	u := usecase.New(cfg, nil, c)

	l := logger.New(logger.Config{
		File: logger.File{
			IsActive: false,
			LogFile:  "log/app.log",
			Format:   "json",
		},
		Console: logger.Console{
			Format: "text",
		},
	})
	cntr := &container.Container{
		Config:  cfg,
		UseCase: u,
		Logger:  l,
	}

	type args struct {
		c *container.Container
	}
	tests := []struct {
		name string
		args args
		want Handler
	}{
		{
			name: "should return a new handler",
			args: args{
				c: cntr,
			},
			want: &handler{
				config:  cfg,
				useCase: u.Health,
				logger:  l,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := New(tt.args.c)
			assert.Equal(t, tt.want, got)
		})
	}
}
