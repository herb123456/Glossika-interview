package mycache

import (
	"Glossika_interview/myredis"
	"context"
	"encoding/json"
	"github.com/eko/gocache/lib/v4/cache"
	"github.com/eko/gocache/lib/v4/store"
	redis_store "github.com/eko/gocache/store/redis/v4"
	"time"
)

type RedisCache struct {
	cacheManager *cache.Cache[string]
}

func NewRedisCache(ctx context.Context) *RedisCache {
	redisStore := redis_store.NewRedis(myredis.RedisConnection())

	cacheManager := cache.New[string](redisStore)

	return &RedisCache{
		cacheManager: cacheManager,
	}
}

func (c *RedisCache) Set(ctx context.Context, key string, value interface{}, expiredSeconds int64) error {
	// json encode value
	jsonByte, err := json.Marshal(value)
	if err != nil {
		return err
	}
	stringValue := string(jsonByte)

	// set to cache
	return c.cacheManager.Set(ctx, key, stringValue, store.WithExpiration(time.Duration(expiredSeconds)*time.Second))
}

func (c *RedisCache) Get(ctx context.Context, key string) (interface{}, error) {
	res, err := c.cacheManager.Get(ctx, key)
	if err != nil {
		return nil, err
	}

	// json decode value
	var value interface{}
	err = json.Unmarshal([]byte(res), &value)

	return value, err
}
