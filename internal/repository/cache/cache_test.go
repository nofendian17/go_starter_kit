package cache

import (
	"context"
	"errors"
	"github.com/nofendian17/gostarterkit/internal/config"
	"github.com/nofendian17/gostarterkit/internal/entity"
	"github.com/nofendian17/gostarterkit/internal/infra/cache"
	mockCacheClient "github.com/nofendian17/gostarterkit/internal/mocks/infra/cache"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	cfg := config.New()

	cacheClient := cache.New(cfg)
	tests := []struct {
		name        string
		cacheClient cache.Client
		want        Repository
	}{
		{
			name:        "success",
			cacheClient: cacheClient,
			want: &repository{
				cache: cacheClient,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := New(tt.cacheClient)
			assert.Equal(t, tt.want, actual)
		})
	}
}

func Test_repository_Delete(t *testing.T) {
	type args struct {
		ctx context.Context
		key string
	}
	tests := []struct {
		name    string
		args    args
		error   error
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				ctx: context.Background(),
				key: "key",
			},
			error:   nil,
			wantErr: false,
		},
		{
			name: "failed - got error while deleting key",
			args: args{
				ctx: context.Background(),
				key: "key",
			},
			error:   errors.New("error while deleting key"),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_mockCacheClient := &mockCacheClient.Client{}
			_mockCacheClient.On("Del", mock.Anything, mock.Anything).Return(tt.error)
			c := &repository{
				cache: _mockCacheClient,
			}
			if err := c.Delete(tt.args.ctx, tt.args.key); (err != nil) != tt.wantErr {
				assert.Error(t, err)
			}
		})
	}
}

func Test_repository_Get(t *testing.T) {
	type args struct {
		ctx context.Context
		key string
	}
	tests := []struct {
		name    string
		args    args
		want    *entity.Cache
		result  []interface{}
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				ctx: context.Background(),
				key: "key",
			},
			want: &entity.Cache{},
			result: []interface{}{
				`{"foo":"bar"}`, nil,
			},
			wantErr: false,
		},
		{
			name: "failed - got error while unmarshalling the result",
			args: args{
				ctx: context.Background(),
				key: "key",
			},
			want: nil,
			result: []interface{}{
				``, nil,
			},
			wantErr: true,
		},
		{
			name: "failed - got error while fetching key",
			args: args{
				ctx: context.Background(),
				key: "key",
			},
			want: nil,
			result: []interface{}{
				``, errors.New("error while fetching key"),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_mockCacheClient := &mockCacheClient.Client{}
			_mockCacheClient.On("Get", mock.Anything, mock.Anything).Return(tt.result...)
			c := &repository{
				cache: _mockCacheClient,
			}
			got, err := c.Get(tt.args.ctx, tt.args.key)
			if (err != nil) != tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_repository_Set(t *testing.T) {

	type args struct {
		ctx   context.Context
		key   string
		cache entity.Cache
		ttl   time.Duration
	}
	tests := []struct {
		name    string
		args    args
		error   error
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				ctx: context.Background(),
				key: "key",
				cache: entity.Cache{
					Async: entity.AsyncCache{
						Identifier:  "key_123",
						CallbackURL: "http://localhost:8080/callback",
					},
				},
				ttl: time.Second * 60,
			},
			error:   nil,
			wantErr: false,
		},
		{
			name: "failed - got error while storing cache",
			args: args{
				ctx: context.Background(),
				key: "key",
				cache: entity.Cache{
					Async: entity.AsyncCache{
						Identifier:  "key_123",
						CallbackURL: "http://localhost:8080/callback",
					},
				},
				ttl: time.Second * 60,
			},
			error:   errors.New("error while storing cache"),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_mockCacheClient := &mockCacheClient.Client{}
			_mockCacheClient.On("Set", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(tt.error)
			c := &repository{
				cache: _mockCacheClient,
			}
			if err := c.Set(tt.args.ctx, tt.args.key, tt.args.cache, tt.args.ttl); (err != nil) != tt.wantErr {
				assert.Error(t, err)
			}
		})
	}
}
