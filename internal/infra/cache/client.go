package cache

import (
	"context"
	"fmt"
	"github.com/nofendian17/gostarterkit/internal/config"
	"github.com/redis/go-redis/v9"
	"time"
)

type Client interface {
	Ping(ctx context.Context) error
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, value string, ttl time.Duration) error
	Del(ctx context.Context, key string) error
}

type client struct {
	rdb *redis.Client
}

func New(cfg *config.Config) (Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Cache.Host, cfg.Cache.Port),
		Username: cfg.Cache.Username,
		Password: cfg.Cache.Password,
		DB:       cfg.Cache.Database,
	})

	err := rdb.Ping(context.Background()).Err()
	if err != nil {
		return nil, err
	}

	return &client{rdb: rdb}, nil
}

func (c *client) Ping(ctx context.Context) error {
	_, err := c.rdb.Ping(ctx).Result()
	if err != nil {
		return fmt.Errorf("client ping failed: %w", err)
	}
	return nil
}

func (c *client) Get(ctx context.Context, key string) (string, error) {
	val, err := c.rdb.Get(ctx, key).Result()
	if err != nil {
		return "", fmt.Errorf("client get failed: %w", err)
	}
	return val, nil
}

func (c *client) Set(ctx context.Context, key string, value string, ttl time.Duration) error {
	err := c.rdb.Set(ctx, key, value, ttl).Err()
	if err != nil {
		return fmt.Errorf("client set failed: %w", err)
	}
	return nil
}

func (c *client) Del(ctx context.Context, key string) error {
	err := c.rdb.Del(ctx, key).Err()
	if err != nil {
		return fmt.Errorf("client delete failed: %w", err)
	}
	return nil
}
