package cmd

import (
	"context"
	"errors"
	"github.com/nofendian17/gostarterkit/internal/config"
	"github.com/nofendian17/gostarterkit/internal/container"
	"github.com/nofendian17/gostarterkit/internal/delivery/rest"
	mocks "github.com/nofendian17/gostarterkit/internal/mocks/delivery/rest"
	mockCacheClient "github.com/nofendian17/gostarterkit/internal/mocks/infra/cache"
	"github.com/nofendian17/gostarterkit/internal/usecase"
	"github.com/nofendian17/gostarterkit/pkg/logger"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		cntr *container.Container
		rest rest.Server
	}
	tests := []struct {
		name string
		args args
		want Server
	}{
		{
			name: "success",
			args: args{
				cntr: nil,
				rest: nil,
			},
			want: &server{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := New(tt.args.cntr, tt.args.rest)
			assert.Equal(t, tt.want, actual)
		})
	}
}

func Test_server_StartRestServer(t *testing.T) {
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
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name               string
		args               args
		restServerStartErr error
		restServerStopErr  error
		wantErr            bool
	}{
		{
			name: "success start and stop server",
			args: args{
				ctx: context.Background(),
			},
			restServerStartErr: nil,
			restServerStopErr:  nil,
			wantErr:            false,
		},
		{
			name: "failed - start server got error",
			args: args{
				ctx: context.Background(),
			},
			restServerStartErr: errors.New("start server error"),
			restServerStopErr:  nil,
			wantErr:            true,
		},
		{
			name: "failed - stop server got error",
			args: args{
				ctx: context.Background(),
			},
			restServerStartErr: nil,
			restServerStopErr:  errors.New("stop server error"),
			wantErr:            true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_mockRestServer := &mocks.Server{}
			_mockRestServer.On("Start", mock.Anything).Return(tt.restServerStartErr).Once()
			_mockRestServer.On("Stop", mock.Anything).Return(tt.restServerStopErr).Once()
			s := &server{
				cntr: cntr,
				rest: _mockRestServer,
			}
			err := s.StartRestServer(tt.args.ctx)
			if tt.wantErr && err != nil {
				assert.Error(t, err)
			}
		})
	}
}
