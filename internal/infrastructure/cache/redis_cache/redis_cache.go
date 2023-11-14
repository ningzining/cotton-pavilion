package redis_cache

import (
	"github.com/redis/go-redis/v9"
)

type RedisCache struct {
	client *redis.Client
}

func (r RedisCache) Client() *redis.Client {
	return r.client
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
