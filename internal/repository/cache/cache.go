package cache

import (
	"context"
	"encoding/json"
	"github.com/nofendian17/gostarterkit/internal/entity"
	cacheClient "github.com/nofendian17/gostarterkit/internal/infra/cache"
	"time"
)

type Repository interface {
	Set(ctx context.Context, key string, cache entity.Cache, ttl time.Duration) error
	Get(ctx context.Context, key string) (*entity.Cache, error)
	Delete(ctx context.Context, key string) error
}

type repository struct {
	cache cacheClient.Client
}

func (c *repository) Set(ctx context.Context, key string, cache entity.Cache, ttl time.Duration) error {
	data, err := json.Marshal(cache)
	if err != nil {
		return err
	}
	return c.cache.Set(ctx, key, string(data), ttl)
}

func (c *repository) Get(ctx context.Context, key string) (*entity.Cache, error) {
	val, err := c.cache.Get(ctx, key)
	if err != nil {
		return nil, err
	}

	var result entity.Cache
	err = json.Unmarshal([]byte(val), &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *repository) Delete(ctx context.Context, key string) error {
	return c.cache.Del(ctx, key)
}

func New(cache cacheClient.Client) Repository {
	return &repository{
		cache: cache,
	}
}
