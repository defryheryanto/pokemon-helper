package main

import (
	"fmt"

	"github.com/defry256/pokemon-helper/config"
	"github.com/go-redis/redis/v8"
)

func setupRedis() *redis.Client {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.RedisHost(), config.RedisPort()),
		Username: config.RedisUsername(),
		Password: config.RedisPassword(),
	})

	return redisClient
}
