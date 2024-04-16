package container

import (
	"github.com/nofendian17/gostarterkit/internal/config"
	"github.com/nofendian17/gostarterkit/internal/infra/cache"
	"github.com/nofendian17/gostarterkit/internal/infra/database"
	"github.com/nofendian17/gostarterkit/internal/usecase"
	"github.com/nofendian17/gostarterkit/pkg/logger"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	cfg := config.New()
	l := logger.New(logger.Config{
		File: logger.File{
			IsActive: false,
			LogFile:  "/tmp/app.log",
			Format:   "json",
		},
		Console: logger.Console{
			Format: "text",
		},
	})
	type args struct {
		cfg *config.Config
		db  *database.DB
		c   cache.Client
		l   logger.Logger
	}
	tests := []struct {
		name string
		args args
		want *Container
	}{
		{
			name: "success",
			args: args{
				cfg: cfg,
				db:  nil,
				c:   nil,
				l:   l,
			},
			want: &Container{
				Config:  cfg,
				UseCase: usecase.New(cfg, nil, nil),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.cfg, tt.args.db, tt.args.c, tt.args.l); !reflect.DeepEqual(got, tt.want) {
				assert.IsType(t, tt.want, got)
			}
		})
	}
}
