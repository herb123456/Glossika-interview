package myredis

import (
	"Glossika_interview/config"
	"github.com/redis/go-redis/v9"
)

func RedisConnection() *redis.Client {
	addr := config.GetString("redis.host") + ":" + config.GetString("redis.port")

	return redis.NewClient(&redis.Options{
		Addr: addr,
	})
}
