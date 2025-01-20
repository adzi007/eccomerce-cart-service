package cachestore

import (
	"cart-service/internal/domain"
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type redisCache struct {
	client *redis.Client
	ctx    context.Context
}

// NewRedisCache creates a new Redis cache instance
func NewRedisCache(ctx context.Context, addr string, password string, db int) domain.CacheRepository {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	return &redisCache{
		client: client,
		ctx:    ctx,
	}
}

func (r *redisCache) Get(key string) (string, error) {
	val, err := r.client.Get(r.ctx, key).Result()
	if err == redis.Nil {
		return "", nil // Key not found
	}
	return val, err
}

func (r *redisCache) Set(key string, value string, expiration int64) error {
	return r.client.Set(r.ctx, key, value, time.Duration(expiration)*time.Second).Err()
}

func (r *redisCache) Delete(key string) error {
	return r.client.Del(r.ctx, key).Err()
}
