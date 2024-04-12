package healthcheck

import (
	"github.com/nofendian17/gostarterkit/internal/config"
	"github.com/nofendian17/gostarterkit/internal/infra/cache"
	"github.com/nofendian17/gostarterkit/internal/infra/database"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	cfg := config.New()
	xtime := time.Now()
	type args struct {
		t           time.Time
		c           *config.Config
		db          *database.DB
		cacheClient cache.Client
	}
	tests := []struct {
		name string
		args args
		want *useCase
	}{
		{
			name: "success",
			args: args{
				t:           xtime,
				c:           cfg,
				db:          nil,
				cacheClient: nil,
			},
			want: &useCase{
				startAt: xtime,
				config:  cfg,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, New(tt.args.t, tt.args.c, tt.args.db, tt.args.cacheClient), "New(%v, %v, %v, %v)", tt.args.t, tt.args.c, tt.args.db, tt.args.cacheClient)
		})
	}
}
