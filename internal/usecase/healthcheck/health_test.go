package healthcheck

import (
	"context"
	"github.com/nofendian17/gostarterkit/internal/config"
	"github.com/nofendian17/gostarterkit/internal/delivery/rest/model/response"
	"github.com/nofendian17/gostarterkit/internal/infra/cache"
	"github.com/nofendian17/gostarterkit/internal/infra/database"
	mockCacheClient "github.com/nofendian17/gostarterkit/internal/mocks/infra/cache"
	mockDB "github.com/nofendian17/gostarterkit/internal/mocks/infra/database"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_useCase_Health(t *testing.T) {
	cfg := config.New()
	_mockDB, _, _ := mockDB.New("mysql")
	_mockCacheClient := &mockCacheClient.Client{}
	type fields struct {
		startAt     time.Time
		config      *config.Config
		db          *database.DB
		cacheClient cache.Client
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *response.HealthResponse
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				startAt:     time.Now(),
				config:      cfg,
				db:          (*database.DB)(_mockDB),
				cacheClient: _mockCacheClient,
			},
			args: args{
				ctx: context.Background(),
			},
			want:    &response.HealthResponse{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &useCase{
				startAt:     tt.fields.startAt,
				config:      tt.fields.config,
				db:          tt.fields.db,
				cacheClient: tt.fields.cacheClient,
			}
			got, err := u.Health(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.IsType(t, tt.want, got)
		})
	}
}
