package cache

import (
	"context"
	"errors"
	"github.com/go-redis/redismock/v9"
	"github.com/nofendian17/gostarterkit/internal/config"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNew(t *testing.T) {

	cfg := config.New()
	cacheClient, _ := New(cfg)
	type args struct {
		cfg *config.Config
	}
	tests := []struct {
		name string
		args args
		want Client
	}{
		{
			name: "success",
			args: args{
				cfg: cfg,
			},
			want: cacheClient,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := cacheClient
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_client_Del(t *testing.T) {

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
				ctx: context.TODO(),
				key: "key",
			},
			error:   nil,
			wantErr: false,
		},
		{
			name: "failed - got error delete data from redis",
			args: args{
				ctx: context.TODO(),
				key: "key",
			},
			error:   errors.New("error delete data from redis"),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock := redismock.NewClientMock()
			if !tt.wantErr {
				mock.ExpectDel(tt.args.key).SetVal(0)
			} else {
				mock.ExpectGet(tt.args.key).SetErr(tt.error)
			}
			defer mock.ClearExpect()
			c := &client{
				rdb: db,
			}
			if err := c.Del(tt.args.ctx, tt.args.key); (err != nil) != tt.wantErr {
				t.Errorf("Del() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_client_Get(t *testing.T) {

	type args struct {
		ctx context.Context
		key string
	}
	tests := []struct {
		name    string
		args    args
		error   error
		want    string
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				ctx: context.TODO(),
				key: "key",
			},
			want:    "value",
			wantErr: false,
		},
		{
			name: "failed - got error get data from redis",
			args: args{
				ctx: context.TODO(),
				key: "key",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock := redismock.NewClientMock()
			if !tt.wantErr {
				mock.ExpectGet(tt.args.key).SetVal(tt.want)
			} else {
				mock.ExpectGet(tt.args.key).SetErr(tt.error)
			}
			defer mock.ClearExpect()
			c := &client{
				rdb: db,
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

func Test_client_Ping(t *testing.T) {

	type args struct {
		ctx context.Context
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
				ctx: context.TODO(),
			},
			error:   nil,
			wantErr: false,
		},
		{
			name: "failed - got error ping redis",
			args: args{
				ctx: context.TODO(),
			},
			error:   errors.New("error ping redis"),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock := redismock.NewClientMock()
			if !tt.wantErr {
				mock.ExpectPing().SetVal("OK")
			} else {
				mock.ExpectPing().SetErr(tt.error)
			}
			defer mock.ClearExpect()
			c := &client{
				rdb: db,
			}
			if err := c.Ping(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Ping() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_client_Set(t *testing.T) {
	type args struct {
		ctx   context.Context
		key   string
		value string
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
				ctx:   context.Background(),
				key:   "key",
				value: "value",
				ttl:   time.Minute * 1,
			},
			error:   nil,
			wantErr: false,
		},
		{
			name: "failed - got error when set to redis",
			args: args{
				ctx:   context.Background(),
				key:   "key",
				value: "value",
				ttl:   time.Minute * 1,
			},
			error:   errors.New("error set to redis"),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock := redismock.NewClientMock()
			if !tt.wantErr {
				mock.ExpectSet(tt.args.key, tt.args.value, tt.args.ttl).SetVal("OK")
			} else {
				mock.ExpectSet(tt.args.key, tt.args.value, tt.args.ttl).SetErr(tt.error)
			}
			defer mock.ClearExpect()
			c := &client{
				rdb: db,
			}
			if err := c.Set(tt.args.ctx, tt.args.key, tt.args.value, tt.args.ttl); (err != nil) != tt.wantErr {
				assert.Error(t, err)
			}
		})
	}
}
