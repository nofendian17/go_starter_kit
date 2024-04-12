package repository

import (
	"github.com/nofendian17/gostarterkit/internal/config"
	"github.com/nofendian17/gostarterkit/internal/infra/cache"
	cacheRepository "github.com/nofendian17/gostarterkit/internal/repository/cache"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	cfg := config.New()
	cacheClient, _ := cache.New(cfg)
	type args struct {
		client cache.Client
	}
	tests := []struct {
		name string
		args args
		want *Repository
	}{
		{
			name: "success",
			args: args{
				client: cacheClient,
			},
			want: &Repository{
				CacheRepository: cacheRepository.New(cacheClient),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.client); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
