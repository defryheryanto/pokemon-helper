package main

import (
	"fmt"

	"github.com/defry256/pokemon-helper/config"
	"github.com/go-redis/redis/v8"
)

func setupRedis() *redis.Client {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.REDIS_HOST(), config.REDIS_PORT()),
		Username: config.REDIS_USERNAME(),
		Password: config.REDIS_PASSWORD(),
	})

	return redisClient
}
