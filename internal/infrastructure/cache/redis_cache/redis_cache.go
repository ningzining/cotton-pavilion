package redis_cache

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
	"user-center/internal/infrastructure/cache"
)

type RedisCache struct {
	client *redis.Client
}

func (r RedisCache) Get(key string) string {
	return r.client.Get(context.Background(), key).String()
}

func (r RedisCache) Set(key string, value string, expiration time.Duration) error {
	return r.client.Set(context.Background(), key, value, expiration).Err()
}

func (r RedisCache) Remove(key ...string) error {
	return r.client.Del(context.Background(), key...).Err()
}

func NewRedisCache(addr, password string, db int) *RedisCache {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db, // use default DB
		Protocol: 3,  // specify 2 for RESP 2 or 3 for RESP 3
	})

	return &RedisCache{
		client: client,
	}
}

var _ cache.ICache = RedisCache{}
